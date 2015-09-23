package core

import (
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

type HostConfig struct {
}

type Container struct {
	Id         string
	Names      []string
	Image      string
	Command    string
	Created    int
	Ports      []Port
	Labels     map[string]string
	Status     string
	HostConfig HostConfig
}

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
