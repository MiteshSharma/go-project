package repository

import "github.com/MiteshSharma/project/model"

type Repository interface {
	Close() error
	User() UserRepository
}

type UserRepository interface {
	CreateUser(user *model.User)
}
