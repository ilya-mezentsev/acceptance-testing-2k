#!/usr/bin/env bash

SCRIPTS_FOLDER=$(pwd)/scripts
declare -A scriptsDescriptions=(
  ['help']='show this help'
  ['calc_go_lines']='calculate lines of *.go files'
  ['test_runner_tests']='run tests of test runner'
  ['push_all']='push all files to repository'
)

function run() {
  if [[ -f ./.env ]]; then
    set -o allexport
    source ./.env
    set +o allexport
    scriptName=$1
    shift
    bash "${SCRIPTS_FOLDER}"/"${scriptName}" "$*"
  else
    echo file "$(pwd)"/.env not found
    exit 1
  fi
}

function showHelp {
  echo 'usage bash run.sh <command>'
  echo 'available commands:'
  echo -e '\t-h, -help, help - ' "${scriptsDescriptions['help']}"
  for scriptName in "${SCRIPTS_FOLDER}"/*.sh; do
    # private scripts
    if [[ $(basename "${scriptName}") == _* ]]; then
      continue
    fi

    scriptName=$(basename "${scriptName}" | sed 's/\.sh$//1')
    scriptDescription=${scriptsDescriptions[$scriptName]}
    if [[ ${scriptDescription} = '' ]]; then
      scriptDescription='no description'
    fi
    printf "\t%s - %s\n" "${scriptName}" "${scriptDescription}"
  done
}

if [[ $1 = '-h' || $1 = 'help' || $1 = '-help' || $1 = '' ]]; then
  showHelp
  exit 0
fi

scriptName="$1.sh"
if [[ -f ${SCRIPTS_FOLDER}/${scriptName} ]]; then
  shift
  run "${scriptName}" "$*"
else
  echo file "${SCRIPTS_FOLDER}"/"${scriptName}" not found
  showHelp
  exit 1
fi
