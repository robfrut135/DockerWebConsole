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
  - running docker service
  - enable ssh to host
  - gotty installed in your host and in your PATH

do:
    #cd [your path]/src/runtime
    #dockerwebconsole

### Demo on Win32

***1- Login***

***2- Load main page***

***3- Main page***

***4- Docker command***

***5- Container console web***

***6- Docker inspect***
