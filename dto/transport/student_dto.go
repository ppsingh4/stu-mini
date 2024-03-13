package transport

type Student struct {
	ID        string   `gorm:"primary_key"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Grade     string    `json:"grade"`
	Phno      string `json:"phno"`
	EmailID   string `json:"emailID"`
	Address   string `json:"address"`
}
