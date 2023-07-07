package register_layers

import (
	"github.com/golang-test-task/internal/repository/storage"
	"github.com/golang-test-task/internal/repository/students"
	"github.com/golang-test-task/internal/repository/tasks"
)

type GlobalRepository struct {
	StudentRepository *students.StudentDAO
	TaskRepository    *tasks.TaskDAO
}

func NewGlobalRepository(storage *storage.Storage) *GlobalRepository {
	return &GlobalRepository{
		StudentRepository: students.NewStudent(storage),
		TaskRepository:    tasks.NewTask(storage),
	}
}
