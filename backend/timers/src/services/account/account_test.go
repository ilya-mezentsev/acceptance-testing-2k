package account

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"test_utils"
	"testing"
	"timers_meta/mock/services"
)

var (
	testHash      = "some-hash"
	filesRootPath = path.Join(os.TempDir(), "at2k")
	accountPath   = path.Join(filesRootPath, testHash)
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestService_DeleteExpiredAccountsSuccess(t *testing.T) {
	_ = os.MkdirAll(accountPath, 0777)
	defer func() {
		_ = os.RemoveAll(filesRootPath)
	}()
	test_utils.AssertTrue(test_utils.MustFileExists(accountPath), t)

	hashes := New(
		services.DefaultAccountRepositoryMock{Hashes: []string{testHash}},
		filesRootPath,
	).DeleteExpiredAccounts()

	test_utils.AssertTrue(strings.Contains(
		strings.Join(hashes, "|"),
		testHash,
	), t)
	test_utils.AssertFalse(test_utils.MustFileExists(accountPath), t)
}

func TestService_DeleteExpiredAccountsNoHashes(t *testing.T) {
	_ = os.MkdirAll(accountPath, 0777)
	defer func() {
		_ = os.RemoveAll(filesRootPath)
	}()
	test_utils.AssertTrue(test_utils.MustFileExists(accountPath), t)

	hashes := New(
		services.DefaultAccountRepositoryMock{Hashes: []string{}},
		filesRootPath,
	).DeleteExpiredAccounts()

	test_utils.AssertEqual(0, len(hashes), t)
	test_utils.AssertTrue(test_utils.MustFileExists(accountPath), t)
}

func TestService_DeleteExpiredAccountsUnableToFetch(t *testing.T) {
	_ = os.MkdirAll(accountPath, 0777)
	defer func() {
		_ = os.RemoveAll(filesRootPath)
	}()
	test_utils.AssertTrue(test_utils.MustFileExists(accountPath), t)

	hashes := New(
		services.BadGetHashesAccountRepositoryMock{},
		filesRootPath,
	).DeleteExpiredAccounts()

	test_utils.AssertEqual(0, len(hashes), t)
	test_utils.AssertTrue(test_utils.MustFileExists(accountPath), t)
}

func TestService_DeleteExpiredAccountsUnableToDeleteInDB(t *testing.T) {
	_ = os.MkdirAll(accountPath, 0777)
	defer func() {
		_ = os.RemoveAll(filesRootPath)
	}()
	test_utils.AssertTrue(test_utils.MustFileExists(accountPath), t)

	hashes := New(
		services.BadDeleteAccountRepositoryMock{},
		filesRootPath,
	).DeleteExpiredAccounts()

	test_utils.AssertEqual(0, len(hashes), t)
	test_utils.AssertTrue(test_utils.MustFileExists(accountPath), t)
}

func TestService_DeleteExpiredAccountsUnableToDeleteFromHost(t *testing.T) {
	_ = os.MkdirAll(accountPath, 0)
	defer func() {
		_ = os.RemoveAll(filesRootPath)
	}()

	hashes := New(
		services.DefaultAccountRepositoryMock{Hashes: []string{testHash}},
		filesRootPath,
	).DeleteExpiredAccounts()

	test_utils.AssertEqual(0, len(hashes), t)
}
