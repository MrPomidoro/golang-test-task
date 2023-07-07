package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-test-task/internal/domain/task/dto"
	"github.com/golang-test-task/internal/service/tasks"
	"github.com/golang-test-task/pkg/common/api"
	"net/http"
)

type Task struct {
	service tasks.TaskService
}

func NewTask(service tasks.TaskService) *Task {
	return &Task{service: service}
}

// Create task
//
// @Summary Create a task.
// @Description Create a task.
// @Tags Task
// @Accept json
// @Produce json
// @Param task body dto.CreateTaskDTO true "Task object to update"
// @Success 200 {object} api.Success
// @Failure 400 {object} api.Error
// @Failure 500 {object} api.Error
// @Router /tasks/create [post]
func (t *Task) Create(c *gin.Context) {
	var dto dto.CreateTaskDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		api.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	task, err := t.service.Create(c, dto)
	if err != nil {
		api.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	api.SuccessResponse(c, http.StatusOK, "Success create task", task)
}

// Get task
//
// @Summary Get a task.
// @Description Get a task.
// @Tags Task
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Success 200 {object} api.Success
// @Failure 400 {object} api.Error
// @Failure 500 {object} api.Error
// @Router /tasks/{id} [get]
func (t *Task) Get(c *gin.Context) {

	id := c.Param("id")

	task, err := t.service.Get(c, id)
	if err != nil {
		api.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	api.SuccessResponse(c, http.StatusOK, "Success get task", task)
}

// GetAll tasks
//
// @Summary Get all tasks.
// @Description Get all tasks.
// @Tags Task
// @Accept json
// @Produce json
// @Success 200 {object} api.Success
// @Failure 500 {object} api.Error
// @Router /tasks/ [get]
func (t *Task) GetAll(c *gin.Context) {

	tasks, err := t.service.GetAll(c)
	if err != nil {
		api.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	api.SuccessResponse(c, http.StatusOK, "Success get all tasks", tasks)
}
