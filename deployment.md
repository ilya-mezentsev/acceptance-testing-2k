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
```bash
$ make env.prepare_workspace
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
