package main

import (
	"events/emitter"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"logger"
	accountRepository "repositories/account"
	"services/account"
	"services/timers"
	"time"
	"utils"
)

var (
	radio           emitter.Emitter
	accountsService account.Service
)

func init() {
	accountsService = account.New(
		accountRepository.New(
			sqlx.MustOpen("sqlite3", utils.MustGetEnv("PROJECT_DB_FILE_PATH")),
		),
		utils.MustGetEnv("FILES_ROOT_PATH"),
	)

	radio = emitter.Get()

	radio.Init(utils.MustGetEnv("RABBIT_CONNECTION"))
}

func main() {
	logger.Info("Start timers ...")

	connectionCacheCleanTimer := time.NewTicker(timers.ConnectionCacheCleanTimout())
	checkExpiredAccountTimer := time.NewTicker(timers.CheckExpiredAccountTimeout())

	for {
		select {
		case <-connectionCacheCleanTimer.C:
			radio.Emit.System.CleanExpiredDBConnections(timers.ConnectionCacheLifetime())

		case <-checkExpiredAccountTimer.C:
			expiredAccountHashes := accountsService.DeleteExpiredAccounts()
			if len(expiredAccountHashes) > 0 {
				for _, hash := range expiredAccountHashes {
					radio.Emit.Admin.DeleteAccount(hash)
				}
			}
		}
	}
}
