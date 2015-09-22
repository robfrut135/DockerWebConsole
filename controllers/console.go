package controllers

import (
	"dockerwebconsole/config"
	"dockerwebconsole/core"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var defaultConfig = config.GetConfig()

// HomeHandler render the main page with read docker containers
func (this *MainController) HomeHandler() {
	this.activeContent("index")

	//******** This page requires login
	sess := this.GetSession("acme")
	if sess == nil {
		this.Redirect("/user/login/main", 302)
		return
	}
	m := sess.(map[string]interface{})
	fmt.Println("username is", m["username"])
	fmt.Println("logged in at", m["timestamp"])

	containers := make(map[string]core.Container)

	/*
		To have: -H unix:///var/run/docker.sock -H tcp://0.0.0.0:4243
	*/

	/* Option A */
	url := "http://" + defaultConfig.Host + ":4243/containers/json"

	response, err := http.Get(url)

	if err != nil {
		core.PrintError(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	fmt.Println(string(body))

	// Handle errors
	if err != nil {
		panic("Can't run remote command: " + err.Error())
	} else {

		dec := json.NewDecoder(strings.NewReader(string(body)))

		// read open bracket
		_, err := dec.Token()
		if err != nil {
			log.Fatal(err)
		}

		// while the array contains values
		for dec.More() {
			var m core.Container

			// decode an array value (Message)
			err := dec.Decode(&m)
			if err != nil {
				log.Fatal(err)
			}

			containers[m.Id[0:10]] = m
		}
	}

	this.Data["Containers"] = containers
}

// ConsoleHandler run container console for user
func (this *MainController) ConsoleHandler() {

	action := this.Ctx.Input.Param(":action")

	if action == "intro" {
		c := make(chan bool)
		go func() {
			id := this.GetString("id")
			cmd := defaultConfig.GottyPath + "gotty --once -w -p 9999 docker exec -ti " + id + " bash"
			out, err := defaultConfig.Ssh.Run(cmd)
			if err != nil {
				panic("Can't run remote command: " + err.Error() + out)
			}
			c <- true
		}()
		// wait for the blocking function to finish if it hasn't already
		<-c
	} else if action == "logs" {
		c := make(chan bool)
		go func() {
			id := this.GetString("id")
			cmd := defaultConfig.GottyPath + "gotty --once -p 8888 docker logs -f " + id
			out, err := defaultConfig.Ssh.Run(cmd)
			if err != nil {
				panic("Can't run remote command: " + err.Error() + out + cmd)
			}
			c <- true
		}()
		// wait for the blocking function to finish if it hasn't already
		<-c
	} else if action == "cmd" {
		c := make(chan bool)
		go func() {
			cmd := defaultConfig.GottyPath + "gotty --once -w -p 7777 docker " + this.GetString("command")
			out, err := defaultConfig.Ssh.Run(cmd)
			if err != nil {
				panic("Can't run remote command: " + err.Error() + out + cmd)
			}
			c <- true
		}()
		// wait for the blocking function to finish if it hasn't already
		<-c
	}
}

//ContactHandler receive a request contact
func ContactHandler(w http.ResponseWriter, r *http.Request, title string) {
	if title == "send" {
		/*
			TODO: send data
		*/

		http.Redirect(w, r, "/main", 301)
	} else {
		http.NotFound(w, r)
	}
}
