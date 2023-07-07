package server

import (
	"github.com/golang-test-task/pkg/validate"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func New() *gin.Engine {
	app := gin.Default()
	if os.Getenv("GIN_MODE") == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	binding.Validator = new(validate.DefaultValidator)

	return app
}

// Redirect the user to the API documentation page using a 301 redirect
func Redirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/api/v1/swagger/index.html")
}
