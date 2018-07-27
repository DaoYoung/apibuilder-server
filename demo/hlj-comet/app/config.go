package app

import (
	"hlj-rest/rest"
	"github.com/jinzhu/configor"
	"github.com/liamylian/jsontime"
)

const Version = "1.0.1"

var Json = jsontime.ConfigWithCustomTimeFormat

var Config = struct {
	Port         uint64 `default:"6050"`
	HttpsPort    uint64 `default:"6051"`
	HttpsPemPath string `default:""`
	HttpsKeyPath string `default:""`
	LogPath      string `default:"runtime/comet.log"`
	LogLevel     int    `default:"5"`
	LogSql       bool   `default:"false"`
	SqlLogPath   string `default:"runtime/sql.log"`

	Comet struct {
		MsgBufSize   int `default:"100"`
		MsgTrunkSize int `default:"100"`
		GcInterval   int `default:"120"`
	}

	Peer struct {
		BufSize      int `default:"100"`
		PingTry      int `default:"5"`
		PingInterval int `default:"5"`
		ReadDeadLine int `default:"5"`
	}

	Nats struct {
		Host     string `default:"localhost"`
		Port     uint64 `default:"4222"`
		User     string `default:""`
		Password string `default:""`
	}

	Db struct {
		Host     string `default:"localhost"`
		Port     uint64 `default:"3306"`
		Name     string `default:"comet"`
		User     string `default:""`
		Password string `default:""`
	}

	WedDb struct {
		Host     string `default:"localhost"`
		Port     uint64 `default:"3306"`
		Name     string `default:"comet"`
		User     string `default:""`
		Password string `default:""`
	}
}{}

func InitConfig(configPath string) error {
	rest.SetJson(Json)
	return configor.Load(&Config, configPath)
}
