package services

import (
    "bus_booking/models"
    "bus_booking/repositories"
)

type RatingService struct {
    ratingRepository *repositories.RatingRepository
}

func NewRatingService(ratingRepository *repositories.RatingRepository) *RatingService {
    return &RatingService{ratingRepository: ratingRepository}
}

func (s *RatingService) CreateRating(rating *models.Rating) error {
    return s.ratingRepository.Create(rating)
}

// Implement more methods as needed
