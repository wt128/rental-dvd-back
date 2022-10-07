package infrastructure

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strings"
)

var Db *gorm.DB

func DbInit() error {
	dsnSummary := []string{
		"host=localhost",
		"user=wara",
		"password=password",
		"dbname=mydb1",
		"port=5432",
		"TimeZone=Asia/Tokyo",
	}
	dsn := strings.Join(dsnSummary, " ")
	var err error
	Db, err = gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})
	return err
}
