package models

type User struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
