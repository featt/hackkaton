package service

import (
	"backend/pkg/repository"
	"context"
)

type TechnicalService struct {
	repo *repository.TechnicalRepository
}

func NewTechnicalService(repo *repository.TechnicalRepository) *TechnicalService {
	return &TechnicalService{repo: repo}
}

func (s *TechnicalService) CheckDatabaseConnection(ctx context.Context) error {
	return s.repo.CheckDatabaseConnection(ctx)
}

func (s *TechnicalService) GetAllTables(ctx context.Context) ([]string, error) {
	return s.repo.GetAllTables(ctx)
}

func (s *TechnicalService) GetTableData(ctx context.Context, tableName string, limit int) ([]string, error) {
	return s.repo.GetTableData(ctx, tableName, limit)
}

func (s *TechnicalService) GetTableFields(ctx context.Context, tableName string) ([]string, error) {
	return s.repo.GetTableFields(ctx, tableName)
}
