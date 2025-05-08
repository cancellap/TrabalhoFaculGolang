package main

import (
	"fmt"
	"log"

	"github.com/cancellap/TrabalhoFaculGolang/config"
	_ "github.com/cancellap/TrabalhoFaculGolang/docs"
	"github.com/cancellap/TrabalhoFaculGolang/routes"
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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Erro ao carregar o arquivo .env")
	}

	err = config.ConnectDB()
	if err != nil {
		log.Fatalf("❌ Erro ao conectar ao banco de dados: %v", err)
	}

	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.SetupRoutes(router)

	fmt.Println("Servidor rodando na porta 8080")
	fmt.Println("Acesse a documentação Swagger em: http://localhost:8080/swagger/index.html")
	router.Run(":8080")
}
