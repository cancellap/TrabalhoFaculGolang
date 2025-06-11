package main

import (
	"fmt"
	"log"
	"os"

	config "TrabalhoFaculGolang/internal/config"
	_ "TrabalhoFaculGolang/docs"
	"TrabalhoFaculGolang/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API TodoList GO
// @version 1.0
// @description API feita em Golang para gerenciamento de tarefas
// @host todolistgo-57814d96dd24.herokuapp.com
// @schemes https
// @BasePath /
func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("❌ Erro ao carregar o arquivo .env")
	}

	config.InitDB()
	if err != nil {
		log.Fatalf("❌ Erro ao conectar ao banco de dados: %v", err)
	}

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.SetupRoutes(router, config.DB)

	fmt.Println("Servidor rodando na porta 8080")
	fmt.Println("Acesse a documentação Swagger em: http://localhost:8080/swagger/index.html")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
