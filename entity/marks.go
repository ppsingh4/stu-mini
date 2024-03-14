
package entity
// Marks represents a marks entity for a student in the database.
type Mark struct {
    ID string `gorm:"primary_key" json:"id"`
	Sub_1  uint `json:"sub_1"`
	Sub_2  uint `json:"sub_2"`
	Sub_3  uint `json:"sub_3"`

}
