package gorm

import (
	"goods-square/pkg/database"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewRDBClient(cfg *database.Config) *gorm.DB {
	db, err := gorm.Open("mysql", cfg.OutputDBConfig().FormatDSN())
	if err != nil {
		panic(err)
	}

	db = db.LogMode(true)
	db = db.BlockGlobalUpdate(true)
	return db
}
