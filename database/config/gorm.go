package config

import (
	"log"
	"os"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func defaultLogger() logger.Interface {
	f, _ := os.OpenFile("sql.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	return logger.New(
		log.New(f, "\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		},
	)
}

func NewGormConfig(logger logger.Interface) *gorm.Config {
	if logger == nil {
		return NewDefaultGormConfig()
	}
	return &gorm.Config{
		Logger: logger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		PrepareStmt:     true,
		CreateBatchSize: 100,
		QueryFields:     true,
	}
}

func NewDefaultGormConfig() *gorm.Config {
	return &gorm.Config{
		Logger: defaultLogger(),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		PrepareStmt:     true,
		CreateBatchSize: 100,
		QueryFields:     true,
	}
}
