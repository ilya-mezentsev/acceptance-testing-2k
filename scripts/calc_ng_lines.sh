#!/usr/bin/env bash
rootFolder="$1"
if [[ ${rootFolder} = '' ]]; then
  echo 'root folder was not provided'
  exit 1
fi
set -o allexport; source ${rootFolder}/.env; set +o allexport

cd "${rootFolder}"/frontend/src/app/ || exit;
find . -name '*' -type f -print0 | xargs -0 cat | wc -l
