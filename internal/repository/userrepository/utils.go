package userrepository

import (
	"database/sql"

	"github.com/alprnemn/yollapi/internal/domain"
)

func extractUsersFromRows(rows *sql.Rows) ([]domain.User, error) {
	var users []domain.User
	var username string
	var firstName string
	var lastName string
	var email string
	var phone string

	for rows.Next() {
		if err := rows.Scan(&username, &firstName, &lastName, &email, &phone); err != nil {
			return nil, err
		}
		users = append(users, domain.User{
			Username:  username,
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
			Phone:     phone,
		})
	}
	return users, nil
}
