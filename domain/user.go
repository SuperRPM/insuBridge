package domain

import (
	"insuBridge/internal/model"

	"github.com/gin-gonic/gin"
)

// UserHandler 인터페이스
type UserHandler interface {
	CreateUser(c *gin.Context)
}

// UserRepository 인터페이스
type UserRepository interface {
	Create(user *model.User) error
}

// UserService 인터페이스
type UserService interface {
	CreateUser(req model.UserRequest) (*model.UserResponse, error)
}
