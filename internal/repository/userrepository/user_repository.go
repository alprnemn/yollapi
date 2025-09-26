package userrepository

import (
	"context"
	"database/sql"
	"github.com/alprnemn/yollapi/internal/domain"
	"log"
	"time"
)

type UserRepository struct {
	Db *sql.DB
}

func (repo *UserRepository) Create(ctx context.Context, user *domain.User) error {
	log.Print("from database")
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

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
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
		return err
	}

	return nil
}
