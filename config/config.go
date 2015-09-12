package config

import (
	"dockerwebconsole_v0/myssh"
	"flag"
	"html/template"
	"regexp"
)

type Config struct {
	Addr           *bool
	Host           string
	Ssh            *myssh.MakeConfig
	Resources_path string
	Templates_path string
	Text_path      string
	Templates      *template.Template
	ValidPath      *regexp.Regexp
}

var defaultConfig = Config{
	Addr: flag.Bool("addr", false, "find open address and print to final-port.txt"),
	Host: "localhost",
	Ssh: &myssh.MakeConfig{
		User:     "root",
		Password: "Robfrut.512",
		Server:   "localhost",
		Port:     "22",
	},
	Resources_path: "../resources/",
	Templates_path: "../resources/templates/",
	Text_path:      "../resources/text/",
	Templates:      template.Must(template.ParseFiles("../resources/web/index.html", "../resources/web/contact.html", "../resources/web/login.html")),
	ValidPath:      regexp.MustCompile("^/(console|contact|login)/([a-zA-Z0-9]+)$"),
}

func GetConfig() Config {
	return defaultConfig
}
