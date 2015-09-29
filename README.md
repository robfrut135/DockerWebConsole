## Docker WebConsole

### Administration of containers from a web interface

Uses
- beego framweork
- custom user management inherited from https://github.com/horrido/acme3
- styles by W3layouts
- gotty https://github.com/yudai/gotty

Features
- show containers in host
- open console in your browser
- on close, close session on container
- container logs in real-time
- run docker commands
- user/admin login
- administration users
- search containers by label

ToDo:
- management different Docker host
- users control based on Labels

### Config

Edit [!(config/config.go)](config/config.go)

    var defaultConfig = Config{
    	Addr: flag.Bool("addr", false, "find open address and print to final-port.txt"),
    	Host: "localhost",
    	Ssh: &myssh.MakeConfig{
    		User:     "root",
    		Password: "XXXXXXX",
    		Server:   "localhost",
    		Port:     "22",
    	},
    	MailHost:       "smtp.gmail.com",
    	MailPort:       465,
    	MailFrom:       "robfrut@gmail.com",
    	MailMailerUser: "robfrut@gmail.com",
    	MailMailerPass: "YYYYYYYYYYYYY",
    	GottyPath:      "/Path/to/your/gotty/",
    }

### Build and run

    bee run


### Demo on Fedora21 - current version
TODO

### Demo on Fedora21 - v0

Video demostration  
[![Live demo](https://raw.githubusercontent.com/robfrut135/DockerWebConsole/master/media/v0/default.png)](https://drive.google.com/file/d/0BymCGWR0IjzkbzBwLWpoMGtjYjA/view?usp=sharing)
