package tests_file

import (
	"io"
	"logger"
	"os"
	"path"
	"services/plugins/hash"
)

type Manager struct {
	filesRootPath string
}

func New(filesRootPath string) Manager {
	return Manager{filesRootPath}
}

// Returns full path to created file
func (m Manager) CreateTestCasesFile(
	accountHash string,
	fileFromRequest io.Reader,
) (string, error) {
	testCasesPath := m.getFilePath(accountHash, m.generateFilename())
	f, err := os.OpenFile(testCasesPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		logger.WithFields(logger.Fields{
			MessageTemplate: "Unable to open file: %v",
			Args: []interface{}{
				err,
			},
			Optional: map[string]interface{}{
				"file_path": testCasesPath,
			},
		}, logger.Error)

		return "", err
	}
	defer func() {
		_ = f.Close()
	}()

	_, err = io.Copy(f, fileFromRequest)
	if err != nil {
		logger.WithFields(logger.Fields{
			MessageTemplate: "Unable to copy to test cases file: %v",
			Args: []interface{}{
				err,
			},
			Optional: map[string]interface{}{
				"file_path": testCasesPath,
			},
		}, logger.Error)

		return "", err
	}

	return f.Name(), nil
}

func (m Manager) generateFilename() string {
	return "test_cases_" + hash.Md5WithTimeAsKey("tmp") + ".txt"
}

func (m Manager) getFilePath(accountHash, filename string) string {
	return path.Join(m.filesRootPath, accountHash, filename)
}

func (m Manager) RemoveFile(filePath string) error {
	return os.Remove(filePath)
}
