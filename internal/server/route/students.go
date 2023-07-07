package route

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/golang-test-task/internal/delivery/http/v1"
)

func Students(r *gin.RouterGroup, controller v1.Student) {
	student := r.Group("/students")
	student.POST("/create", controller.Create)
}
