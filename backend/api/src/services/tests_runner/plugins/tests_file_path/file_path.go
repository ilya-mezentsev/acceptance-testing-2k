package tests_file_path

import (
	"os"
	"path"
)

const prefix = "aat-tmp"

func BuildDirPath(accountHash string) string {
	return path.Join(
		os.TempDir(),
		prefix,
		accountHash,
	)
}

func BuildFilePath(accountHash, filename string) string {
	return path.Join(
		BuildDirPath(accountHash),
		filename,
	)
}
