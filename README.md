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

*** 1- Login ***
[image]

*** 2- Load main page ***
[image]

*** 3- Main page ***
[image]

*** 4- Docker command ***
[image]

*** 5- Container console web ***
[image]

*** 6- Docker inspect ***
[image]
