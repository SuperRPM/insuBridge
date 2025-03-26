package handler

import (
	"net/http"

	"insuBridge/domain"
	"insuBridge/internal/model"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService domain.UserService
}

func NewUserHandler(userService domain.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req model.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "잘못된 요청입니다",
			"error":   err.Error(),
		})
		return
	}

	response, err := h.userService.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "사용자 생성 중 오류가 발생했습니다",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
