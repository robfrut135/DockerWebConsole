## Docker WebConsole

### Administration of containers from ui web

- show containers in host
- open console in your browser
- on close, close session on container
- container logs in real-time
- run docker commands
- fake user login

ToDo:

- management different Docker host
- login
- users control based on Labels

### Config

In config/config.go set your config options

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

### Build

    #cd $GOPATH/src/runtime
    #go build

### Run

before:
  - running docker service
  - enable ssh to host
  - gotty installed in your host

do:

    #cd $GOPATH/src/dockerwebconsole/runtime
    #./runtime

### Demo on Fedora21

Video demostration

[![Live demo](https://raw.githubusercontent.com/robfrut135/DockerWebConsole/master/media/default.png)](https://drive.google.com/file/d/0BymCGWR0IjzkbzBwLWpoMGtjYjA/view?usp=sharing)
