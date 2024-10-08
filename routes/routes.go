package routes

import (
	"gym-management/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Configurar arquivos estáticos
	r.Static("/static", "./static")

	// Carregar templates
	r.LoadHTMLGlob("templates/*")

	// Rotas da API
	api := r.Group("/api")
	{
		api.POST("/users", controllers.CreateUser)
		api.GET("/users", controllers.GetUsers)
		api.GET("/users/:id", controllers.GetUser)
		api.PUT("/users/:id", controllers.UpdateUser)
		api.DELETE("/users/:id", controllers.DeleteUser)
	}

	// Rota principal para servir a página frontend
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	return r
}
