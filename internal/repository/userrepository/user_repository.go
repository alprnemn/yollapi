package userrepository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	cmn "github.com/alprnemn/yollapi/common"
	"github.com/alprnemn/yollapi/internal/domain"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type UserRepository struct {
	Db *sql.DB
}

func (repo *UserRepository) Create(ctx context.Context, user *domain.User) error {

	query := `INSERT INTO users (
    first_name,
    last_name,
    username,
    phone,
    email,
    age,
    password
	) VALUES (
    $1, $2, $3, $4, $5, $6, $7
	)`

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	_, err := repo.Db.ExecContext(ctx,
		query,
		user.FirstName,
		user.LastName,
		user.Username,
		user.Phone,
		user.Email,
		user.Age,
		user.Password,
	)

	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			switch pqErr.Constraint {
			case "users_username_key":
				return cmn.ErrDuplicateUsername
			case "users_email_key":
				return cmn.ErrDuplicateEmail
			case "users_phone_key":
				return cmn.ErrDuplicatePhone
			}
		}
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (repo *UserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {

	query := `
		SELECT  username, first_name, last_name, email, phone
		FROM users
		WHERE email = $1
	`
	user := &domain.User{}
	err := repo.Db.QueryRowContext(ctx, query, email).Scan(
		&user.Username,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Phone,
	)

	if err != nil {
		return nil, err
	}
	return user, nil
}
