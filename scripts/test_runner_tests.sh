#!/usr/bin/env bash
rootFolder="$1"
if [[ ${rootFolder} = '' ]]; then
  echo 'root folder was not provided'
  exit 1
fi
set -o allexport; source ${rootFolder}/.env; set +o allexport

echo "Running Test runner tests..."

export GOPATH=${TEST_RUNNER_PATH}
export GO_SRC=${TEST_RUNNER_PATH}
export REPORT_FOLDER=${TEST_RUNNER_REPORT_FOLDER}
cd "${PROJECT_ROOT}" && source scripts/go_tests.sh

echo
echo '=============================='
echo
