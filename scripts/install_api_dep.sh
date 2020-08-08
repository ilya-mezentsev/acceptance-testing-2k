#!/usr/bin/env bash
if [[ ${ENV_VARS_WERE_SET} != '1' ]]; then
  echo 'env variables are not set'
  exit 1
fi

GOPATH=${BACKEND_API_PATH} go get -u -v "$1"
