package register_layers

import (
	"github.com/golang-test-task/internal/service/students"
	"github.com/golang-test-task/internal/service/tasks"
)

type GlobalService struct {
	StudentService students.StudentsService
	TaskService    tasks.TaskService
}

func NewGlobalService(repository *GlobalRepository, identity students.IdentityGenerator, clock students.Clock) *GlobalService {
	return &GlobalService{
		StudentService: students.NewStudentsService(repository.StudentRepository, identity, clock),
		TaskService:    tasks.NewTaskService(repository.TaskRepository, identity, clock),
	}
}
