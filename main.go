package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type User struct {
	Id    int
	Name  string
	Email string
	Age   int
}

func Read(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println("Server failed to handle", err)
		return
	}

	defer rows.Close()
	data := make([]User, 0)
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Age)
		if err != nil {
			fmt.Println("Server failed to handle", err)
			return
		}
		data = append(data, user)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Server failed to handle", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)

}

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	u := User{}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		fmt.Println("Server failed to handle", err)
		return
	}

	_, err = db.Exec("INSERT INTO users (name, email, age) VALUES ($1, $2, $3)", u.Name, u.Email, u.Age)
	if err != nil {
		fmt.Println("Server failed to handle", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

var db *sql.DB

func init() {

	// Capturando as variáveis de ambiente
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	dbname := os.Getenv("DB_NAME")

	// Montando a string de conexão
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, password, host, dbname)

	var err error
	// Abrindo uma conexão com o postgres
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("You connected to your database.")
}

func main() {
	http.HandleFunc("/users/read", Read)
	http.HandleFunc("/users/create", Create)
	http.ListenAndServe(":8080", nil)
}
