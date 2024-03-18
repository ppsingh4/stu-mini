package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct{
	*gorm.DB
}
// NewDatabase creates a new database connection.
func NewDatabase(dsn string) (*gorm.DB, error) {
	// Open a connection to the database
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
