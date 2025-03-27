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
		Name:  req.Name,
		Phone: req.Phone,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return &model.UserResponse{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		Phone:     user.Phone,
	}, nil
}
