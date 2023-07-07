package route

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/golang-test-task/internal/delivery/http/v1"
)

func Job(r *gin.RouterGroup, controller v1.Job) {
	job := r.Group("/jobs")
	job.POST("/slow_task_missing_numbers", controller.GetSlowTaskMissingNumbers)
	job.POST("/add_credit", controller.AddCredit)
}
