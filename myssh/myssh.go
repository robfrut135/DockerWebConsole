// ssh project ssh.go
package myssh

import (
	"fmt"
	"golang.org/x/crypto/ssh"
)

type MakeConfig struct {
	User     string
	Server   string
	Key      string
	Port     string
	Password string
}

func (ssh_conf *MakeConfig) Run(command string) (string, error) {

	client, session, err := connectToHost(ssh_conf.User, ssh_conf.Password, ssh_conf.Server+":"+ssh_conf.Port)
	if err != nil {
		panic(err)
	}
	out, err := session.CombinedOutput(command)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
	client.Close()

	return string(out), err
}

func connectToHost(user, pass, host string) (*ssh.Client, *ssh.Session, error) {

	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{ssh.Password(pass)},
	}

	client, err := ssh.Dial("tcp", host, sshConfig)
	if err != nil {
		return nil, nil, err
	}

	session, err := client.NewSession()
	if err != nil {
		client.Close()
		return nil, nil, err
	}

	return client, session, nil
}
