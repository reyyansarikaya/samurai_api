package service

import (
	"context"
<<<<<<< HEAD
	"errors"
	"log/slog"
=======
>>>>>>> origin/main
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
<<<<<<< HEAD
	if samurai.Name == "" {
		slog.Warn("Samurai name is required")
		return errors.New("samurai name is required")
	}

	if samurai.ClanID == "" {
		slog.Warn("Samurai must belong to a clan")
		return errors.New("samurai must belong to a clan")
	}

	existing, err := s.repo.FindByName(ctx, samurai.Name)
	if err != nil {
		return err
	}
	if existing != nil {
		slog.Warn("Samurai already exists", "name", samurai.Name)
		return errors.New("samurai with this name already exists")
	}

	slog.Info("Creating samurai", "name", samurai.Name)
=======
>>>>>>> origin/main
	return s.repo.Create(ctx, samurai)
}

func (s *samuraiService) GetAllSamurais(ctx context.Context) ([]models.Samurai, error) {
	return s.repo.GetAll(ctx)
}
