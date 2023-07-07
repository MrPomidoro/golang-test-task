package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-test-task/internal/domain/students/dto"
	"github.com/golang-test-task/internal/service/students"
	"github.com/golang-test-task/pkg/common/api"
	"net/http"
)

type Student struct {
	service students.Service
}

func NewStudent(service students.Service) *Student {
	return &Student{service: service}
}

// Create student
//
// @Summary Create a user.
// @Description Create a user.
// @ID update-user
// @Tags User
// @Accept json
// @Produce json
// @Param user body dto.CreateStudentDTO true "User object to update"
// @Success 200 {object} api.Success
// @Failure 400 {object} api.Error
// @Failure 500 {object} api.Error
// @Router /students/create [post]
func (s *Student) Create(c *gin.Context) {

	var dto dto.CreateStudentDTO

	if err := c.BindJSON(&dto); err != nil {
		api.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	st, err := s.service.Create(c, dto)
	if err != nil {
		api.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	api.SuccessResponse(c, http.StatusOK, "Success create student", st)
}
