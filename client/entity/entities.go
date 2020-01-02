package entity

import (
	"time"
)

// User represents user data
type User struct {
	ID          uint   `json:"id"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	UserName    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	ProfilePic  string `json:"profilepic"`
	PhoneNumber string `json:"phonenum"`
	Role        string `json:"role"`
}

// Rating represents users rating
type Rating struct {
	ID             uint
	UserID         uint
	HealthCenterID uint
	PlacedAt       time.Time `sql:"DEFAULT:'current_timestamp'"`
}

// Comment represents users comment
type Comment struct {
	ID             uint      `json:"id"`
	Comment        string    `json:"comment"`
	UserID         uint      `json:"userid"`
	HealthCenterID uint      `json:"healthcenterid"`
	PlacedAt       time.Time `json:"placedat" sql:"DEFAULT:'current_timestamp'"`
}

// HealthCenter represents health centers data
type HealthCenter struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phonenumber"`
	City        string `json:"city"`
	ProfilePic  string `json:"profilepic"`
	AgentID     uint   `json:"agentid"`
	User        User   `json:"user"`
}

// Service represents health centers services
type Service struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	HealthCenterID uint   `json:"healthcenterid"`
	HealthCenter   HealthCenter
	Status         string `json:"status"`
}
