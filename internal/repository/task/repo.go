package task

import (
	"context"
	"fmt"

	taskdomain "TrabalhoFaculGolang/internal/domain/task"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	Create(ctx context.Context, task *taskdomain.Task) error
	List(ctx context.Context) ([]taskdomain.Task, error)
	UpdateStatus(ctx context.Context, id string, completed bool) error
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, task *taskdomain.Task) error {
	query := "INSERT INTO tasks (id, title, completed) VALUES ($1, $2, $3)"
	_, err := r.db.Exec(ctx, query, task.ID, task.Title, task.Completed)
	if err != nil {
		return fmt.Errorf("erro ao inserir task: %w", err)
	}
	return nil
}

func (r *repository) List(ctx context.Context) ([]taskdomain.Task, error) {
	rows, err := r.db.Query(ctx, "SELECT id, title, completed FROM tasks")
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar tasks: %w", err)
	}
	defer rows.Close()

	var tasks []taskdomain.Task
	for rows.Next() {
		var t taskdomain.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Completed); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (r *repository) UpdateStatus(ctx context.Context, id string, completed bool) error {
	_, err := r.db.Exec(ctx, "UPDATE tasks SET completed = $1 WHERE id = $2", completed, id)
	if err != nil {
		return fmt.Errorf("erro ao atualizar status: %w", err)
	}
	return nil
}
