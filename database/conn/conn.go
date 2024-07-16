package conn

import (
	"database/sql"
	"time"

	"github.com/riete/gorm-manager/database/config"

	"gorm.io/gorm"
)

func SetConn(db *gorm.DB, maxConn int, maxConnLifetime time.Duration) (*sql.DB, error) {
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	if maxConn == 0 {
		maxConn = config.DefaultMaxConn
	}
	if maxConnLifetime == 0 {
		maxConnLifetime = config.DefaultMaxConnLifetime
	}
	sqlDB.SetMaxIdleConns(maxConn)
	sqlDB.SetMaxOpenConns(maxConn)
	sqlDB.SetConnMaxLifetime(maxConnLifetime)
	return sqlDB, nil
}
