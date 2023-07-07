package register_layers

import (
	v1 "github.com/golang-test-task/internal/delivery/http/v1"
)

type GlobalDelivery struct {
	StudentDelivery *v1.Student
	TaskDelivery    *v1.Task
	JobDelivery     *v1.Job
}

func NewGlobalDelivery(service *GlobalService) *GlobalDelivery {
	return &GlobalDelivery{
		StudentDelivery: v1.NewStudent(service.StudentService),
		TaskDelivery:    v1.NewTask(service.TaskService),
		JobDelivery:     v1.NewJob(service.JobService),
	}
}
