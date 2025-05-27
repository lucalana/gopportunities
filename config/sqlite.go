package config

import (
	"os"

	"github.com/lucalana/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"
	// Check if the database exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("Database file not found, creating...")
		// Create databse file and directory
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}
	// Create database and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Sqlite opening error: %v", err)
		return nil, err
	}
	// Migrate the schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("Sqlite automigration error: %v", err)
		return nil, err
	}
	// Return DB
	return db, nil
}
