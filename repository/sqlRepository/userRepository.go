package sqlRepository

import (
	"net/http"

	"github.com/MiteshSharma/project/model"
)

type UserRepository struct {
	*SqlRepository
}

func NewSqlUserRepository(sqlRepository *SqlRepository) UserRepository {
	userRepository := UserRepository{sqlRepository}

	hasTable := userRepository.DB.HasTable(&model.User{})
	if !hasTable {
		userRepository.DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.User{})
		userRepository.DB.Model(&model.User{}).AddIndex("idx_email", "email")
	}
	userRepository.DB.Model(&model.User{}).AddIndex("idx_email", "email")
	if (!userRepository.DB.HasTable(&model.UserDetail{})) {
		// will append "ENGINE=InnoDB" to the SQL statement when creating table `users`
		userRepository.DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.UserDetail{})
	}
	if (!userRepository.DB.HasTable(&model.UserRole{})) {
		// will append "ENGINE=InnoDB" to the SQL statement when creating table `users`
		userRepository.DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.UserRole{})
	}
	return userRepository
}

// CreateUser func is used to create user object in db
func (ur UserRepository) CreateUser(user *model.User) *model.StorageResult {
	if err := ur.DB.Create(&user).Error; err != nil {
		return model.NewStorageResult(nil, model.NewAppError(err.Error(), http.StatusInternalServerError))
	}
	return model.NewStorageResult(user, nil)
}

func (ur UserRepository) UpdateUser(user *model.User) *model.StorageResult {
	if err := ur.DB.Save(&user).Error; err != nil {
		return model.NewStorageResult(nil, model.NewAppError(err.Error(), http.StatusInternalServerError))
	}
	return model.NewStorageResult(user, nil)
}

func (ur UserRepository) GetUser(userID int) *model.StorageResult {
	var user model.User
	if err := ur.DB.First(&user, userID).Error; err != nil {
		return model.NewStorageResult(nil, model.NewAppError(err.Error(), http.StatusInternalServerError))
	}
	return model.NewStorageResult(user, nil)
}

func (ur UserRepository) GetAllUsers() *model.StorageResult {
	var users []model.User
	if err := ur.DB.First(&users).Error; err != nil {
		return model.NewStorageResult(nil, model.NewAppError(err.Error(), http.StatusInternalServerError))
	}
	return model.NewStorageResult(users, nil)
}

func (ur UserRepository) GetUserByEmail(email string) *model.StorageResult {
	var user model.User
	if err := ur.DB.Where("email = ?", email).Find(&user).Error; err != nil {
		return model.NewStorageResult(nil, model.NewAppError(err.Error(), http.StatusInternalServerError))
	}
	return model.NewStorageResult(user, nil)
}

func (ur UserRepository) DeleteUser(userID int) *model.StorageResult {
	if err := ur.DB.Where("user_id = ?", userID).Delete(model.User{}).Error; err != nil {
		return model.NewStorageResult(nil, model.NewAppError(err.Error(), http.StatusInternalServerError))
	}
	return model.NewStorageResult(nil, nil)
}

func (ur UserRepository) CreateUserDetail(userDetail *model.UserDetail) *model.StorageResult {
	if err := ur.DB.Create(&userDetail).Error; err != nil {
		return model.NewStorageResult(nil, model.NewAppError(err.Error(), http.StatusInternalServerError))
	}
	return model.NewStorageResult(userDetail, nil)
}

func (ur UserRepository) UpdateUserDetail(userDetail *model.UserDetail) *model.StorageResult {
	if err := ur.DB.Save(&userDetail).Error; err != nil {
		return model.NewStorageResult(nil, model.NewAppError(err.Error(), http.StatusInternalServerError))
	}
	return model.NewStorageResult(userDetail, nil)
}

func (ur UserRepository) GetRoles(userID int) *model.StorageResult {
	var userRoles []model.UserRole
	if err := ur.DB.Where("user_id = ?", userID).Find(&userRoles).Error; err != nil {
		return model.NewStorageResult(nil, model.NewAppError(err.Error(), http.StatusInternalServerError))
	}
	return model.NewStorageResult(userRoles, nil)
}

func (ur UserRepository) AttachRole(userRole *model.UserRole) *model.StorageResult {
	if err := ur.DB.Create(&userRole).Error; err != nil {
		return model.NewStorageResult(nil, model.NewAppError(err.Error(), http.StatusInternalServerError))
	}
	return model.NewStorageResult(userRole, nil)
}
