package main

import (
	"gym-management/config"
	"gym-management/models"
	"gym-management/routes"
	"fmt"
)

func main() {
	// Conectar ao banco de dados
	Config.Connect()

	// Migrar o modelo User
	Config.DB.AutoMigrate(&models.User{})

	// Iniciar o servidor
	r := routes.SetupRouter()
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
