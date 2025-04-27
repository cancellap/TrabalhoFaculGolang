package main

import (
	"fmt"
	"log"

	"github.com/cancellap/TrabalhoFaculGolang/config"
	_ "github.com/cancellap/TrabalhoFaculGolang/docs" // Importa a documentação gerada
	"github.com/cancellap/TrabalhoFaculGolang/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Todo List API
// @version 1.0
// @description Esta API gerencia uma lista de tarefas.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email seu-email@example.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /
func main() {
	// Carregar variáveis de ambiente
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Erro ao carregar o arquivo .env")
	}

	// Conectar ao banco de dados
	err = config.ConnectDB()
	if err != nil {
		log.Fatalf("❌ Erro ao conectar ao banco de dados: %v", err)
	}

	// Inicializar o roteador Gin
	r := gin.Default()

	// Adicionar rota do Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Definir a rota para criar tarefas
	r.POST("/tasks", handlers.CreateTask)

	// Iniciar o servidor na porta 8080
	fmt.Println("Servidor rodando na porta 8080")
	fmt.Println("Acesse a documentação Swagger em: http://localhost:8080/swagger/index.html")
	r.Run(":8080")
}
