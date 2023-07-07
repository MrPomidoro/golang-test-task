package api

import "github.com/gin-gonic/gin"

type Success struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SuccessResponse middleware sets the response for successful requests
// It takes in the gin context, the HTTP status code, a message string and some data to be returned
func SuccessResponse(c *gin.Context, status int, message string, data interface{}) {
	// Create a new Success struct with the given message and data
	success := Success{
		Status:  "success",
		Message: message,
		Data:    data,
	}
	// Set the response status code and return the success struct as JSON
	c.JSON(status, success)
}

type Error struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

// ErrorResponse middleware sets the response for failed requests
func ErrorResponse(c *gin.Context, status int, err error) {
	var s string
	switch status {
	case 400:
		s = "Bad Request"
	case 500:
		s = "Internal Server Error"
	}
	// Create a JSON response with the given status code and error message
	c.AbortWithStatusJSON(status, Error{Status: s, Error: err.Error()})
}
