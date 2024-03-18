package db

import "gorm.io/gorm"

// StudentModel represents the student table in the database.
type Student struct {
	gorm.Model
	ID        int32 `gorm:"primaryKey"`
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
	Grade     string `gorm:"column:grade"`
	Phno      string `gorm:"column:phno"`
	EmailID   string `gorm:"column:email_id"`
	Address   string `gorm:"column:address"`
}

// MarksModel represents the marks table in the database.
type Mark struct {
	gorm.Model
	ID     string   `gorm:"column:id"`
	Sub_1  int32    `gorm:"column:sub_1"`
	Sub_2  int32    `gorm:"column:sub_2"`
	Sub_3  int32    `gorm:"column:sub_3"`
}
