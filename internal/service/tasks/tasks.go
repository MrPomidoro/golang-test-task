package tasks

import (
	"context"
	"github.com/golang-test-task/internal/domain/task/dto"
	"github.com/golang-test-task/internal/domain/task/model"
	"github.com/golang-test-task/internal/repository/tasks"
	"github.com/golang-test-task/pkg/common/errors"
	"github.com/google/uuid"
	"time"
)

type repository interface {
	Create(ctx context.Context, task tasks.RepositoryTask) error
	Get(ctx context.Context, id string) (tasks.RepositoryTask, error)
	GetAll(ctx context.Context) ([]tasks.RepositoryTask, error)
}

type IdentityGenerator interface {
	GenerateUUIDv4() uuid.UUID
}

type Clock interface {
	Now() time.Time
}

type TaskService struct {
	repository repository
	identity   IdentityGenerator
	clock      Clock
}

func NewTaskService(repository repository, identity IdentityGenerator, clock Clock) TaskService {
	return TaskService{repository: repository, identity: identity, clock: clock}
}

func (s *TaskService) Create(ctx context.Context, dto dto.CreateTaskDTO) (model.Task, error) {

	tr := tasks.RepositoryTask{
		ID:          s.identity.GenerateUUIDv4(),
		Description: dto.Description,
		Cost:        dto.Cost,
		CreateAT:    s.clock.Now(),
	}

	if err := s.repository.Create(ctx, tr); err != nil {
		return model.Task{}, errors.Wrap(err, "Repository.Create")
	}

	std := tr.ToDomain()
	return std, nil
}

func (s *TaskService) Get(ctx context.Context, id string) (model.Task, error) {

	tr, err := s.repository.Get(ctx, id)
	if err != nil {
		return model.Task{}, errors.Wrap(err, "Repository.Get")
	}

	std := tr.ToDomain()
	return std, nil
}

func (s *TaskService) GetAll(ctx context.Context) ([]model.Task, error) {
	trs, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Repository.GetAll")
	}

	var stds []model.Task
	for _, tr := range trs {
		std := tr.ToDomain()
		stds = append(stds, std)
	}
	return stds, nil
}
