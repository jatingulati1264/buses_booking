package services

import (
    "bus_booking/repositories"
)

type AuthService struct {
    userRepository *repositories.UserRepository
}

func NewAuthService(userRepository *repositories.UserRepository) *AuthService {
    return &AuthService{userRepository: userRepository}
}

func (s *AuthService) Login(username, password string) (string, error) {
    user, err := s.userRepository.GetByUsername(username)
    if err != nil || user.Password != password {
        return "", err
    }
    // Return a token (JWT or similar) for the user
    return "token", nil
}
