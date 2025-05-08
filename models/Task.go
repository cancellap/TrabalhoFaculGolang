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

func GetTasks(db *pgxpool.Pool) ([]Task, error) {
	query := "SELECT * FROM tasks"
	ctx := context.Background()

	rows, err := db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar tasks no banco: %w", err)
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Completed); err != nil {
			return nil, fmt.Errorf("erro ao escanear task: %w", err)
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("erro ao iterar sobre as tasks: %w", err)
	}

	return tasks, nil
}
func UpdateTaskStatus(db *pgxpool.Pool, id string, completed bool) error {
	query := "UPDATE tasks SET completed = $1 WHERE id = $2"
	ctx := context.Background()

	_, err := db.Exec(ctx, query, completed, id)
	if err != nil {
		return fmt.Errorf("erro ao atualizar status da task (ID: %s): %w", id, err)
	}

	return nil
}
