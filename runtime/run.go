package main

import (
	"dockerwebconsole_v0/config"
	"dockerwebconsole_v0/controller"
	"flag"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

func main() {

	defaultConfig := config.GetConfig()

	flag.Parse()
	http.HandleFunc("/", controller.LoginHandlerNil)
	http.HandleFunc("/contact", controller.ContactHandlerNil)
	http.HandleFunc("/main", controller.HomeHandler)

	http.HandleFunc("/login/", controller.MakeHandler(controller.LoginHandler))
	http.HandleFunc("/console/", controller.MakeHandler(controller.ConsoleHandler))
	http.HandleFunc("/contact/", controller.MakeHandler(controller.ContactHandler))
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("../resources/web"))))

	if *defaultConfig.Addr {
		l, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			log.Fatal(err)
		}
		err = ioutil.WriteFile("final-port.txt", []byte(l.Addr().String()), 0644)
		if err != nil {
			log.Fatal(err)
		}
		s := &http.Server{}
		s.Serve(l)
		return
	}

	http.ListenAndServe(":8080", nil)
}
