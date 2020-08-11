#!/usr/bin/env bash
rootFolder="$1"
if [[ ${rootFolder} = '' ]]; then
  echo 'root folder was not provided'
  exit 1
fi

mkdir -p "$1"/backend/tests_runner/test_report
mkdir -p "$1"/backend/api/test_report
mkdir -p "$1"/backend/libs/test_report

mkdir -p "$1"/backend/test_data
mkdir -p "$1"/backend/test_data/some-hash

rm -f "$1"/backend/test_data/some-hash/test_cases.txt
echo "BEGIN
// some comment (will be ignored)
CREATE USER {\"hash\": \"some-hash\", \"userName\": \"Piter\"}

user = GET USER {\"hash\": \"some-hash\"}

ASSERT user.hash EQUALS 'some-hash'
ASSERT user.userName EQUALS 'Piter'
END" >> "$1"/backend/test_data/some-hash/test_cases.txt

rm -f "$1"/backend/test_data/test.db
touch "$1"/backend/test_data/test.db

rm -f "$1"/backend/test_data/some-hash/db.db
touch "$1"/backend/test_data/some-hash/db.db