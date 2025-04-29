package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/mdwiastika/the-cows-reog-fiber/domain"
)

type UserRepository struct {
	db *goqu.Database
}

func NewUser(con *sql.DB) domain.UserRepository {
	return &UserRepository{
		db: goqu.New("default", con),
	}
}

func (ur *UserRepository) FindAll(ctx context.Context) (result []domain.User, err error) {
	dataset := ur.db.From("users").Where(goqu.C("deleted_at").IsNull())
	err = dataset.ScanStructsContext(ctx, &result)
	return
}

func (ur *UserRepository) FindById(ctx context.Context, id string) (result domain.User, err error) {
	dataset := ur.db.From("users").
		Where(goqu.C("deleted_at").IsNull(), goqu.C("id").Eq(id))

	_, err = dataset.ScanStructContext(ctx, &result)
	return
}

func (ur *UserRepository) Save(ctx context.Context, c *domain.User) error {
	executor := ur.db.Insert("users").Rows(c).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (ur *UserRepository) Update(ctx context.Context, c *domain.User) error {
	executor := ur.db.Update("users").Where(goqu.C("id").Eq(c.ID)).Set(c).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (ur *UserRepository) Delete(ctx context.Context, id string) error {
	executor := ur.db.Update("users").Where(goqu.C("id").Eq(id)).Set(goqu.Record{"deleted_at": sql.NullTime{Valid: true, Time: time.Now()}}).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}
