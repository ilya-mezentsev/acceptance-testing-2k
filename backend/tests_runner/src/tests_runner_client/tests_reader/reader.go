package tests_reader

import (
	"io/ioutil"
	"os"
	"path"
	"plugins/logger"
)

type Reader struct {
	rootTestsPath string
}

func New(rootTestsPath string) Reader {
	return Reader{rootTestsPath}
}

func (r Reader) Read(accountHash, testCasesFilename string) (string, error) {
	data, err := ioutil.ReadFile(path.Join(r.rootTestsPath, accountHash, testCasesFilename))

	switch {
	case err == nil:
		return string(data), nil
	case os.IsNotExist(err):
		return "", TestsFileNotFound
	default:
		logger.WithFields(logger.Fields{
			MessageTemplate: "Unexpected error: %v",
			Args: []interface{}{
				err,
			},
			Optional: map[string]interface{}{
				"root_tests_path": r.rootTestsPath,
				"account_hash":    accountHash,
				"tests_filename":  testCasesFilename,
			},
		}, logger.Error)

		return "", UnknownError
	}
}
