package config

import (
	"os"

	"github.com/gaspartv/API-GO-opportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQlite() (*gorm.DB, error){
	logger := GetLogger("sqlite")
    dbPatch := "./db/main.db"

    // Check if the database file exists
    _, err := os.Stat(dbPatch)
    if os.IsNotExist(err) {
        logger.Info("database file not found, creating...")

        // Create the database file and directory
        err = os.MkdirAll("./db", os.ModePerm)
        if err != nil {
            return nil, err
        }
        file, err := os.Create(dbPatch)
        if err != nil {
            return nil, err
        }
        file.Close()
    }

    // Create DB and connect
    db, err := gorm.Open(sqlite.Open(dbPatch), &gorm.Config{})
    if err != nil {
        logger.ErrorF("sqlite opening error: %v", err)
        return nil, err
    }

    // Migration the schema
    err = db.AutoMigrate(&schemas.Opening{})
    if err != nil {
        logger.ErrorF("sqlite auto migration error: %v", err)
        return nil, err
    }

    // Return the DB
    return db, nil
}
