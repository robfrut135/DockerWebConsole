package config

import (
	"dockerwebconsole_v0/myssh"
	"flag"
	"html/template"
	"regexp"
)

//Config parameter for docker web console
type Config struct {
	Addr          *bool
	Host          string
	Ssh           *myssh.MakeConfig
	ResourcesPath string
	TemplatesPath string
	TextPath      string
	Templates     *template.Template
	ValidPath     *regexp.Regexp
	GottyPath     string
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
	TextPath:  "../resources/text/",
	Templates: template.Must(template.ParseFiles("../resources/web/index.html", "../resources/web/contact.html", "../resources/web/login.html")),
	ValidPath: regexp.MustCompile("^/(console|contact|login)/([a-zA-Z0-9]+)$"),
	GottyPath: "/home/roberto/GoLangWorkspace/bin/",
}

//GetConfig return configuration by default
func GetConfig() Config {
	return defaultConfig
}
