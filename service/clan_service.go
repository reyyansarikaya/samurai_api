package service

import (
	"context"
	"samurai_api/models"
	"samurai_api/repository"
)

type ClanService interface {
	CreateClan(ctx context.Context, clan models.Clan) error
	GetAllClans(ctx context.Context) ([]models.Clan, error)
}

type clanService struct {
	repo repository.ClanRepository
}

func NewClanService(repo repository.ClanRepository) ClanService {
	return &clanService{repo: repo}
}

func (s *clanService) CreateClan(ctx context.Context, clan models.Clan) error {
	// Buraya business logic ekleyebilirsin (örn. aynı isimli klan varsa ekleme)
	return s.repo.Create(ctx, clan)
}

func (s *clanService) GetAllClans(ctx context.Context) ([]models.Clan, error) {
	return s.repo.GetAll(ctx)
}
