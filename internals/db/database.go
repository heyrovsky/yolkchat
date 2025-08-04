package db

import (
	"fmt"
	"log"
	"path/filepath"
	"sync"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "modernc.org/sqlite"
)

var db *gorm.DB
var dbMutex sync.Mutex

const maxRetries = 5
const retryDelay = 3 * time.Second

func GetDbInstance(path string) (*gorm.DB, error) {
	dbMutex.Lock()
	defer dbMutex.Unlock()

	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			return nil, fmt.Errorf("error getting DB connection: %w", err)
		}
		err = sqlDB.Ping()
		if err == nil {
			return db, nil
		}
		if err := sqlDB.Close(); err != nil {
			return nil, fmt.Errorf("error closing DB connection: %w", err)
		}
	}

	var err error
	for retry := 1; retry <= maxRetries; retry++ {
		db, err = connect(path)
		if err == nil {
			return db, nil
		}
		log.Printf("Database connection failed (Attempt %d/%d): %s", retry, maxRetries, err)
		time.Sleep(retryDelay)
	}

	return nil, fmt.Errorf("failed to establish a database connection after %d attempts", maxRetries)
}

func connect(path string) (*gorm.DB, error) {
	dbPath := filepath.Join(path, "app.db")
	dsn := fmt.Sprintf("file:%s?_pragma=foreign_keys(1)", dbPath)
	return gorm.Open(sqlite.Open(dsn), &gorm.Config{})
}
