package service

import (
	"backend/pkg/model"
	"backend/pkg/repository"
	"context"
)

type InternshipService struct {
	repo *repository.InternshipRepository
}

func NewInternshipService(repo *repository.InternshipRepository) *InternshipService {
	return &InternshipService{repo: repo}
}

func (s *InternshipService) CreateInternship(ctx context.Context, internship model.Internship) (int64, error) {
	return s.repo.CreateInternship(ctx, internship)
}

func (s *InternshipService) GetInternshipByID(ctx context.Context, id int64) (model.Internship, error) {
	return s.repo.GetInternshipByID(ctx, id)
}

func (s *InternshipService) UpdateInternship(ctx context.Context, id int64, internship model.Internship) error {
	return s.repo.UpdateInternship(ctx, id, internship)
}

func (s *InternshipService) DeleteInternship(ctx context.Context, id int64) error {
	return s.repo.DeleteInternship(ctx, id)
}

func (s *InternshipService) GetAllInternships(ctx context.Context) ([]model.Internship, error) {
	return s.repo.GetAllInternships(ctx)
}
