package register_layers

import (
	"github.com/golang-test-task/internal/repository/storage"
	"github.com/golang-test-task/internal/repository/students"
)

type GlobalRepository struct {
	StudentRepository *students.StudentDAO
}

func NewGlobalRepository(storage storage.Storage) *GlobalRepository {
	return &GlobalRepository{
		StudentRepository: students.NewStudent(storage),
	}
}
