
package entity
// Marks represents a marks entity for a student in the database.
type Marks struct {
    ID string `gorm:"primary_key" json:"student_id"`
	Maths   uint `json:"maths"`
	Social  uint `json:"social"`
	Science uint `json:"science"`
	English uint `json:"english"`
	Hindi   uint `json:"hindi"`
}