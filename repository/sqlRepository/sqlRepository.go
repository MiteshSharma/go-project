package sqlRepository

import (
	"os"

	"github.com/MiteshSharma/project/logger"
	"github.com/MiteshSharma/project/model"
	"github.com/jinzhu/gorm"

	// This package is used as mysql driver with gorm
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type SqlRepository struct {
	DB             *gorm.DB
	Log            logger.Logger
	Config         model.DatabaseConfig
	UserRepository UserRepository
}

func NewSqlRepository(logger logger.Logger, config model.DatabaseConfig) *SqlRepository {
	sqlRepository := &SqlRepository{
		Log:    logger,
		Config: config,
	}
	sqlRepository.DB = sqlRepository.getDb(config)

	sqlRepository.UserRepository = NewSqlUserRepository(sqlRepository)
	return sqlRepository
}

func (s *SqlRepository) User() UserRepository {
	return s.UserRepository
}

func (s *SqlRepository) getDb(config model.DatabaseConfig) *gorm.DB {
	var db *gorm.DB
	switch config.Type {
	case "mysql":
		mysqlDb, err := gorm.Open("mysql", config.ConnectionString)
		if err != nil {
			s.Log.Error("Connecting mysql failed due to error ", logger.Error(err))
			os.Exit(1)
		}
		db = mysqlDb
		break
	default:
		break
	}
	return db
}

func (s *SqlRepository) Close() error {
	return s.DB.Close()
}
