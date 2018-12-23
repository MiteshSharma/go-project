package repository

import "github.com/MiteshSharma/project/model"

type Repository interface {
	Close() error
	User() UserRepository
}

type UserRepository interface {
	CreateUser(user *model.User) *model.StorageResult
	UpdateUser(user *model.User) *model.StorageResult
	GetUser(userID int) *model.StorageResult
	GetAllUsers() *model.StorageResult
	GetUserByEmail(email string) *model.StorageResult
	DeleteUser(userID int) *model.StorageResult
	CreateUserDetail(userDetail *model.UserDetail) *model.StorageResult
	UpdateUserDetail(userDetail *model.UserDetail) *model.StorageResult
	GetRoles(userID int) *model.StorageResult
	AttachRole(userRole *model.UserRole) *model.StorageResult
}
