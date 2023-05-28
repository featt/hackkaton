package service

import (
	"backend/pkg/model"
	"backend/pkg/repository"
	"context"
	"strings"
	"time"
)

type CandidateService struct {
	repository *repository.CandidateRepository
}

var DirectionKeywords = map[string][]string{
	"Инжиниринг":  {"инженер", "дизайн", "технический", "проект"},
	"Бизнес":      {"бизнес", "финансы", "продажи", "маркетинг"},
	"Наука":       {"исследование", "лаборатория", "эксперимент", "теория"},
	"Образование": {"учитель", "школа", "студент", "образование"},
	"Медицина":    {"больница", "медицинский", "пациент", "здоровье"},
	"ИТ":          {"код", "программное обеспечение", "разработка", "алгоритм"},
	"Искусство":   {"искусство", "дизайн", "творческий", "рисовать"},
}

func NewCandidateService(repository *repository.CandidateRepository) *CandidateService {
	return &CandidateService{
		repository: repository,
	}
}

func (s *CandidateService) SaveCandidate(ctx context.Context, candidate *model.CandidateModel) error {
	return s.repository.Create(ctx, candidate)
}

func (s *CandidateService) IsCandidateValid(candidate *model.CandidateModel) bool {
	if candidate.Citizenship != "РФ" {
		return false
	}

	birthDate, err := time.Parse("2006-01-02", candidate.DateOfBirth)
	if err != nil {
		return false
	}
	age := time.Now().Year() - birthDate.Year()
	if age < 18 || age > 35 {
		return false
	}

	keywords, ok := DirectionKeywords[candidate.Education]
	if !ok {
		return false
	}

	for _, keyword := range keywords {
		if strings.Contains(strings.ToLower(candidate.WorkExperience), keyword) {
			return true
		}
	}

	return false
}
