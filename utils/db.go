package utils

import (
	"stt-service/conf"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func OpenSQLiteDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(conf.Config.DB_NAME), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
