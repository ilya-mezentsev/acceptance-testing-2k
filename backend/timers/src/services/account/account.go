package account

import (
	"logger"
	"os"
	"path"
	"timers_meta/interfaces"
)

type Service struct {
	repository    interfaces.AccountsRepository
	filesRootPath string
}

func New(
	repository interfaces.AccountsRepository,
	filesRootPath string,
) Service {
	return Service{
		repository:    repository,
		filesRootPath: filesRootPath,
	}
}

func (s Service) DeleteExpiredAccounts() (deletedHashes []string) {
	var err error
	deletedHashes, err = s.repository.GetNonVerifiedAccountsCreatedDayAgo()
	if err != nil {
		logger.ErrorF("Unable to fetch non-verified expired accounts: %v", err)
		return nil
	}

	if len(deletedHashes) == 0 {
		return nil
	}

	err = s.repository.DeleteAccounts(deletedHashes)
	if err != nil {
		logger.ErrorF("Unable to delete non-verified expired accounts: %v", err)
		return nil
	}

	err = s.deleteAccountsData(deletedHashes)
	if err != nil {
		logger.ErrorF("Unable to delete accounts data: %v", err)
		return nil
	}

	return deletedHashes
}

func (s Service) deleteAccountsData(accountHashes []string) error {
	for _, accountHash := range accountHashes {
		err := os.RemoveAll(path.Join(s.filesRootPath, accountHash))
		if err != nil {
			return err
		}
	}

	return nil
}
