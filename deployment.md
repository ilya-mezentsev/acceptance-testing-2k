# Project for manage and run acceptance tests

## Requirements:
* Go, 1.15
* g++, 9.3
* protoc, 3.11
* Node.js, 14.5
* NPM, 6.14
* Angular CLI, 10.0
* Docker, 19.03
* Docker-compose, 1.25
* GNU Make, 4.2

## Local deployment
* Prepare workspace
```bash
$ make env.prepare_workspace
```
* Create ```.docker/rabbitmq/etc/rabbitmq.conf``` with content:
```
loopback_users.guest = false
listeners.tcp.default = 5672
default_pass = guest
default_user = guest
hipe_compile = false
management.listener.port = 15672
management.listener.ssl = false
```


## Run backend tests
```bash
$ make tests.backend
```

## Run project
### Build project:
```bash
$ make build.project
```
### Run backend containers:
```bash
$ docker-compose up
```
### ... and frontend:
```bash
$ cd frontend && npm run start
```
