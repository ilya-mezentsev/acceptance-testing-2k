#!/usr/bin/env bash
rootFolder="$1"
if [[ ${rootFolder} = '' ]]; then
  echo 'root folder was not provided'
  exit 1
fi
set -o allexport; source ${rootFolder}/.env; set +o allexport

for appsDir in "${TEST_RUNNER_PATH}" "${BACKEND_LIBS_PATH}" "${BACKEND_API_PATH}" "${BACKEND_TIMERS_PATH}"
do
  if test -f "${appsDir}"/main.go; then
    (( linesCount=linesCount+$(wc -l  < "${appsDir}"/main.go) ))
  fi

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
