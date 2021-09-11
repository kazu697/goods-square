package testutil

import (
	"goods-square/pkg/database"

	"github.com/jinzhu/gorm"
)

type Env struct {
	TestDBHost     string `required:"true"`
	TestDBPort     int    `required:"true"`
	TestDBNet      string `required:"true"`
	TestDBUser     string `required:"true"`
	TestDBPassword string `required:"true"`
}

func TestSQLClient() *gorm.DB {
	// var env Env
	// if err := envconfig.Process("", &env); err != nil {
	// 	panic(err)
	// }
	cfg := &database.Config{
		DBHost:     "localhost",
		DBPort:     3306,
		DBNet:      "tcp",
		DBUser:     "root",
		DBPassword: "MYSQL_ROOT_PASSWORD",
		DBName:     "tests",
	}
	db, err := gorm.Open("mysql", cfg.OutputDBConfig().FormatDSN())

	if err != nil {
		panic(err)
	}
	return db
}
