package db

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/the-sanctuary/waddles/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type WadlDB struct {
	DB *gorm.DB
}

var (
	wadlDB WadlDB
)

func CurrentWadlDB() *WadlDB {
	return &wadlDB
}

func NewWadlDB() WadlDB {
	var dsn string

	if util.Cfg.Db.URL == "" {
		dsn = fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
			util.Cfg.Db.User,
			util.Cfg.Db.Pass,
			util.Cfg.Db.Host,
			util.Cfg.Db.Port,
			util.Cfg.Db.Name,
		)
		log.Info().Msg("Using database config for connection.")
	} else {
		dsn = util.Cfg.Db.URL
		log.Info().Msg("Using database URL for connection.")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if util.DebugError(err) {
		log.Info().Msg("[WADL] Unable to open connection to database.  Quitting....")
		// log.Debug().Msg("[IERR] " + err.Error())
		os.Exit(1)
	}
	wadlDB = WadlDB{DB: db}
	return wadlDB
}

func (wdb *WadlDB) Migrate() {
	wdb.DB.AutoMigrate(&UserActivity{})
}

func (w WadlDB) GetVersion() {

}
