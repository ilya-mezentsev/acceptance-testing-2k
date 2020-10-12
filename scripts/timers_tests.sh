#!/usr/bin/env bash
rootFolder="$1"
if [[ ${rootFolder} = '' ]]; then
  echo 'root folder was not provided'
  exit 1
fi
set -o allexport; source ${rootFolder}/.env; set +o allexport

echo "Running Timers tests..."

export GOPATH=${BACKEND_TIMERS_PATH}
export GO_SRC=${BACKEND_TIMERS_PATH}
export REPORT_FOLDER=${TIMERS_REPORT_FOLDER}
cd "${PROJECT_ROOT}" && source scripts/go_tests.sh

echo
echo '=============================='
echo
