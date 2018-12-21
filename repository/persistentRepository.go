package repository

import (
	"github.com/MiteshSharma/project/metrics"
	"github.com/MiteshSharma/project/repository/sqlRepository"

	"github.com/MiteshSharma/project/logger"
	"github.com/MiteshSharma/project/model"
)

type PersistentRepository struct {
	SqlRepository *sqlRepository.SqlRepository
	Log           logger.Logger
	Config        *model.Config
	Metrics       metrics.Metrics

	UserRepository *sqlRepository.UserRepository
}

func NewPersistentRepository(log logger.Logger, config *model.Config, metrics metrics.Metrics) *PersistentRepository {
	repository := &PersistentRepository{
		Log:     log,
		Config:  config,
		Metrics: metrics,
	}

	repository.SqlRepository = sqlRepository.NewSqlRepository(log, config.DatabaseConfig)

	return repository
}

func (s *PersistentRepository) User() UserRepository {
	return s.SqlRepository.UserRepository
}

func (s *PersistentRepository) Close() error {
	return s.SqlRepository.Close()
}
