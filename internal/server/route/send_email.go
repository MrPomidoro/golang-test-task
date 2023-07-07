package route

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/golang-test-task/internal/delivery/http/v1"
)

func Email(r *gin.RouterGroup, controller v1.Email) {
	student := r.Group("/email")
	student.POST("/send", controller.Send)
}
