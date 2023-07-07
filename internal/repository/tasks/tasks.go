package tasks

import (
	"context"
	"fmt"
	"github.com/golang-test-task/internal/repository"
	"github.com/golang-test-task/internal/repository/storage"
	"github.com/golang-test-task/pkg/common/logging"
	"github.com/golang-test-task/pkg/databases/postgresql"
)

type TaskDAO struct {
	storage *storage.Storage
}

func NewTask(storage *storage.Storage) *TaskDAO {
	return &TaskDAO{storage: storage}
}

func (r *TaskDAO) Create(ctx context.Context, task RepositoryTask) error {

	logging.WithFields(
		ctx,
		logging.StringField("id", task.ID.String()),
		logging.StringField("description", task.Description),
		logging.StringField("cost", string(task.Cost)),
		logging.StringField("create_at", task.CreateAT.String()),
	).Info("create student")

	sql, args, err := r.storage.QueryBuilder.
		Insert(repository.TasksTable).
		Columns(
			"id",
			"description",
			"cost",
			"create_at",
		).Values(
		task.ID.String(),
		task.Description,
		task.Cost,
		task.CreateAT,
	).ToSql()
	if err != nil {
		err = postgresql.ErrCreateQuery(err)
		return err
	}

	logging.WithFields(ctx, logging.StringField("sql", sql)).Info("create task")

	logging.L(ctx).Info(fmt.Sprintf("args %v", args))

	exec, execErr := r.storage.Client.Exec(ctx, sql, args...)
	if execErr != nil {
		err = postgresql.ErrCreateQuery(execErr)
		return err
	}

	if exec.RowsAffected() != 0 {
		return repository.ErrNothingInserted
	}

	return nil

}

func (r *TaskDAO) Get(ctx context.Context, id string) (RepositoryTask, error) {

	logging.WithFields(
		ctx,
		logging.StringField("id", id),
	).Info("get task")

	sql, args, err := r.storage.QueryBuilder.Select("id", "description", "cost", "create_at").
		From(repository.TasksTable).
		Where("id = ?", id).ToSql()

	if err != nil {
		err = postgresql.ErrCreateQuery(err)
		return RepositoryTask{}, err
	}

	logging.WithFields(ctx, logging.StringField("sql", sql)).Info("get task")

	logging.L(ctx).Info(fmt.Sprintf("args %v", args))

	row := r.storage.Client.QueryRow(ctx, sql, args...)

	rt := RepositoryTask{}

	if err := row.Scan(&rt.ID, &rt.Description, &rt.Cost, &rt.CreateAT); err != nil {
		err = postgresql.ErrScan(err)
		return RepositoryTask{}, err
	}

	return rt, nil

}

func (r *TaskDAO) GetAll(ctx context.Context) ([]RepositoryTask, error) {
	sql, args, err := r.storage.QueryBuilder.Select("id", "description", "cost", "create_at").From(repository.TasksTable).ToSql()

	if err != nil {
		err = postgresql.ErrCreateQuery(err)
		return []RepositoryTask{}, err
	}

	logging.WithFields(ctx, logging.StringField("sql", sql)).Info("get all tasks")

	logging.L(ctx).Info(fmt.Sprintf("args %v", args))

	rows, err := r.storage.Client.Query(ctx, sql, args...)
	if err != nil {
		err = postgresql.ErrDoQuery(err)
		return []RepositoryTask{}, err
	}

	defer rows.Close()

	var tasks []RepositoryTask

	for rows.Next() {
		var rt RepositoryTask
		if err := rows.Scan(&rt.ID, &rt.Description, &rt.Cost, &rt.CreateAT); err != nil {
			err = postgresql.ErrScan(err)
			return []RepositoryTask{}, err
		}
		logging.WithFields(ctx,
			logging.StringField("id", rt.ID.String()),
			logging.StringField("description", rt.Description),
			logging.StringField("cost", string(rt.Cost)),
		).Info("get all tasks")
		tasks = append(tasks, rt)
	}

	return tasks, nil
}
