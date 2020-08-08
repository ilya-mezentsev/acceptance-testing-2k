#!/usr/bin/env bash
if [[ ${ENV_VARS_WERE_SET} != '1' ]]; then
  echo 'env variables are not set'
  exit 1
fi

export GOPATH=${BACKEND_LIBS_PATH}
export GO_SRC=${BACKEND_LIBS_PATH}
export REPORT_FOLDER=${BACKEND_LIBS_REPORT_FOLDER}
cd "${PROJECT_ROOT}" && source scripts/_go_tests.sh
