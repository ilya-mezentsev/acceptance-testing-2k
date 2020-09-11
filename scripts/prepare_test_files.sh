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
mkdir -p "$1"/backend/test_data/33d1ff478677b1ac49e4305785a63d70
mkdir -p "$1"/backend/test_data/registration

rm -f "$1"/backend/test_data/33d1ff478677b1ac49e4305785a63d70/test_cases.txt
echo "BEGIN
// some comment (will be ignored)
user = GET USER {\"hash\": \"some-hash\"}

ASSERT user.hash EQUALS 'some-hash'
ASSERT user.userName EQUALS 'Piter'
END" >> "$1"/backend/test_data/33d1ff478677b1ac49e4305785a63d70/test_cases.txt

rm -f "$1"/backend/test_data/test.db
touch "$1"/backend/test_data/test.db

rm -f "$1"/backend/test_data/33d1ff478677b1ac49e4305785a63d70/db.db
touch "$1"/backend/test_data/33d1ff478677b1ac49e4305785a63d70/db.db
