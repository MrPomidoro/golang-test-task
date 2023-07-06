package service

import (
	"context"
	"github.com/golang-test-task/internal/domain/students/model"
	"github.com/golang-test-task/internal/repository/students"
	"github.com/golang-test-task/pkg/common/errors"
)

type repository interface {
	Create(context.Context, students.RepositoryStudent) error
}

type StudentsService struct {
	repository repository
}

func NewStudentsService(repository repository) *StudentsService {
	return &StudentsService{repository: repository}
}

// Create
func (s *StudentsService) Create(ctx context.Context, st students.RepositoryStudent) (model.Student, error) {

	err := s.repository.Create(ctx, st)
	if err != nil {
		return model.Student{}, errors.Wrap(err, "Repository.Create")
	}

	std := st.ToDomain()

	return model.NewStudent(
		std.ID,
		std.FullName,
		std.GroupNum,
		std.Email,
		std.Username,
		std.VerifyEmail,
		std.CreateAT,
		std.UpdateAT,
	), nil
}
