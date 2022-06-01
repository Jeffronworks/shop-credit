package models

import "time"

type User struct {
	ID           uint `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	EmailAddress string `json:"email_address"`
	PhoneNumber  string `json:"phone_number"`
}
