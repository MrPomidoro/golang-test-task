package register_layers

import (
	"github.com/golang-test-task/internal/service/job"
	"github.com/golang-test-task/internal/service/students"
	"github.com/golang-test-task/internal/service/tasks"
)

type GlobalService struct {
	StudentService students.Service
	TaskService    tasks.TaskService
	JobService     job.Service
}

func NewGlobalService(repository *GlobalRepository, identity students.IdentityGenerator, clock students.Clock, email job.EmailService) *GlobalService {
	return &GlobalService{
		StudentService: students.NewStudentsService(repository.StudentRepository, identity, clock),
		TaskService:    tasks.NewTaskService(repository.TaskRepository, identity, clock),
		JobService:     job.NewJobService(repository.CreditRepository, repository.EmailRepository, email),
	}
}
