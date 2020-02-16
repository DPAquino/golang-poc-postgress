package model

type User struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (User) TableName() string { return "users" }
