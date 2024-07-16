package sqlite

import (
	"database/sql"

	"github.com/riete/gorm-manager/database/config"

	"github.com/riete/gorm-manager/database/conn"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func New(c *gorm.Config, path string, cnc config.ConnConfig) (*gorm.DB, *sql.DB, error) {
	db, err := gorm.Open(sqlite.Open(path), c)
	if err != nil {
		return nil, nil, err
	}
	sqlDB, err := conn.SetConn(db, cnc.MaxConn, cnc.MaxConnLifetime)
	return db, sqlDB, err
}
