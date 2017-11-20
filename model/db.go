package model

import (
	_ "github.com/lib/pq"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"

	"github.com/partyzanex/test-task/config"
)

// реализация паттерна - singleton
type DBConnection struct {
	*gorm.DB
	Config *config.DBConfig
}

var DBConn *DBConnection

// конструктор
func NewDBConnection() (*DBConnection, error) {
	// если не подлючено
	if DBConn == nil {
		conf, err := config.GetDBConfig("./config.ini")
		if err != nil {
			return nil, err
		}

		c, err := gorm.Open(conf.Dialect, conf.GetDsn())
		if err != nil {
			return nil, err
		}

		DBConn = &DBConnection{
			DB:     c,
			Config: conf,
		}
	}

	return DBConn, nil
}
