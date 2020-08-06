#!/usr/bin/env bash

function prepareTestRunnerFiles() {
  mkdir -p "$1"/backend/tests_runner/test_report
  mkdir -p "$1"/backend/tests_runner/test_data
  mkdir -p "$1"/backend/tests_runner/test_data/some-hash

  rm -f "$1"/backend/tests_runner/test_data/some-hash/test_cases.txt
  echo "BEGIN
  // some comment (will be ignored)
  CREATE USER {\"hash\": \"some-hash\", \"userName\": \"Piter\"}

  user = GET USER {\"hash\": \"some-hash\"}

  ASSERT user.hash EQUALS 'some-hash'
  ASSERT user.userName EQUALS 'Piter'
END" >> "$1"/backend/tests_runner/test_data/some-hash/test_cases.txt

  rm -f "$1"/backend/tests_runner/test_data/test.db
  touch "$1"/backend/tests_runner/test_data/test.db

  rm -f "$1"/backend/tests_runner/test_data/some-hash/db.db
  touch "$1"/backend/tests_runner/test_data/some-hash/db.db
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
  ['PROTO_PATH']="${rootFolder}/backend/proto"
  ['TEST_RUNNER_REPORT_FOLDER']="${rootFolder}/backend/tests_runner/test_report"
  ['TEST_CASES_ROOT_PATH']="${rootFolder}/backend/tests_runner/test_data/"
  ['TEST_CASES_FILENAME']="test_cases.txt"
  ['TEST_RUNNER_DB_FILE']="${rootFolder}/backend/tests_runner/test_data/test.db"
  ['TEST_ACCOUNT_HASH']="some-hash"
  ['TEST_RUNNER_PATH']="${rootFolder}"/backend/tests_runner
  ['BACKEND_LIBS_PATH']="${rootFolder}"/backend/libs
)

prepareTestRunnerFiles "${rootFolder}"
prepareFiles "${rootFolder}"

for envKey in "${!env[@]}"
do
  echo "${envKey}=${env[$envKey]}" >> "${rootFolder}"/.env
done
