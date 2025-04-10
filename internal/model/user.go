package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name              string `gorm:"type:varchar(100);not null"`
	Phone             string `gorm:"type:varchar(20);uniqueIndex;not null"`
	HbA1c             string `gorm:"type:varchar(20);"`
	FastingBloodSugar string `gorm:"type:varchar(20);"`
}

type UserRequest struct {
	Name              string `json:"name" binding:"required"`
	Phone             string `json:"phone" binding:"required"`
	HbA1c             string `json:"hbA1c"`
	FastingBloodSugar string `json:"fastingBloodSugar"`
}

type UserResponse struct {
	ID                uint      `json:"id"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	Name              string    `json:"name"`
	Phone             string    `json:"phone"`
	HbA1c             string    `json:"hbA1c"`
	FastingBloodSugar string    `json:"fastingBloodSugar"`
}
