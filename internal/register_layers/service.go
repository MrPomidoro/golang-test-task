package register_layers

import (
	"github.com/golang-test-task/internal/service/students"
)

type GlobalService struct {
	StudentService students.StudentsService
}

func NewGlobalService(repository *GlobalRepository, identity students.IdentityGenerator, clock students.Clock) *GlobalService {
	return &GlobalService{
		StudentService: students.NewStudentsService(repository.StudentRepository, identity, clock),
	}
}
