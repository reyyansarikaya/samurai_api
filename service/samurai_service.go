package service

import (
	"context"
	"samurai_api/models"
	"samurai_api/repository"
)

type SamuraiService interface {
	CreateSamurai(ctx context.Context, samurai models.Samurai) error
	GetAllSamurais(ctx context.Context) ([]models.Samurai, error)
}

type samuraiService struct {
	repo repository.SamuraiRepository
}

func NewSamuraiService(repo repository.SamuraiRepository) SamuraiService {
	return &samuraiService{repo: repo}
}

func (s *samuraiService) CreateSamurai(ctx context.Context, samurai models.Samurai) error {
	return s.repo.Create(ctx, samurai)
}

func (s *samuraiService) GetAllSamurais(ctx context.Context) ([]models.Samurai, error) {
	return s.repo.GetAll(ctx)
}
