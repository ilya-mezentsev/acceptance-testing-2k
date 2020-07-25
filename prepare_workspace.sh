#!/usr/bin/env bash

function prepareFolders() {
  mkdir -p "$1"/test_report
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
  ['REPORT_FOLDER']="${rootFolder}/backend/test_report"
  ['DB_FILE']="${rootFolder}/backend/test_data/test.db"
  ['GOPATH']="${rootFolder}"/backend
)

prepareFolders "${rootFolder}"/backend
prepareFiles "${rootFolder}"

for envKey in "${!env[@]}"
do
  echo "${envKey}=${env[$envKey]}" >> "${rootFolder}"/.env
done
