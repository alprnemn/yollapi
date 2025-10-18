package models

import (
	"github.com/alprnemn/yollapi/common"
	"golang.org/x/crypto/bcrypt"
)

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
