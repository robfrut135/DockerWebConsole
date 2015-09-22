package config

import (
	"dockerwebconsole_v0/myssh"
	"flag"
)

//Config parameter for docker web console
type Config struct {
	Addr      *bool
	Host      string
	Ssh       *myssh.MakeConfig
	GottyPath string
}

var defaultConfig = Config{
	Addr: flag.Bool("addr", false, "find open address and print to final-port.txt"),
	Host: "localhost",
	Ssh: &myssh.MakeConfig{
		User:     "root",
		Password: "XXXXXXX",
		Server:   "localhost",
		Port:     "22",
	},
	GottyPath: "/home/roberto/GoLangWorkspace/bin/",
}

//GetConfig return configuration by default
func GetConfig() Config {
	return defaultConfig
}
