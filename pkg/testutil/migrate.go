package testutil

import (
	"goods-square/internal/infrastracture/dto"

	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&dto.Product{},
	)
}

func TearDown(db *gorm.DB) {
	db.DropTableIfExists(
		&dto.Product{},
	)
}

func SetUp(db *gorm.DB) {
	TearDown(db)
	Migrate(db)
}
