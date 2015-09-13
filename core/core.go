package core

import (
	"dockerwebconsole_v0/config"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type Port struct {
	IP          string
	PrivatePort int
	PublicPort  int
	Type        string
}

type Label struct {
}

type HostConfig struct {
}

type Container struct {
	Id         string
	Names      []string
	Image      string
	Command    string
	Created    int
	Ports      []Port
	Labels     Label
	Status     string
	HostConfig HostConfig
}

type Page struct {
	Title      string
	Body       []byte
	Containers map[string]Container
}

var defaultConfig = config.GetConfig()

func PrintCommand(cmd *exec.Cmd) {
	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func PrintError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
	}
}

func PrintOutput(outs []byte) {
	if len(outs) > 0 {
		fmt.Printf("==> Output: %s\n", string(outs))
	}
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(defaultConfig.TextPath+filename, p.Body, 0600)
}

func LoadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(defaultConfig.TextPath + filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
