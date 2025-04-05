package service

import (
	"context"
	"errors"
	"samurai_api/models"
	"samurai_api/repository"
)

type SamuraiService interface {
	CreateSamurai(ctx context.Context, s *models.Samurai) error
	GetAllSamurais(ctx context.Context) ([]models.Samurai, error)
}

type samuraiService struct {
	repo repository.SamuraiRepository
}

func NewSamuraiService(r repository.SamuraiRepository) SamuraiService {
	return &samuraiService{repo: r}
}

func (s *samuraiService) CreateSamurai(ctx context.Context, samurai *models.Samurai) error {
	// Validation: same name samurai should not exist
	existing, err := s.repo.FindByName(ctx, samurai.Name)
	if err != nil {
		return err
	}
	if existing != nil {
		return errors.New("a samurai with this name already exists")
	}

	if samurai.Name == "" || samurai.ClanID == "" {
		return errors.New("samurai name and clan_id are required")
	}

	return s.repo.Create(ctx, samurai)
}

func (s *samuraiService) GetAllSamurais(ctx context.Context) ([]models.Samurai, error) {
	return s.repo.FindAll(ctx)
}
