package controller

import (
	"dockerwebconsole_v0/config"
	"dockerwebconsole_v0/core"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var defaultConfig = config.GetConfig()

// HomeHandler render the main page with read docker containers
func HomeHandler(w http.ResponseWriter, r *http.Request) {

	p := &core.Page{}

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

	/* Option B

	command := "echo -e 'GET /containers/json HTTP/1.0\r\n' | nc -U /var/run/docker.sock | awk 'END{print}'"

	// Call Run method with command you want to run on remote server.
	body, err := defaultConfig.Ssh.Run(command)
	*/
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

	p = &core.Page{Containers: containers}

	renderTemplate(w, "index", p)
}

// ConsoleHandler run container console for user
func ConsoleHandler(w http.ResponseWriter, r *http.Request, title string) {

	if title == "intro" {
		c := make(chan bool)
		go func() {
			id := r.FormValue("id")
			cmd := defaultConfig.GottyPath + "gotty --once -w -p 9999 docker exec -ti " + id + " bash"
			out, err := defaultConfig.Ssh.Run(cmd)
			if err != nil {
				panic("Can't run remote command: " + err.Error() + out)
			}
			c <- true
		}()
		// wait for the blocking function to finish if it hasn't already
		<-c
	} else if title == "logs" {
		c := make(chan bool)
		go func() {
			id := r.FormValue("id")
			cmd := defaultConfig.GottyPath + "gotty --once -p 8888 docker logs -f " + id
			out, err := defaultConfig.Ssh.Run(cmd)
			if err != nil {
				panic("Can't run remote command: " + err.Error() + out + cmd)
			}
			c <- true
		}()
		// wait for the blocking function to finish if it hasn't already
		<-c
	} else if title == "cmd" {
		c := make(chan bool)
		go func() {
			cmd := defaultConfig.GottyPath + "gotty --once -w -p 7777 docker " + r.FormValue("command")
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

// ContactHandlerNil render a template with contact form and location
func ContactHandlerNil(w http.ResponseWriter, r *http.Request) {

	renderTemplate(w, "contact", &core.Page{})
}

//LoginHandlerNil render a template with login page
func LoginHandlerNil(w http.ResponseWriter, r *http.Request) {

	renderTemplate(w, "login", &core.Page{})
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

//LoginHandler manage a login request
func LoginHandler(w http.ResponseWriter, r *http.Request, title string) {
	if title == "send" {
		/*
			TODO: send data
		*/

		http.Redirect(w, r, "/main", 301)
	} else {
		http.NotFound(w, r)
	}
}

func renderTemplateNil(w http.ResponseWriter, tmpl string) {
	err := defaultConfig.Templates.ExecuteTemplate(w, tmpl+".html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *core.Page) {
	err := defaultConfig.Templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//MakeHandler create a handler for controller
func MakeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := defaultConfig.ValidPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}
