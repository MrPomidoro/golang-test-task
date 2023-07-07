package route

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/golang-test-task/internal/delivery/http/v1"
)

func Tasks(r *gin.RouterGroup, controller v1.Task) {
	task := r.Group("/tasks")

	task.POST("/create", controller.Create)
	task.GET("/:id", controller.Get)
	task.GET("/", controller.GetAll)
}
