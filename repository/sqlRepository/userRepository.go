package sqlRepository

import "github.com/MiteshSharma/project/model"

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
	return userRepository
}

func (ur UserRepository) CreateUser(user *model.User) error {
	if err := ur.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}
