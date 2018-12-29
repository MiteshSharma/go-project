package app

import (
	"fmt"
	"net/http"

	"github.com/MiteshSharma/project/model"
	"github.com/MiteshSharma/project/util"
)

func (a *App) CreateUser(user *model.User) (*model.UserAuth, *model.AppError) {
	user.Salt = util.RandStringBytes(6)
	hashedPassword, err := util.HashPassword(user.Password, user.Salt)
	if err != nil {
		a.Log.Error(fmt.Sprintf("Hashing password returned error for userId %d", user.UserID))
	}
	user.Password = hashedPassword

	storageResult := a.Repository.User().CreateUser(user)
	if storageResult.Err != nil {
		return nil, storageResult.Err
	}

	user = storageResult.Data.(*model.User)

	userDetail := &model.UserDetail{
		UserID: user.UserID,
	}
	storageResult = a.Repository.User().CreateUserDetail(userDetail)

	if storageResult.Err != nil {
		return nil, storageResult.Err
	}

	roles := []model.Role{model.ADMIN}
	userRole := &model.UserRole{
		UserID: user.UserID,
		Role:   model.ADMIN,
	}

	storageResult = a.Repository.User().AttachRole(userRole)

	token, err := a.SignToken(user.UserID, roles)
	if err != nil {
		return nil, err
	}
	userAuth := &model.UserAuth{
		User:  user,
		Token: token,
	}
	return userAuth, nil
}

func (a *App) UpdateUser(user *model.User) (*model.User, *model.AppError) {
	storageResult := a.Repository.User().GetUser(user.UserID)
	if storageResult.Err != nil {
		return nil, storageResult.Err
	}
	existingUser := storageResult.Data.(*model.User)
	if user.FirstName != "" {
		existingUser.FirstName = user.FirstName
	}
	if user.LastName != "" {
		existingUser.LastName = user.LastName
	}

	storageResult = a.Repository.User().UpdateUser(existingUser)
	if storageResult.Err != nil {
		return nil, storageResult.Err
	}

	user = storageResult.Data.(*model.User)
	return user, nil
}

func (a *App) GetUser(userID int) (*model.User, *model.AppError) {
	storageResult := a.Repository.User().GetUser(userID)
	if storageResult.Err != nil {
		return nil, storageResult.Err
	}

	user := storageResult.Data.(*model.User)
	return user, nil
}

func (a *App) GetAllUser() ([]*model.User, *model.AppError) {
	storageResult := a.Repository.User().GetAllUsers()
	if storageResult.Err != nil {
		return nil, storageResult.Err
	}

	users := storageResult.Data.([]*model.User)
	return users, nil
}

func (a *App) DeleteUser(userID int) (*model.User, *model.AppError) {
	storageResult := a.Repository.User().DeleteUser(userID)
	if storageResult.Err != nil {
		return nil, storageResult.Err
	}

	return nil, nil
}

func (a *App) UpdateUserDetail(userDetail *model.UserDetail) (*model.UserDetail, *model.AppError) {
	storageResult := a.Repository.User().GetUserDetail(userDetail.UserID)
	if storageResult.Err != nil {
		return nil, storageResult.Err
	}
	dbUserDetail := storageResult.Data.(*model.UserDetail)
	userDetail.UserDetailID = dbUserDetail.UserDetailID
	storageResult = a.Repository.User().UpdateUserDetail(userDetail)
	if storageResult.Err != nil {
		return nil, storageResult.Err
	}

	userDetail = storageResult.Data.(*model.UserDetail)
	return userDetail, nil
}

// UserLogin function
func (a *App) UserLogin(user *model.User) (*model.UserAuth, *model.AppError) {
	storageResult := a.Repository.User().GetUserByEmail(user.Email)
	if storageResult.Err != nil {
		return nil, storageResult.Err
	}
	dbUser := storageResult.Data.(*model.User)
	if util.CheckPasswordHash(user.Password, dbUser.Salt, dbUser.Password) {
		storageResult = a.Repository.User().GetRoles(dbUser.UserID)
		if storageResult.Err != nil {
			return nil, storageResult.Err
		}
		roles := storageResult.Data.([]model.UserRole)
		var rolesObj []model.Role
		for _, role := range roles {
			rolesObj = append(rolesObj, role.Role)
		}
		token, err := a.SignToken(dbUser.UserID, rolesObj)
		if err != nil {
			return nil, err
		}
		userAuth := &model.UserAuth{
			User:  dbUser,
			Token: token,
		}
		return userAuth, nil
	}
	return nil, model.NewAppError("User not exist", http.StatusNotFound)
}

// UserLogout function
func (a *App) UserLogout(userID int) {
	a.ResetToken(userID)
}
