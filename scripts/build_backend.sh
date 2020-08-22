#!/usr/bin/env bash
rootFolder="$1"
if [[ ${rootFolder} = '' ]]; then
  echo 'root folder was not provided'
  exit 1
fi
set -o allexport; source ${rootFolder}/.env; set +o allexport

cd "${PROJECT_ROOT}"/backend || exit 1

cd tests_runner && \
  GOPATH=${TEST_RUNNER_PATH}:${PROTO_PATH}:${BACKEND_LIBS_PATH} go build main.go && \
  cd ..
echo "Tests runner is built"

cd api && \
  GOPATH=${BACKEND_API_PATH}:${PROTO_PATH}:${BACKEND_LIBS_PATH} go build main.go && \
  cd ..
echo "API is built"
