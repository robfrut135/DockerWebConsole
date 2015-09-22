package config

import (
	"dockerwebconsole_v0/myssh"
	"flag"
)

//Config parameter for docker web console
type Config struct {
	Addr           *bool
	Host           string
	Ssh            *myssh.MakeConfig
	MailHost       string
	MailPort       int
	MailFrom       string
	MailMailerUser string
	MailMailerPass string
	GottyPath      string
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
	MailHost:       "smtp.gmail.com",
	MailPort:       587,
	MailFrom:       "robfrut@gmail.com",
	MailMailerUser: "robfrut@gmail.com",
	MailMailerPass: "135Transceptor135",
	GottyPath:      "/home/roberto/GoLangWorkspace/bin/",
}

//GetConfig return configuration by default
func GetConfig() Config {
	return defaultConfig
}
