package util

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DbClient *gorm.DB

func DbConfig() {
	var err error
	DbClent, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

}
