#!/usr/bin/env bash
rootFolder="$1"
if [[ ${rootFolder} = '' ]]; then
  echo 'root folder was not provided'
  exit 1
fi

declare -a backendDeps=(
  "github.com/jmoiron/sqlx"
  "github.com/mattn/go-sqlite3"
  "github.com/golang/protobuf/protoc-gen-go"
  "google.golang.org/grpc"
  "github.com/gorilla/mux"
  "github.com/gorilla/websocket"
)
for lib in "${backendDeps[@]}"
do
  GOPATH="$1"/backend/libs go get -v -u "$lib"
done
