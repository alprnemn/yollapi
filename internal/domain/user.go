package domain

import (
	"github.com/alprnemn/yollapi/common"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64    `json:"id,omitempty"`
	FirstName string   `json:"firstName,omitempty"`
	LastName  string   `json:"lastName,omitempty"`
	Username  string   `json:"username,omitempty"`
	Phone     string   `json:"phone,omitempty"`
	Email     string   `json:"email,omitempty"`
	Age       uint8    `json:"age,omitempty"`
	Password  password `json:"-"`
	CreatedAt string   `json:"created-at,omitempty"`
}

type password struct {
	Text *string
	Hash []byte
}

func (pw *password) Set(text string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return common.ErrGeneratePassword
	}

	pw.Text = &text
	pw.Hash = hash

	return nil
}

func (pw *password) Compare(text string) error {
	err := bcrypt.CompareHashAndPassword(pw.Hash, []byte(text))
	if err != nil {
		return common.ErrPasswordInvalid
	}
	return nil
}
