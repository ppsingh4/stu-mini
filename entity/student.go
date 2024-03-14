package entity
// Student represents a student entity in the database.
type Student struct {
    ID        uint   `gorm:"primary_key"`
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Grade     string `json:"grade"`
    Phno      string `json:"phno"`
    EmailID   string `json:"email_id"`
    Address   string `json:"address"`
}