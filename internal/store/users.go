package store

import (
	"context"
	"database/sql"
	"github.com/alprnemn/yollapi/internal/types"
)

type IUserStore interface {
	Create(ctx context.Context, user *types.User) error
}

type UserStore struct {
	db *sql.DB
}

func (store *UserStore) Create(ctx context.Context, user *types.User) error {

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

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := store.db.ExecContext(ctx,
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
		return err
	}

	return nil
}
