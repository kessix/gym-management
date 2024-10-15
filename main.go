package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type User struct {
	Id    int
	Plan  Plan
	Name  string
	Email string
	Age   int
}

type Plan struct {
	Id    int
	Name  string
	Price float64
}

type Payment struct {
	Id          int
	UserId      int
	Month       string
	Status      bool
	PaymentDate *time.Time
}

func Read(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	query := `
        SELECT u.id, u.name, u.email, u.age, p.id, p.name, p.price
        FROM users u
        LEFT JOIN plans p ON u.plan_id = p.id
    `
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Server failed to handle", err)
		return
	}

	defer rows.Close()
	data := make([]User, 0)
	for rows.Next() {
		user := User{}
		plan := Plan{}
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Age, &plan.Id, &plan.Name, &plan.Price)
		if err != nil {
			fmt.Println("Server failed to handle", err)
			return
		}
		user.Plan = plan
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

	_, err = db.Exec("INSERT INTO users (name, email, age, plan_id) VALUES ($1, $2, $3, $4)", u.Name, u.Email, u.Age, u.Plan.Id)
	if err != nil {
		fmt.Println("Server failed to handle", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	up := User{}
	err := json.NewDecoder(r.Body).Decode(&up)
	if err != nil {
		fmt.Println("Server failed to handle", err)
		return
	}

	_, err = db.Exec("UPDATE users SET name=$1, email=$2, age=$3, plan_id=$4 WHERE id=$5", up.Name, up.Email, up.Age, up.Plan.Id, id)
	if err != nil {
		fmt.Println("Server failed to handle", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	_, err := db.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		fmt.Println("Server failed to handle", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func ReadPlan(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query("SELECT * FROM plans")
	if err != nil {
		fmt.Println("Server failed to handle", err)
		return
	}

	defer rows.Close()
	plans := make([]Plan, 0)
	for rows.Next() {
		plan := Plan{}
		err := rows.Scan(&plan.Id, &plan.Name, &plan.Price)
		if err != nil {
			fmt.Println("Server failed to handle", err)
			return
		}
		plans = append(plans, plan)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Server failed to handle", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(plans)
}

func CreatePlan(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	p := Plan{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		fmt.Println("Server failed to handle", err)
		return
	}

	_, err = db.Exec("INSERT INTO plans (name, price) VALUES ($1, $2)", p.Name, p.Price)
	if err != nil {
		fmt.Println("Server failed to handle", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func UpdatePlan(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	up := Plan{}
	err := json.NewDecoder(r.Body).Decode(&up)
	if err != nil {
		fmt.Println("Server failed to handle", err)
		return
	}

	_, err = db.Exec("UPDATE plans SET name=$1, price=$2 WHERE id=$3", up.Name, up.Price, id)
	if err != nil {
		fmt.Println("Server failed to handle", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeletePlan(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	_, err := db.Exec("DELETE FROM plans WHERE id=$1", id)
	if err != nil {
		fmt.Println("Server failed to handle", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func ReadPayment(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query("SELECT id, user_id, month, status, payment_date FROM payments")
	if err != nil {
		fmt.Println("Server failed to handle", err)
		return
	}
	defer rows.Close()

	payments := []Payment{}

	for rows.Next() {
		var p Payment
		err := rows.Scan(&p.Id, &p.UserId, &p.Month, &p.Status, &p.PaymentDate)
		if err != nil {
			fmt.Println("Server failed to handle sql:", err)
			return
		}
		payments = append(payments, p)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Server failed to handle", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(payments)
}

func CreatePayment(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	var p Payment
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		fmt.Println("Server failed to handle json:", err)
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Inserir no banco de dados, `p.PaymentDate` pode ser nil
	_, err = db.Exec("INSERT INTO payments (user_id, month, status, payment_date) VALUES ($1, $2, $3, $4)",
		p.UserId, p.Month, p.Status, p.PaymentDate)
	if err != nil {
		fmt.Println("Server failed to insert payment:", err)
		http.Error(w, "Failed to create payment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Payment created successfully")
}

func UpdatePayment(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing payment ID", http.StatusBadRequest)
		return
	}

	var p Payment
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		fmt.Println("Server failed to handle json:", err)
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Verificando se o pagamento existe antes de atualizar
	row := db.QueryRow("SELECT id FROM payments WHERE id = $1", id)
	var existingPayment Payment
	err = row.Scan(&existingPayment.Id)
	if err == sql.ErrNoRows {
		http.Error(w, "Payment not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	// Atualizando apenas os campos enviados
	_, err = db.Exec("UPDATE payments SET user_id = $1, month = $2, status = $3, payment_date = $4 WHERE id = $5",
		p.UserId, p.Month, p.Status, p.PaymentDate, id)
	if err != nil {
		fmt.Println("Server failed to update payment:", err)
		http.Error(w, "Failed to update payment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Payment updated successfully")
}

func DeletePayment(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	_, err := db.Exec("DELETE FROM payments WHERE id=$1", id)
	if err != nil {
		fmt.Println("Server failed to handle", err)
		return
	}

	w.WriteHeader(http.StatusOK)
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
	http.HandleFunc("/users/update", Update)
	http.HandleFunc("/users/delete", Delete)

	http.HandleFunc("/plans/read", ReadPlan)
	http.HandleFunc("/plans/create", CreatePlan)
	http.HandleFunc("/plans/update", UpdatePlan)
	http.HandleFunc("/plans/delete", DeletePlan)

	http.HandleFunc("/payments/read", ReadPayment)
	http.HandleFunc("/payments/create", CreatePayment)
	http.HandleFunc("/payments/update", UpdatePayment)
	http.HandleFunc("/payments/delete", DeletePayment)

	http.ListenAndServe(":8080", nil)
}
