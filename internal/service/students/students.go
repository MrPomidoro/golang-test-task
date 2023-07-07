package students

import (
	"context"
	"github.com/golang-test-task/internal/domain/students/dto"
	"github.com/golang-test-task/internal/domain/students/model"
	"github.com/golang-test-task/internal/repository/students"
	"github.com/golang-test-task/pkg/common/errors"
	"github.com/google/uuid"
	"time"
)

type repository interface {
	Create(context.Context, students.RepositoryStudent) error
}

type IdentityGenerator interface {
	GenerateUUIDv4() uuid.UUID
}

type Clock interface {
	Now() time.Time
}

type Service struct {
	repository repository
	identity   IdentityGenerator
	clock      Clock
}

func NewStudentsService(repository repository, identity IdentityGenerator, clock Clock) Service {
	return Service{repository: repository, identity: identity, clock: clock}
}

func (s *Service) Create(ctx context.Context, dto dto.CreateStudentDTO) (model.Student, error) {

	rst := students.RepositoryStudent{
		ID:          s.identity.GenerateUUIDv4(),
		FullName:    dto.FullName,
		GroupNum:    dto.GroupNum,
		Email:       dto.Email,
		Username:    dto.Username,
		VerifyEmail: false,
		CreateAT:    s.clock.Now(),
	}

	if err := s.repository.Create(ctx, rst); err != nil {
		return model.Student{}, errors.Wrap(err, "Repository.Create")
	}

	std := rst.ToDomain()

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
