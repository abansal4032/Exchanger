package main

import (
	"Exchanger/server"
	"Exchanger/server/dbclient"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

var (
	cfgFlag = flag.String("config", "", "config file for Zeus")
)

func main() {
	parseFlags()
	connectDB(*loadConfig(*cfgFlag))
	server.Start()
}

func parseFlags() {
	flag.Parse()
	if *cfgFlag == "" {
		flag.PrintDefaults()
		log.Fatalln("missing argument '-config'")
	}
}

func loadConfig(fileName string) *server.Config {
	//fileName := *cfgFlag
	cfgFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln(fmt.Errorf("unable to read config from '%s': %s", fileName, err.Error()))
	}
	conf := &server.Config{}
	if err := json.Unmarshal(cfgFile, conf); err != nil {
		log.Fatalln(fmt.Errorf("unable to parse config from '%s': %s", fileName, err.Error()))
	}
	server.Conf = conf
	return conf
}

func connectDB(cfg server.Config) {
	if err := dbclient.Connect(cfg.DBConf); err != nil {
		log.Fatalln(err)
	}
}
