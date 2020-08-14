#!/usr/bin/env bash
rootFolder="$1"
if [[ ${rootFolder} = '' ]]; then
  echo 'root folder was not provided'
  exit 1
fi

declare -A env=(
  ['PROJECT_ROOT']="${rootFolder}"
  ['PROTO_PATH']="${rootFolder}/backend/proto"
  ['TEST_RUNNER_REPORT_FOLDER']="${rootFolder}/backend/tests_runner/test_report"
  ['API_REPORT_FOLDER']="${rootFolder}/backend/api/test_report"
  ['BACKEND_LIBS_REPORT_FOLDER']="${rootFolder}/backend/libs/test_report"
  ['TEST_CASES_ROOT_PATH']="${rootFolder}/backend/test_data/"
  ['REGISTRATION_ROOT_PATH']="${rootFolder}/backend/test_data/registration"
  ['TEST_CASES_FILENAME']="test_cases.txt"
  ['TEST_DB_FILE']="${rootFolder}/backend/test_data/test.db"
  ['TEST_ACCOUNT_HASH']="some-hash"
  ['TEST_RUNNER_PATH']="${rootFolder}"/backend/tests_runner
  ['BACKEND_LIBS_PATH']="${rootFolder}"/backend/libs
  ['BACKEND_API_PATH']="${rootFolder}"/backend/api
)

rm -f "${rootFolder}"/.env
for envKey in "${!env[@]}"
do
  echo "${envKey}=${env[$envKey]}" >> "${rootFolder}"/.env
done