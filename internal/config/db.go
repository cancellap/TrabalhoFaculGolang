package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitDB() {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("A variável de ambiente DATABASE_URL não está definida")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var err error
	DB, err = pgxpool.New(ctx, databaseURL)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	err = DB.Ping(ctx)
	if err != nil {
		log.Fatalf("Erro ao testar conexão com banco: %v", err)
	}

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS tasks (
		id VARCHAR(100) PRIMARY KEY,
		title VARCHAR(100),
		completed BOOLEAN
	);
	`
	_, err = DB.Exec(ctx, createTableSQL)
	if err != nil {
		log.Fatalf("Erro ao criar tabela: %v", err)
	}

	fmt.Println("✅ Conectado ao banco de dados com sucesso e tabela criada (se necessário)")
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
