package tests_reader

import (
	"io/ioutil"
	"logger"
	"os"
)

func Read(testCasesPath string) (string, error) {
	data, err := ioutil.ReadFile(testCasesPath)

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
				"tests_filename": testCasesPath,
			},
		}, logger.Error)

		return "", UnknownError
	}
}
