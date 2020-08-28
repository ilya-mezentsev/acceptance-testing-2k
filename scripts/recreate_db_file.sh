#!/usr/bin/env bash
rootFolder="$1"
if [[ ${rootFolder} = '' ]]; then
  echo 'root folder was not provided'
  exit 1
fi

cd "${rootFolder}" || exit
rm -f data/data.db
touch data/data.db