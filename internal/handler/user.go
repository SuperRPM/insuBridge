package handler

import (
	"net/http"

	"insuBridge/internal/model"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
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

	c.JSON(http.StatusOK, gin.H{
		"name":            req.Name,
		"phone":           req.Phone,
		"location":        req.Location,
		"preffered_time":  req.PrefferedTime,
		"monthly_premium": req.MonthlyPremium,
	})

	// c.Status(http.StatusOK)
}
