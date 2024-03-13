
package entity
// Marks represents a marks entity for a student in the database.
type Marks struct {
    Stu_ID uint `gorm:"primary_key" json:"student_id"`
	Sub_1  uint `json:"sub1"`
	Sub_2  uint `json:"sub2"`
	Sub_3  uint `json:"sub3"`

}