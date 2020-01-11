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


type Admin struct {
	ID          uint `json:"id"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	UserName    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	ProfilePic  string `json:"profilepic"`
	PhoneNumber string `json:"phonenum" `
}

type Agent struct {
	ID          uint `json:"id"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	UserName    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	ProfilePic  string `json:"profilepic"`
	PhoneNumber string `json:"phonenum"`
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
	Rating         uint      `json:"rating"`
	HealthCenterID uint      `json:"healthcenterid"`
	PlacedAt       time.Time `json:"placedat" sql:"DEFAULT:'current_timestamp'"`
}
// UserComment joins feedback givers first name with feedback
type UserComment struct {
	FirstName string
	Comment
}

// HealthCenter represents health centers data
type HealthCenter struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phonenumber"`
	City        string `json:"city"`
	ProfilePic  string `json:"profilepic"`
	AgentID     uint   `json:"agentid"`
	User        User   `json:"user"`
}

// Hcrating represents healthcenters with rating
type Hcrating struct {
	HealthCenter
	Rating float64 `json:"rating"`
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

//Session represents login user session
type Session struct {
	ID         uint
	UUID       string `gorm:"type:varchar(255);not null"`
	Expires    int64  `gorm:"type:varchar(255);not null"`
	SigningKey []byte `gorm:"type:varchar(255);not null"`
}