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

### Build

    #cd [your path]/src/runtime
    #go build -o dockerwebconsole


### Run

before:

  running docker service
  enable ssh to host
  gotty installed in your host

do:

    #cd [your path]/src/runtime
    #dockerwebconsole
