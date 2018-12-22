package app

import (
	"github.com/MiteshSharma/project/logger"
	"github.com/MiteshSharma/project/model"
)

func (a *App) CreateUser(user *model.User) error {
	err := a.Repository.User().CreateUser(user)
	if err != nil {
		a.Log.Info("Error creating user object", logger.Error(err))
	}
	return err
}
