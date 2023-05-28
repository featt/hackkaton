package service

import (
	"backend/pkg/repository"
	"context"
	"log"
)

type LoyaltyService struct {
	userRepo *repository.UserRepository
}

func NewLoyaltyService(userRepo *repository.UserRepository) *LoyaltyService {
	return &LoyaltyService{
		userRepo: userRepo,
	}
}

func (s *LoyaltyService) GetUserLoyalty(ctx context.Context, userID int64) (int, error) {
	user, err := s.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return 0, err
	}
	return user.LoyaltyPoints, nil
}

func (s *LoyaltyService) UpdateUserLoyalty(ctx context.Context, userID int64, loyaltyPoints int) error {
	err := s.userRepo.UpdateUserLoyalty(ctx, userID, loyaltyPoints)
	if err != nil {
		log.Printf("Error in repository updating user loyalty: %v", err)
		return err
	}
	return nil
}
