package transport

type Student struct {
	ID        int32   `gorm:"primary_key"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_Name"`
	Grade     string `json:"grade"`
	Phno      string `json:"phno"`
	EmailID   string `json:"email_id"`
	Address   string `json:"address"`
}
