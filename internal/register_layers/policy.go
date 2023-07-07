package register_layers

import (
	"github.com/golang-test-task/internal/domain/policy/students"
	"github.com/golang-test-task/internal/domain/students/service"
)

type GlobalPolicy struct {
	StudentPolicy *students.Policy
}

func NewGlobalPolicy(studentService *service.StudentsService, identity students.IdentityGenerator, clock students.Clock) GlobalPolicy {
	return GlobalPolicy{
		StudentPolicy: students.NewProductPolicy(studentService, identity, clock),
	}
}
