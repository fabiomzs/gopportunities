package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	// Init SQLite
	db, err = InitializeSQLite()

	if err != nil {
		return fmt.Errorf("error initializing sqlite;: %v", err)
	}

	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	// Init logger
	logger = NewLogger(p)
	return logger
}
