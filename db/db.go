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
	} else {
		dsn = util.Cfg.Db.URL
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if util.DebugError(err) {
		log.Info().Msg("[WADL] Unable to open connection to database.  Quitting....")
		// log.Debug().Msg("[IERR] " + err.Error())
		os.Exit(1)
	}
	return WadlDB{DB: db}
}

func (wdb *WadlDB) Migrate() {
}

func (w WadlDB) GetVersion() {

}
