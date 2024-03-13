package entity
// Student represents a student entity in the database.
type Student struct {
    ID        string `gorm:"primary_key" json:"id"`
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Grade     string `json:"grade"`
    Phno   string `json:"phone_no"`
    EmailID   string `json:"email_id"`
    Address   string `json:"address"`
}