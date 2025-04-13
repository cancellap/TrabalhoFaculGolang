package main

import (
	"log"

	"github.com/cancellap/TrabalhoFaculGolang/config" // Muda isso pro nome real do seu módulo
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Erro ao carregar o arquivo .env")
	}

	err = config.ConnectDB()
	if err != nil {
		log.Fatalf("❌ Erro ao conectar ao banco de dados: %v", err)
	}

	// Agora você pode usar config.DB para fazer queries
}
