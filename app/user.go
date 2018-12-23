package app

import (
	"github.com/MiteshSharma/project/model"
)

func (a *App) CreateUser(user *model.User) (*model.User, *model.AppError) {
	storageResult := a.Repository.User().CreateUser(user)
	return user, storageResult.Err
}
