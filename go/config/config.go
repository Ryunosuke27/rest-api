package config

import (
	"gopkg.in/ini.v1"
	"log"
)

type ConfigList struct {
	Port     string
	DBDriver string
	DBName   string
	LogFile  string
}

var Config ConfigList

func init() {
	LoadConfig()
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		Port:     cfg.Section("web").Key("port").MustString("8080"),
		LogFile:  cfg.Section("web").Key("logfile").String(),
		DBDriver: cfg.Section("db").Key("driver").String(),
		DBName:   cfg.Section("db").Key("name").String(),
	}
}
