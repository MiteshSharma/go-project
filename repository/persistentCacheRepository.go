package repository

import (
	"github.com/MiteshSharma/project/metrics"
	"github.com/MiteshSharma/project/repository/redisRepository"

	"github.com/MiteshSharma/project/logger"
	"github.com/MiteshSharma/project/model"
)

type PersistentCacheRepository struct {
	*PersistentRepository
	RedisRepository *redisRepository.RedisRepository
}

func NewPersistentCacheRepository(log logger.Logger, config *model.Config, metrics metrics.Metrics) *PersistentCacheRepository {
	repository := &PersistentCacheRepository{
		PersistentRepository: NewPersistentRepository(log, config, metrics),
	}

	repository.RedisRepository = redisRepository.NewRedisRepository(log, config.CacheConfig, repository.SqlRepository)
	return repository
}

func (s *PersistentCacheRepository) User() UserRepository {
	return s.RedisRepository.UserRepository
}
