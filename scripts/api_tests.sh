#!/usr/bin/env bash
rootFolder="$1"
if [[ ${rootFolder} = '' ]]; then
  echo 'root folder was not provided'
  exit 1
fi
set -o allexport; source ${rootFolder}/.env; set +o allexport

echo "Running API tests..."

export GOPATH=${BACKEND_API_PATH}
export GO_SRC=${BACKEND_API_PATH}
export REPORT_FOLDER=${API_REPORT_FOLDER}
cd "${PROJECT_ROOT}" && source scripts/go_tests.sh

echo
echo '=============================='
echo
