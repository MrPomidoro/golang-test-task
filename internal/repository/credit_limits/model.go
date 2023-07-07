package credit_limits

import (
	"database/sql"
	"github.com/golang-test-task/internal/domain/credit/model"
	"github.com/golang-test-task/pkg/utils"
	"github.com/google/uuid"
	"time"
)

type RepositoryCredit struct {
	ID          uuid.UUID    `json:"id"`
	StudentID   uuid.UUID    `json:"student_id"`
	CreditLimit int          `json:"credit_limit"`
	CreatedAT   time.Time    `json:"create_at"`
	UpdateAT    sql.NullTime `json:"update_at"`
}

func (s *RepositoryCredit) ToDomain() model.Credit {
	var updateAT *time.Time
	if s.UpdateAT.Valid {
		updateAT = utils.Pointer(s.UpdateAT.Time)
	}
	return model.NewCreditModel(s.ID, s.StudentID, s.CreditLimit, s.CreatedAT, updateAT)
}
