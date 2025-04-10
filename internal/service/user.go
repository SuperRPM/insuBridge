package service

import (
	"insuBridge/domain"
	"insuBridge/internal/model"
)

type userService struct {
	userRepo domain.UserRepository
}

func NewUserService(userRepo domain.UserRepository) domain.UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) CreateUser(req model.UserRequest) (*model.UserResponse, error) {
	user := &model.User{
		Name:              req.Name,
		Phone:             req.Phone,
		HbA1c:             req.HbA1c,
		FastingBloodSugar: req.FastingBloodSugar,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return &model.UserResponse{
		ID:                user.ID,
		CreatedAt:         user.CreatedAt,
		UpdatedAt:         user.UpdatedAt,
		Name:              user.Name,
		Phone:             user.Phone,
		HbA1c:             user.HbA1c,
		FastingBloodSugar: user.FastingBloodSugar,
	}, nil
}
