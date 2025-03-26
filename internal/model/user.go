package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name           string `json:"name" gorm:"not null"`
	Phone          string `json:"phone" gorm:"not null"`
	Location       string `json:"location"`
	PrefferedTime  string `json:"preffered_time"`
	MonthlyPremium int    `json:"monthly_premium"`
}

type UserRequest struct {
	Name           string `json:"name" binding:"required"`
	Phone          string `json:"phone" binding:"required"`
	Location       string `json:"location"`
	PrefferedTime  string `json:"preffered_time"`
	MonthlyPremium int    `json:"monthly_premium"`
}

type UserResponse struct {
	ID             uint      `json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Name           string    `json:"name"`
	Phone          string    `json:"phone"`
	Location       string    `json:"location"`
	PrefferedTime  string    `json:"preffered_time"`
	MonthlyPremium int       `json:"monthly_premium"`
}
