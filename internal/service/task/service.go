package task

import (
	"context"
	"github.com/google/uuid"

	taskdomain "TrabalhoFaculGolang/internal/domain/task"
)

type Repository interface {
	Create(ctx context.Context, t *taskdomain.Task) error
	List(ctx context.Context) ([]taskdomain.Task, error)
	UpdateStatus(ctx context.Context, id string, completed bool) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateTask(ctx context.Context, t *taskdomain.Task) (*taskdomain.Task, error) {

	task := &taskdomain.Task{
		ID:        uuid.NewString(),
		Title:     t.Title,
		Completed: t.Completed,
	}
	err := s.repo.Create(ctx, task)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (s *Service) ListTasks(ctx context.Context) ([]taskdomain.Task, error) {
	return s.repo.List(ctx)
}

func (s *Service) UpdateTaskStatus(ctx context.Context, id string, completed bool) error {
	return s.repo.UpdateStatus(ctx, id, completed)
}
