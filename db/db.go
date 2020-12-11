package db

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/the-sanctuary/waddles/db/model"
	"github.com/the-sanctuary/waddles/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//WadlDB holds the gorm.DB{} database connection
type WadlDB struct {
	DB *gorm.DB
}

var (
	//Instance is the current database connection
	Instance *WadlDB
)

//BuildWadlDB connects to the database and returns a WadlDB{} holding the database connection
func BuildWadlDB(config *util.Config) WadlDB {
	var dsn string

	if config.Db.URL == "" {
		dsn = fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
			config.Db.User,
			config.Db.Pass,
			config.Db.Host,
			config.Db.Port,
			config.Db.Name,
		)
		log.Info().Msg("Using database config for connection.")
	} else {
		dsn = config.Db.URL
		log.Info().Msg("Using database URL for connection.")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if util.DebugError(err) {
		log.Fatal().Err(err).Msg("Unable to open connection to database. Quitting....")
	}

	Instance = &WadlDB{DB: db}

	return *Instance
}

//Migrate calls gorm.DB.AutoMigrate() on all models
func (wdb *WadlDB) Migrate() {
	wdb.DB.AutoMigrate(&model.UserActivity{})
	wdb.DB.AutoMigrate(&model.NitroUserChannel{})
}
