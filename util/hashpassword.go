package util

import (
	"net/http"

	"github.com/MiteshSharma/project/model"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string, salt string) (string, *model.AppError) {
	passwordWithSalt := password + salt
	bytes, err := bcrypt.GenerateFromPassword([]byte(passwordWithSalt), 14)
	if err != nil {
		return "", model.NewAppError(err.Error(), http.StatusInternalServerError)
	}
	return string(bytes), nil
}

func CheckPasswordHash(password, salt, bcryptPassword string) bool {
	passwordWithSalt := password + salt
	err := bcrypt.CompareHashAndPassword([]byte(bcryptPassword), []byte(passwordWithSalt))
	return err == nil
}
