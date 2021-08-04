package database

import "github.com/go-sql-driver/mysql"

type Config struct {
	DBHost     string
	DBPort     int
	DBNet      string
	DBUser     string
	DBPassword string
	DBName     string
}

func (c *Config) OutputDBConfig() *mysql.Config {
	return &mysql.Config{
		User:                 c.DBUser,
		Passwd:               c.DBPassword,
		Net:                  c.DBNet,
		Addr:                 c.DBHost,
		DBName:               c.DBName,
		AllowNativePasswords: true,
		ParseTime:            true,
	}
}
