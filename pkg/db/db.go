package db

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/the-sanctuary/waddles/pkg/cfg"
	"github.com/the-sanctuary/waddles/pkg/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//WadlDB holds the gorm.DB{} database connection
type WadlDB struct {
	*gorm.DB
}

var (
	//Instance is the current database connection
	Instance *WadlDB
)

//BuildWadlDB connects to the database and returns a WadlDB{} holding the database connection
func BuildWadlDB(config *cfg.Config) WadlDB {
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

//Migrate calls gorm.DB.AutoMigrate() on all dbs
func (wdb *WadlDB) Migrate() {
	wdb.AutoMigrate(&UserActivity{})
	wdb.AutoMigrate(&NitroUserChannel{})
	wdb.AutoMigrate(&NicknameUpdate{})

	wdb.AutoMigrate(&Topic{})
	wdb.AutoMigrate(&TopicChannel{})
	wdb.AutoMigrate(&TopicUser{})
	wdb.AutoMigrate(&TopicTag{})
}
