package model

type UserRequest struct {
	Name           string `json:"name" binding:"required"`
	Phone          string `json:"phone" binding:"required"`
	Location       string `json:"location"`
	PrefferedTime  string `json:"preffered_time"`
	MonthlyPremium int    `json:"monthly_premium"`
}
