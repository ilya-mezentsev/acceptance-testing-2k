#!/usr/bin/env bash

function prepareFolders() {
  mkdir -p "$1"/test_report
  mkdir -p "$1"/data
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
  ['REPORT_FOLDER']="${rootFolder}/test_report"
  ['GOPATH']="${rootFolder}"
)

prepareFolders "${rootFolder}"/backend
prepareFiles "${rootFolder}"/backend

for envKey in "${!env[@]}"
do
  echo "${envKey}=${env[$envKey]}" >> "${rootFolder}"/.env
done
