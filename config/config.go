package config

import (
	"github.com/astaxie/beego/config"
	"fmt"
)

const (
	DefaultAdapter = "ini"
	DefaultSection = "database"
)

type DBConfig struct {
	fileName string

	Dialect  string
	Hostname string
	Username string
	Password string
	DbName   string
	SslMode  string
}

func (dc *DBConfig) read() error {
	fullConfigIni, err := config.NewConfig(DefaultAdapter, dc.fileName)
	if err != nil {
		return err
	}

	configIni, err := fullConfigIni.GetSection(DefaultSection)
	if err != nil {
		return err
	}

	// параметры для конекта с БД вынесены в ini-файл
	dc.Dialect = configIni["dialect"]
	dc.Hostname = configIni["hostname"]
	dc.Username = configIni["username"]
	dc.Password = configIni["password"]
	dc.DbName = configIni["dbname"]
	dc.SslMode = configIni["sslmode"]

	return nil
}

// todo: по идее надо сделать и для других СУБД
func (dc DBConfig) GetDsn() string {
	tpl := "host=%s user=%s dbname=%s sslmode=%s password=%s"
	return fmt.Sprintf(tpl, dc.Hostname, dc.Username, dc.DbName, dc.SslMode, dc.Password)
}

func GetDBConfig(fileName string) (*DBConfig, error) {
	conf := &DBConfig{
		fileName: fileName,
	}
	err := conf.read()

	return conf, err
}
