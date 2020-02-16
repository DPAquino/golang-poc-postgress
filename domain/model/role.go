package model

type Role struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Description string `json:"description"`
	RoleName  string `json:"role_name"`
}

func (Role) TableName() string { return "roles" }