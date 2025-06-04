package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB() error {
	databaseUrl, ok := os.LookupEnv("DATABASE_URL")
	if !ok || databaseUrl == "" {
		return fmt.Errorf("DATABASE_URL não está definida")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, databaseUrl)
	if err != nil {
		return fmt.Errorf("erro ao criar pool: %w", err)
	}

	if err = pool.Ping(ctx); err != nil {
		return fmt.Errorf("falha ao conectar: %w", err)
	}

	DB = pool
	fmt.Println("✅ Banco conectado!")
	return nil
}
