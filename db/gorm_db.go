package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetGORM(schemaName string) (*gorm.DB, error) {
	logMode := logger.Silent
	if os.Getenv("MODE") == "DEBUG" {
		logMode = logger.Info
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "pgx",
		DSN:        getURI(schemaName),
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logMode),
	})
	if err != nil {
		return nil, fmt.Errorf("gorm.Open error: %s", err)
	}

	return gormDB, nil
}
