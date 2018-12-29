package redisRepository

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MiteshSharma/project/repository/sqlRepository"

	"github.com/MiteshSharma/project/logger"
	"github.com/MiteshSharma/project/model"
)

type UserRepository struct {
	sqlRepository.UserRepository
	*RedisRepository
}

func NewUserRepository(redisRepository *RedisRepository) UserRepository {
	userRedisRepository := UserRepository{
		RedisRepository: redisRepository,
		UserRepository:  redisRepository.SQLRepository.UserRepository,
	}

	return userRedisRepository
}

// SaveSession func is used to save user session object in db
func (ur UserRepository) CreateSession(session *model.UserSession) *model.StorageResult {
	userAuthKey := fmt.Sprintf("user:%d", session.UserID)
	err := ur.Redis.Set(userAuthKey, session.ToJson(), 0).Err()
	if err != nil {
		ur.Log.Error("Error writing redis for user login token", logger.Error(err))
	}
	return ur.UserRepository.CreateSession(session)
}

func (ur UserRepository) UpdateSession(session *model.UserSession) *model.StorageResult {
	userAuthKey := fmt.Sprintf("user:%d", session.UserID)
	err := ur.Redis.Set(userAuthKey, session.ToJson(), 0).Err()
	if err != nil {
		ur.Log.Error("Error writing redis for user login token", logger.Error(err))
	}
	return ur.UserRepository.UpdateSession(session)
}

func (ur UserRepository) GetSession(userID int) *model.StorageResult {
	_, err := ur.Redis.Ping().Result()
	if err != nil {
		return ur.UserRepository.GetSession(userID)
	}
	userAuthKey := fmt.Sprintf("user:%d", userID)
	result, err := ur.Redis.Get(userAuthKey).Result()
	if err != nil {
		return ur.UserRepository.GetSession(userID)
	}
	if result == "" {
		return ur.UserRepository.GetSession(userID)
	}
	var session model.UserSession
	err = json.Unmarshal([]byte(result), &session)
	if err != nil {
		return model.NewStorageResult(nil, model.NewAppError(err.Error(), http.StatusInternalServerError))
	}
	return model.NewStorageResult(&session, nil)
}

func (ur UserRepository) DeleteSession(userID int) *model.StorageResult {
	userAuthKey := fmt.Sprintf("user:%d", userID)
	ur.Redis.Del(userAuthKey)
	return ur.UserRepository.DeleteSession(userID)
}
