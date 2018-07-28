package app

import (
	"github.com/jinzhu/configor"
)

const Version = "1.0.1"

var Config = struct {
	Port         uint64 `default:"6050"`
	HttpsPort    uint64 `default:"6051"`
	HttpsPemPath string `default:""`
	HttpsKeyPath string `default:""`
	LogPath      string `default:"runtime/comet.log"`
	LogLevel     int    `default:"5"`
	LogSql       bool   `default:"false"`
	SqlLogPath   string `default:"runtime/sql.log"`



	Db struct {
		Host     string `default:"localhost"`
		Port     uint64 `default:"3306"`
		Name     string `default:"apibuilder"`
		User     string `default:""`
		Password string `default:""`
	}

}{}

func InitConfig(configPath string) error {
	return configor.Load(&Config, configPath)
}
