package models

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func CreateTask(db *pgxpool.Pool, task *Task) error {
	query := "INSERT INTO tasks (title, completed) VALUES ($1, $2) RETURNING id"

	ctx := context.Background()

	err := db.QueryRow(ctx, query, task.Title, task.Completed).Scan(&task.ID)
	if err != nil {
		return fmt.Errorf("erro ao inserir task no banco: %w", err)
	}

	return nil
}
