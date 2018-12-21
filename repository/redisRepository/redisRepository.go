package redisRepository

import (
	"fmt"

	"github.com/go-redis/redis"

	"github.com/MiteshSharma/project/logger"
	"github.com/MiteshSharma/project/model"
)

type RedisRepository struct {
	Redis  *redis.Client
	Log    logger.Logger
	Config model.CacheConfig
}

func NewRedisRepository(logger logger.Logger, config model.CacheConfig) *RedisRepository {
	redisRepository := &RedisRepository{
		Log:    logger,
		Config: config,
	}
	redisRepository.Redis = redisRepository.getRedis(config)

	return redisRepository
}

func (s *RedisRepository) getRedis(config model.CacheConfig) *redis.Client {
	var client *redis.Client
	if config.Host != "" {
		client = redis.NewClient(&redis.Options{
			Addr:     s.getRedisURL(config),
			Password: config.Password,
			DB:       0,
		})
	}
	return client
}

func (s *RedisRepository) getRedisURL(config model.CacheConfig) string {
	return fmt.Sprintf("%s:%s", config.Host, config.Port)
}
