package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool // Essa variável será usada para acessar o banco no projeto todo

func ConnectDB() error {
	// 1. Lê a variável de ambiente que contém a URL de conexão com o banco
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		return fmt.Errorf("DATABASE_URL não está definida nas variáveis de ambiente")
	}

	// 2. Cria um contexto com timeout de 5 segundos — se demorar mais, cancela
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 3. Cria o pool de conexões com base na URL
	pool, err := pgxpool.New(ctx, databaseUrl)
	if err != nil {
		return fmt.Errorf("erro ao criar pool de conexão: %w", err)
	}

	// 4. Tenta "pingar" o banco pra garantir que a conexão está funcionando
	err = pool.Ping(ctx)
	if err != nil {
		return fmt.Errorf("não foi possível conectar ao banco: %w", err)
	}

	// 5. Se tudo deu certo, guarda o pool de conexões em uma variável global
	DB = pool
	fmt.Println("✅ Conectado ao banco de dados com sucesso")
	return nil
}
