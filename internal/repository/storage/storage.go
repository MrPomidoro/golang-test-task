package storage

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/golang-test-task/pkg/databases/postgresql"
)

type Storage struct {
	Client       postgresql.Client
	QueryBuilder sq.StatementBuilderType
}

func NewStorage(client postgresql.Client) *Storage {
	return &Storage{
		Client:       client,
		QueryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}
