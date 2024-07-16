package mysql

import (
	"database/sql"

	"github.com/riete/gorm-manager/database/config"

	"github.com/riete/gorm-manager/database/conn"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New(c Config, cc ConnectionConfig, cnc config.ConnConfig) (*gorm.DB, *sql.DB, error) {
	c.Update(cc)
	db, err := gorm.Open(mysql.New(c.gmc), c.gc)
	if err != nil {
		return nil, nil, err
	}
	sqlDB, err := conn.SetConn(db, cnc.MaxConn, cnc.MaxConnLifetime)
	return db, sqlDB, err
}
