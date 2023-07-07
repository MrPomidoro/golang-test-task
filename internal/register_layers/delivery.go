package register_layers

import (
	v1 "github.com/golang-test-task/internal/delivery/http/v1"
	"github.com/golang-test-task/internal/domain/policy/students"
)

type GlobalDelivery struct {
	StudentDelivery *v1.Student
}

func NewGlobalDelivery(policy students.Policy) GlobalDelivery {
	return GlobalDelivery{
		StudentDelivery: v1.NewStudent(policy),
	}
}
