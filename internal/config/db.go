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
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	if user == "" || password == "" || host == "" || port == "" || dbname == "" {
		log.Fatal("Variáveis de ambiente de conexão não estão totalmente definidas")
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		user, password, host, port, dbname,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var err error
	DB, err = pgxpool.New(ctx, dsn)
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
