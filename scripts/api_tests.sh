#!/usr/bin/env bash
if [[ ${ENV_VARS_WERE_SET} != '1' ]]; then
  echo 'env variables are not set'
  exit 1
fi

export GOPATH=${BACKEND_API_PATH}
export GO_SRC=${BACKEND_API_PATH}
export REPORT_FOLDER=${API_REPORT_FOLDER}
cd "${PROJECT_ROOT}" && source scripts/_go_tests.sh
