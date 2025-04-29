package domain

import (
	"context"
	"database/sql"

	"github.com/mdwiastika/the-cows-reog-fiber/dto"
)

type User struct {
	ID        string       `db:"id"`
	Name      string       `db:"name"`
	Email     string       `db:"email"`
	Password  string       `db:"password"`
	GoogleID  string       `db:"google_id"`
	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

type UserRepository interface {
	FindAll(ctx context.Context) ([]User, error)
	FindById(ctx context.Context, id string) (User, error)
	Save(ctx context.Context, c *User) error
	Update(ctx context.Context, c *User) error
	Delete(ctx context.Context, id string) error
}

type UserService interface {
	Index(ctx context.Context) ([]dto.UserData, error)
}
