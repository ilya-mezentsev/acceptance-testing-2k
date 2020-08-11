#!/usr/bin/env bash
rootFolder="$1"
if [[ ${rootFolder} = '' ]]; then
  echo 'root folder was not provided'
  exit 1
fi

cd "${rootFolder}/backend/proto/src/test_case_runner" || exit
protoc --go_out=plugins=grpc:. test_runner.proto
