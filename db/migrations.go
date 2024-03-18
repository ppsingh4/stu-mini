package db

import (
	"fmt"

	"gorm.io/gorm"
)

// Migrate performs database migrations.
func Migrate(db *gorm.DB) error {
	// Auto-migrate database tables based on model definitions
	// if err := db.AutoMigrate(&Student{}).Error; err != nil {
	// 	// Handle the error here
	// 	log.Fatal(err)
	// }
	// if err := db.AutoMigrate(&Mark{}).Error; err != nil {
	// 	// Handle the error here
	// 	log.Fatal(err)
	// }
	err := db.AutoMigrate(&Student{}, &Mark{})
	if err != nil {
		return err
	}
	fmt.Printf("successfully created 2 tables")
	return nil
}
