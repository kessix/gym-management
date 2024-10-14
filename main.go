package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	_ "github.com/lib/pq" // driver do PostgreSQL
)

// Estruturas para usuários, planos e pagamentos
type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	PlanID int    `json:"plan_id"`
}

type Plan struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Payment struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Month       string `json:"month"`
	Status      bool   `json:"status"`
	PaymentDate string `json:"payment_date"`
}

// Conexão com o banco de dados
func connectDB() (*sql.DB, error) {
	// Capturando as variáveis de ambiente
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	dbname := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, password, host, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// CRUD para Usuários

// Criar Usuário
func createUser(db *sql.DB, name, email string, age int, planID int) error {
	query := `INSERT INTO users (name, email, age, plan_id) VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(query, name, email, age, planID)
	return err
}

// Ler Usuário
func readUser(db *sql.DB, id int) (User, error) {
	var user User
	err := db.QueryRow(`SELECT id, name, email, age, plan_id FROM users WHERE id = $1`, id).Scan(&user.ID, &user.Name, &user.Email, &user.Age, &user.PlanID)
	return user, err
}

// Listar Usuários
func listUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query(`SELECT id, name, email, age, plan_id FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age, &user.PlanID)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// CRUD para Planos

// Criar Plano
func createPlan(db *sql.DB, name string, price float64) error {
	query := `INSERT INTO plans (name, price) VALUES ($1, $2)`
	_, err := db.Exec(query, name, price)
	return err
}

// Listar Planos
func listPlans(db *sql.DB) ([]Plan, error) {
	rows, err := db.Query(`SELECT id, name, price FROM plans`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var plans []Plan
	for rows.Next() {
		var plan Plan
		err := rows.Scan(&plan.ID, &plan.Name, &plan.Price)
		if err != nil {
			return nil, err
		}
		plans = append(plans, plan)
	}
	return plans, nil
}

// CRUD para Pagamentos

// Criar Pagamento
func createPayment(db *sql.DB, userID int, month string, status bool, paymentDate string) error {
	query := `INSERT INTO payments (user_id, month, status, payment_date) VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(query, userID, month, status, paymentDate)
	return err
}

// Listar Pagamentos de um Usuário
func listPaymentsByUser(db *sql.DB, userID int) ([]Payment, error) {
	rows, err := db.Query(`SELECT id, month, status, payment_date FROM payments WHERE user_id = $1`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []Payment
	for rows.Next() {
		var payment Payment
		err := rows.Scan(&payment.ID, &payment.UserID, &payment.Month, &payment.Status, &payment.PaymentDate)
		if err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}
	return payments, nil
}

// Handlers da API

// Handler para criar um usuário
func CreateUser(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var user User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = createUser(db, user.Name, user.Email, user.Age, user.PlanID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// Handler para listar usuários
func ListUsers(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	users, err := listUsers(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Handler para criar um plano
func CreatePlanHandler(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var plan Plan
	err = json.NewDecoder(r.Body).Decode(&plan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = createPlan(db, plan.Name, plan.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// Handler para listar planos
func ListPlansHandler(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	plans, err := listPlans(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(plans)
}

// Handler para criar um pagamento
func CreatePaymentHandler(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var payment Payment
	err = json.NewDecoder(r.Body).Decode(&payment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = createPayment(db, payment.UserID, payment.Month, payment.Status, payment.PaymentDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// Handler para listar pagamentos de um usuário
func ListPaymentsHandler(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	db, err := connectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	payments, err := listPaymentsByUser(db, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payments)
}

// Função para lidar com a requisição para a rota "/"
func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Define o Content-Type como HTML
	w.Header().Set("Content-Type", "text/html")

	// Exibe o HTML diretamente, você pode também usar um template se preferir
	http.ServeFile(w, r, "index.html") // Certifique-se de que o index.html está no diretório correto
}

func main() {

	// Rota para a página inicial
	http.HandleFunc("/", indexHandler)

	// Configurar os manipuladores da API
	http.HandleFunc("/users/create", CreateUser)
	http.HandleFunc("/users", ListUsers)

	http.HandleFunc("/plans/create", CreatePlanHandler)
	http.HandleFunc("/plans", ListPlansHandler)

	http.HandleFunc("/payments/create", CreatePaymentHandler)
	http.HandleFunc("/payments", ListPaymentsHandler)

	// Servir arquivos estáticos (CSS, JS)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Iniciar o servidor na porta 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
