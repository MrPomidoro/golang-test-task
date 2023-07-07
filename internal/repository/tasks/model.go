package tasks

import (
	"database/sql"
	"github.com/golang-test-task/internal/domain/task/model"
	"github.com/golang-test-task/pkg/utils"
	"github.com/google/uuid"
	"time"
)

type RepositoryTask struct {
	ID          uuid.UUID    `json:"id"`
	Description string       `json:"description"`
	Cost        int          `json:"cost"`
	CreateAT    time.Time    `json:"create_at"`
	UpdateAT    sql.NullTime `json:"update_at"`
}

func (s *RepositoryTask) ToDomain() model.Task {
	var updateAT *time.Time
	if s.UpdateAT.Valid {
		updateAT = utils.Pointer(s.UpdateAT.Time)
	}
	return model.Task{
		ID:          s.ID.String(),
		Description: s.Description,
		Cost:        s.Cost,
		CreateAT:    s.CreateAT,
		UpdateAT:    updateAT,
	}
}
