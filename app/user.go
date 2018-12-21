package app

import "github.com/MiteshSharma/project/model"

func (a *App) CreateUser(user *model.User) {
	a.Repository.User().CreateUser(user)
}
