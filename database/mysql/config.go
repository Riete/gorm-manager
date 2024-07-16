package mysql

import (
	"fmt"
	"time"

	"github.com/riete/gorm-manager/database/config"

	"github.com/go-sql-driver/mysql"

	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ConnectionConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

type Config struct {
	gmc gmysql.Config
	gc  *gorm.Config
}

func (c *Config) Update(cc ConnectionConfig) {
	c.gmc.DSNConfig.Addr = fmt.Sprintf("%s:%s", cc.Host, cc.Port)
	c.gmc.DSNConfig.Net = "tcp"
	c.gmc.DSNConfig.DBName = cc.DBName
	c.gmc.DSNConfig.User = cc.Username
	c.gmc.DSNConfig.Passwd = cc.Password
}

func (c *Config) GormMySQLConfig() gmysql.Config {
	return c.gmc
}

func (c *Config) GormConfig() *gorm.Config {
	return c.gc
}

func NewConfig(mc *mysql.Config, gmc gmysql.Config, gc *gorm.Config) Config {
	gmc.DSNConfig = mc
	return Config{gmc: gmc, gc: gc}
}

func NewDefaultConfig() Config {
	return Config{
		gmc: NewGormMySQLConfig(),
		gc:  config.NewDefaultGormConfig(),
	}
}

func NewGormMySQLConfig() gmysql.Config {
	return gmysql.Config{
		DSNConfig: &mysql.Config{
			Timeout:   10 * time.Second,
			Loc:       time.Local,
			Params:    map[string]string{"charset": "utf8mb4"},
			ParseTime: true,
		},
		DefaultStringSize: 256,
	}
}
