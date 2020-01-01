package entity

import (
	"time"
)

// User represents user data
type User struct {
	ID          uint `json:"id" gorm:"type:serial"`
	FirstName   string `json:"firstname" gorm:"type:varchar(255);not null"`
	LastName    string `json:"lastname" gorm:"type:varchar(255);not null"`
	UserName    string `json:"username" gorm:"type:varchar(255);not null"`
	Email       string `json:"email" gorm:"type:varchar(255);not null"`
	Password    string `json:"password" gorm:"type:varchar(255);not null"`
	ProfilePic  string `json:"profilepic" gorm:"type:varchar(255);"`
	PhoneNumber string `json:"phonenum" gorm:"type:varchar(255);not null"`
}
type Agent struct {
	ID          uint `json:"id" `
	FirstName   string `json:"firstname" gorm:"type:varchar(255);not null"`
	LastName    string `json:"lastname" gorm:"type:varchar(255);not null"`
	UserName    string `json:"username" gorm:"type:varchar(255);not null"`
	Email       string `json:"email" gorm:"type:varchar(255);not null"`
	Password    string `json:"password" gorm:"type:varchar(255);not null"`
	ProfilePic  string `json:"profilepic" gorm:"type:varchar(255);"`
	PhoneNumber string `json:"phonenum" gorm:"type:varchar(255);not null"`
}



type Admin struct {
	ID          uint `json:"id"`
	FirstName   string `json:"firstname" gorm:"type:varchar(255);not null"`
	LastName    string `json:"lastname" gorm:"type:varchar(255);not null"`
	UserName    string `json:"username" gorm:"type:varchar(255);not null"`
	Email       string `json:"email" gorm:"type:varchar(255);not null"`
	Password    string `json:"password" gorm:"type:varchar(255);not null"`
	ProfilePic  string `json:"profilepic" gorm:"type:varchar(255);"`
	PhoneNumber string `json:"phonenum" gorm:"type:varchar(255);not null"`
}
// HealthCenter represents health centers data
type HealthCenter struct {
	ID          uint `json:"id"`
	Name        string `json:"name" gorm:"type:varchar(255); not null"`
	Email       string `json:"email" gorm:"type:varchar(255); not null"`
	PhoneNumber string `json:"phonenumber" gorm:"type:varchar(255); not null"`
	City        string `json:"city" gorm:"type:varchar(255); not null"`
	ProfilePic  string `json:"profilepic" gorm:"type:varchar(255);"`
	Agent       Agent `gorm:"foreignkey:AgentID"`
	AgentID     uint `json:"agentid"`
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
	ID             uint `json:"id"`
	Comment        string `json:"comment" gorm:"varchar(255); not null"`
	UserID         uint	`json:"userid"`
	HealthCenterID uint	`json:"healthcenterid"`
	PlacedAt       time.Time `json:"placedat" sql:"DEFAULT:'current_timestamp'"`
}


// Service represents health centers services
type Service struct {
	ID             uint `json:"id"`
	Name           string `json:"name" gorm:"varchar(255); not null"`
	Description    string `json:"description" gorm:"varchar(255); not null"`
	HealthCenterID uint `json:"healthcenterid"`
	HealthCenter   HealthCenter
	Status         string `json:"status" gorm:"varchar(255); not null;default:'pending'"`
}
