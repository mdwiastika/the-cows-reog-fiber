package service

import (
	"context"

	"github.com/mdwiastika/the-cows-reog-fiber/domain"
	"github.com/mdwiastika/the-cows-reog-fiber/dto"
)

type userService struct {
	userRepository domain.UserRepository
}

func (us *userService) Index(ctx context.Context) ([]dto.UserData, error) {
	users, err := us.userRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var userData []dto.UserData
	for _, v := range users {
		userData = append(userData, dto.UserData{
			ID:       v.ID,
			Name:     v.Name,
			Email:    v.Email,
			Password: v.Password,
			GoogleID: v.GoogleID,
		})
	}

	return userData, nil
}

func NewUser(userRepository domain.UserRepository) domain.UserService {
	return &userService{
		userRepository: userRepository,
	}
}
