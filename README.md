# Project for manage and run acceptance tests

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
