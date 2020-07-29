#!/usr/bin/env bash

function prepareFolders() {
  mkdir -p "$1"/test_report
  mkdir -p "$1"/test_data
}

function prepareFiles() {
  rm "$1"/.env 2>/dev/null
  touch "$1"/.env
}

rootFolder="$1"
if [[ ${rootFolder} = '' ]]; then
  echo 'root folder was not provided'
  echo 'usage bash prepare_workspace.sh ROOT_FOLDER'
  exit 1
fi

declare -A env=(
  ['ENV_VARS_WERE_SET']="1"
  ['PROJECT_ROOT']="${rootFolder}"
  ['TEST_RUNNER_REPORT_FOLDER']="${rootFolder}/backend/tests_runner/test_report"
  ['TEST_RUNNER_DB_FILE']="${rootFolder}/backend/tests_runner/test_data/test.db"
  ['TEST_RUNNER_PATH']="${rootFolder}"/backend/tests_runner
  ['BACKEND_LIBS_PATH']="${rootFolder}"/backend/libs
)

prepareFolders "${rootFolder}"/backend/tests_runner
prepareFolders "${rootFolder}"/backend/tests_runner
prepareFiles "${rootFolder}"

for envKey in "${!env[@]}"
do
  echo "${envKey}=${env[$envKey]}" >> "${rootFolder}"/.env
done
