package credit_limits

import (
	"context"
	"github.com/golang-test-task/internal/repository"
	"github.com/golang-test-task/internal/repository/storage"
	"github.com/golang-test-task/pkg/databases/postgresql"
)

const MaxCreditLimit = 1000

type CreditDAO struct {
	storage storage.Storage
}

func NewRepositoryCredit(storage storage.Storage) *CreditDAO {
	return &CreditDAO{storage: storage}
}

func (r *CreditDAO) AddCredit(ctx context.Context, username string, amount int) error {

	exec, execErr := r.storage.Client.Exec(
		ctx,
		"UPDATE credit_limits SET credit_limit = credit_limit + $1 WHERE username = $2",
		amount,
		username,
	)
	if execErr != nil {
		execErr = postgresql.ErrDoQuery(execErr)
		return execErr
	}

	if exec.RowsAffected() == 0 {
		return repository.ErrNothingUpdated
	}

	return nil
}

func (r *CreditDAO) SubtractCredit(ctx context.Context, username string, amount int) error {

	exec, err := r.storage.Client.Exec(ctx,
		"UPDATE credit_limits SET credit_limit = credit_limit - $1 WHERE username = $2 AND credit_limit <= $3",
		amount,
		username,
		MaxCreditLimit,
	)

	if err != nil {
		err = postgresql.ErrDoQuery(err)
		return err
	}

	if exec.RowsAffected() == 0 {
		return repository.ErrNothingUpdated
	}

	return nil
}
