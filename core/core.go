package core

import (
	"dockerwebconsole/config"
	"fmt"
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
