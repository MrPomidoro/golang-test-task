package students

import (
	"context"
	"github.com/golang-test-task/internal/repository"
	"github.com/golang-test-task/internal/repository/storage"
	"github.com/golang-test-task/pkg/databases/postgresql"
	"github.com/google/uuid"
)

type StudentDAO struct {
	storage *storage.Storage
}

func NewStudent(storage *storage.Storage) *StudentDAO {
	return &StudentDAO{storage: storage}
}

func (s *StudentDAO) Create(ctx context.Context, st RepositoryStudent) error {

	sql, args, err := s.storage.QueryBuilder.
		Insert(repository.StudentTable).
		Columns(
			"id",
			"full_name",
			"group_num",
			"email",
			"username",
			"create_at",
		).
		Values(
			st.ID,
			st.FullName,
			st.GroupNum,
			st.Username,
			st.CreateAT,
		).ToSql()
	if err != nil {
		err = postgresql.ErrCreateQuery(err)
		return err
	}

	exec, execErr := s.storage.Client.Exec(ctx, sql, args)
	if execErr != nil {
		execErr = postgresql.ErrDoQuery(execErr)
		return execErr
	}

	if exec.RowsAffected() == 0 {
		return repository.ErrNothingInserted
	}

	return s.createDefaultCashAccount(ctx, st.ID)
}

func (s *StudentDAO) createDefaultCashAccount(ctx context.Context, studentID uuid.UUID) error {
	sql, args, err := s.storage.QueryBuilder.
		Insert(repository.CreditLimitTable).Columns("student_id").
		Values(studentID).ToSql()
	if err != nil {
		err = postgresql.ErrCreateQuery(err)
		return err
	}

	exec, execErr := s.storage.Client.Exec(ctx, sql, args)
	if execErr != nil {
		execErr = postgresql.ErrDoQuery(execErr)
		return execErr
	}

	if exec.RowsAffected() == 0 {
		return repository.ErrNothingInserted
	}

	return nil
}

func (s *StudentDAO) Get(ctx context.Context, id uuid.UUID) (*RepositoryStudent, error) {
	var st *RepositoryStudent

	sql, args, err := s.storage.QueryBuilder.
		Select(
			"id",
			"full_name",
			"group_num",
			"email",
			"username",
			"verify_email",
			"create_at",
			"update_at",
		).
		From(repository.StudentTable).
		Where("id = $1", id).ToSql()

	if err != nil {
		err = postgresql.ErrCreateQuery(err)
		return nil, err
	}

	row := s.storage.Client.QueryRow(ctx, sql, args...)
	if err := row.Scan(&st.ID, &st.FullName, &st.GroupNum, &st.Email, &st.Username, &st.VerifyEmail, &st.CreateAT, &st.UpdateAT); err != nil {
		err = postgresql.ErrScan(err)
		return nil, err
	}

	return st, repository.ErrNotFound
}

func (s *StudentDAO) GetAll(ctx context.Context) ([]*RepositoryStudent, error) {
	var sts []*RepositoryStudent
	sql, args, err := s.storage.QueryBuilder.
		Select(
			"id",
			"full_name",
			"group_num",
			"email",
			"username",
			"verify_email",
			"create_at",
			"update_at",
		).
		From(repository.StudentTable).ToSql()

	if err != nil {
		err = postgresql.ErrCreateQuery(err)
		return nil, err
	}

	rows, err := s.storage.Client.Query(ctx, sql, args...)
	if err != nil {
		err = postgresql.ErrDoQuery(err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var st RepositoryStudent
		if err := rows.Scan(&st.ID, &st.FullName, &st.GroupNum, &st.Email, &st.Username, &st.VerifyEmail, &st.CreateAT, &st.UpdateAT); err != nil {
			err = postgresql.ErrScan(err)
			return nil, err
		}
		sts = append(sts, &st)
	}

	return sts, nil
}

func (s *StudentDAO) Delete(ctx context.Context, id uuid.UUID) error {
	sql, args, err := s.storage.QueryBuilder.
		Delete(repository.StudentTable).
		Where("id = $1", id).ToSql()
	if err != nil {
		err = postgresql.ErrCreateQuery(err)
		return err
	}

	exec, execErr := s.storage.Client.Exec(ctx, sql, args)
	if execErr != nil {
		execErr = postgresql.ErrDoQuery(execErr)
		return execErr
	}

	if exec.RowsAffected() == 0 {
		return repository.ErrNothingDeleted
	}

	return nil
}

func (s *StudentDAO) Update(ctx context.Context, st RepositoryStudent) error {
	sql, args, err := s.storage.QueryBuilder.
		Update(repository.StudentTable).
		Set("full_name", st.FullName).
		Set("group_num", st.GroupNum).
		Set("email", st.Email).
		Set("username", st.Username).
		Set("verify_email", st.VerifyEmail).
		Where("id = $1", st.ID).ToSql()
	if err != nil {
		err = postgresql.ErrCreateQuery(err)
		return err
	}

	exec, execErr := s.storage.Client.Exec(ctx, sql, args)
	if execErr != nil {
		execErr = postgresql.ErrDoQuery(execErr)
		return execErr
	}

	if exec.RowsAffected() == 0 {
		return repository.ErrNothingUpdated
	}

	return nil
}
