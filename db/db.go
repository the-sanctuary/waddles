package db

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/the-sanctuary/waddles/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type wadldata struct {
	DB *gorm.DB
}

func NewWadlDB() wadldata {
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		util.Cfg.Db.User,
		util.Cfg.Db.Pass,
		util.Cfg.Db.Host,
		util.Cfg.Db.Port,
		util.Cfg.Db.Name,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Info().Msg("[WADL] Unable to open connection to database.  Quitting....")
		log.Debug().Msg("[IERR] " + err.Error())
		os.Exit(1)
	}
	return wadldata{DB: db}
}

func (w wadldata) GetVersion() {

}
