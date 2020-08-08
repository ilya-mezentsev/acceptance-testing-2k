#!/usr/bin/env bash
if [[ ${ENV_VARS_WERE_SET} != '1' ]]; then
  echo 'env variables are not set'
  exit 1
fi

for appsDir in "${TEST_RUNNER_PATH}" "${BACKEND_LIBS_PATH}" "${BACKEND_API_PATH}"
do
  cd "${appsDir}"/src || exit
  for dir in $(ls)
    do
      if [[
        ${dir} != github.com &&
        ${dir} != golang.org &&
        ${dir} != google.golang.org &&
        ${dir} != bin &&
        ${dir} != pkg
      ]]; then
        cd "${dir}" || exit
        (( linesCount=linesCount+$(find . -name '*.go' -type f -print0 | xargs -0 cat | wc -l) ))
        cd ../
      fi
    done
done

echo ${linesCount}
