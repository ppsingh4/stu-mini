package db
import (
    "gorm.io/gorm"
)
// Initialize initializes the database.
func Initialize(dsn string) (*gorm.DB, error) {
    // Establish a connection to the database
    db, err := NewDatabase(dsn)
    if err != nil {
        return nil, err
    }
    // Run database migrations
    err = Migrate(db)
    if err != nil {
        return nil, err
    }
    return db, nil
}
