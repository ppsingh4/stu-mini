package db

import "gorm.io/gorm"

// Migrate performs database migrations.
func Migrate(db *gorm.DB) error {
	// Auto-migrate database tables based on model definitions
	err := db.AutoMigrate(&StudentModel{}, &MarksModel{})
	if err != nil {
		return err
	}
	return nil
}
