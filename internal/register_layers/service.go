package register_layers

import (
	"github.com/golang-test-task/internal/domain/students/service"
)

type GlobalService struct {
	StudentService *service.StudentsService
}

func NewGlobalService(repository GlobalRepository) GlobalService {
	return GlobalService{
		StudentService: service.NewStudentsService(repository.StudentRepository),
	}
}
