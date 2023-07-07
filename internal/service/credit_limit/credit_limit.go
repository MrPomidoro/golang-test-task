package credit_limit

import "context"

type repository interface {
	AddCredit(ctx context.Context, username string, amount int) error
	SubtractCredit(ctx context.Context, username string, amount int) error
}

type IdentityGenerator interface {
}

type Service struct {
	repository repository
}

func NewService(repository repository) Service {
	return Service{repository: repository}
}

func (s Service) AddCredit(ctx context.Context, username string, amount int) error {
	return s.repository.AddCredit(ctx, username, amount)
}

func (s Service) SubtractCredit(ctx context.Context, username string, amount int) error {
	return s.repository.SubtractCredit(ctx, username, amount)
}
