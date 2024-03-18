
package entity
// Marks represents a marks entity for a student in the database.
type Mark struct {
    ID string `gorm:"primary_key" json:"id"`
	Sub_1  int32 `json:"sub_1"`
	Sub_2  int32 `json:"sub_2"`
	Sub_3  int32 `json:"sub_3"`

}
