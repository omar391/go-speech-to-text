package utils

import (
	"log"
	"stt-service/conf"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func OpenSQLiteDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(conf.Config.DATA_DIR+conf.Config.DB_NAME), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	return db
}
