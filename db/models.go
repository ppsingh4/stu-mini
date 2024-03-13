package db

import "gorm.io/gorm"

// StudentModel represents the student table in the database.
type StudentModel struct {
	gorm.Model
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
	Grade     string `gorm:"column:grade"`
	Phno   string `gorm:"column:phone_no"`
	EmailID   string `gorm:"column:email_id"`
	Address   string `gorm:"column:address"`
}

// MarksModel represents the marks table in the database.
type MarksModel struct {
	gorm.Model
	ID string `gorm:"column:student_id"`
	Maths  int    `gorm:"column:maths"`
	Social  int    `gorm:"column:social"`
	Science  int    `gorm:"column:science"`
	English int    `gorm:"column:english"`
	Hindi int     `gorm:"column:hindi"`
}
