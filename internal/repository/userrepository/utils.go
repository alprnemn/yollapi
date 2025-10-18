package userrepository

import (
	"database/sql"

	"github.com/alprnemn/yollapi/internal/models"
)

func extractUsersFromRows(rows *sql.Rows) ([]models.User, error) {
	var users []models.User
	var username string
	var firstName string
	var lastName string
	var email string
	var phone string

	for rows.Next() {
		if err := rows.Scan(&username, &firstName, &lastName, &email, &phone); err != nil {
			return nil, err
		}
		users = append(users, models.User{
			Username:  username,
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
			Phone:     phone,
		})
	}
	return users, nil
}
