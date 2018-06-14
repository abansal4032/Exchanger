package server

import (
	"Exchanger/server/dbclient"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Config struct {
	DBConf 			dbclient.DBConfig `json:"dbConfig"`
	AccessLog      *lumberjack.Logger `json:"accessLog"`
}


var Conf *Config