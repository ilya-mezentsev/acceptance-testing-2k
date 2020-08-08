#!/usr/bin/env bash
if [[ ${ENV_VARS_WERE_SET} != '1' ]]; then
  echo 'env variables are not set'
  exit 1
fi

function runTests() {
  echo "Running $1 tests..."
  bash "${PROJECT_ROOT}"/run.sh "$2"

  if [[ $3 = 'print_sep' ]]; then
    echo
    echo '=============================='
    echo
  fi
}

runTests "test runner" test_runner_tests "print_sep"
runTests "API" api_tests "print_sep"
runTests "Backend libs" backend_libs_tests
