package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-test-task/internal/service/job"
	"github.com/golang-test-task/pkg/common/api"
	"net/http"
	"strconv"
)

type Job struct {
	service job.Service
}

func NewJob(service job.Service) *Job {
	return &Job{service: service}
}

// GetSlowTaskMissingNumbers
//
// @Summary Get slow task missing numbers
// @Description Get slow task missing numbers
// @Tags Job
// @Accept json
// @Produce json
// @Param slow_task body GetSlowTask true "Get slow task missing numbers"
// @Success 200 {object} api.Success
// @Failure 400 {object} api.Error
// @Failure 500 {object} api.Error
// @Router /jobs/slow_task_missing_numbers [post]
func (j *Job) GetSlowTaskMissingNumbers(c *gin.Context) {

	var body GetSlowTask

	if err := c.ShouldBindJSON(&body); err != nil {
		api.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if err := j.service.SubtractCredit(c, body.Username, body.Amount); err != nil {
		api.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if err := j.service.SlowTaskMissingNumbers(c, body.Nums, body.Username); err != nil {
		api.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	api.SuccessResponse(c, http.StatusOK, "Success get slow task missing numbers", nil)
}

type GetSlowTask struct {
	Nums     []int
	Username string
	Amount   int
}

// AddCredit
//
// @Summary Add credit
// @Description Add credit
// @Tags Job
// @Accept json
// @Produce json
// @Param amount query int true "Amount"
// @Param username query string true "Username"
// @Success 200 {object} api.Success
// @Failure 400 {object} api.Error
// @Failure 500 {object} api.Error
// @Router /jobs/add_credit [post]
func (j *Job) AddCredit(c *gin.Context) {

	amount := c.Query("amount")
	if amount == "" {
		api.ErrorResponse(c, http.StatusBadRequest, fmt.Errorf("amount is empty"))
		return
	}

	amountInt, err := strconv.Atoi(amount)
	if err != nil {
		api.ErrorResponse(c, http.StatusBadRequest, fmt.Errorf("amount is not int"))
		return
	}

	if err := j.service.AddCredit(c, c.Query("username"), amountInt); err != nil {
		api.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	api.SuccessResponse(c, http.StatusOK, "Success add credit", nil)
}
