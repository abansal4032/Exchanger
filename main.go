package main

import (
	"fmt"
	"log"
	"flag"
	"io/ioutil"
	"encoding/json"
	"Exchanger/server"
	"Exchanger/server/dbclient"
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

func loadConfig(fileName string) *dbclient.Config {
	//fileName := *cfgFlag
	cfgFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln(fmt.Errorf("unable to read config from '%s': %s", fileName, err.Error()))
	}
	conf := &dbclient.Config{}
	if err := json.Unmarshal(cfgFile, conf); err != nil {
		log.Fatalln(fmt.Errorf("unable to parse config from '%s': %s", fileName, err.Error()))
	}
	return conf
}

func connectDB(cfg dbclient.Config) {
	fmt.Println(cfg)
	if err := dbclient.Connect(cfg); err != nil {
		log.Fatalln(err)
	}
}