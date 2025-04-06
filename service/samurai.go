package service

import (
	"context"
	"errors"
	"log/slog"
	"math/rand"
	"samurai_api/messaging"
	"samurai_api/models"
	"samurai_api/repository"
	"time"
)

type SamuraiService interface {
	CreateSamurai(ctx context.Context, s *models.Samurai) error
	GetAllSamurais(ctx context.Context) ([]models.Samurai, error)
	Attack(samuraiID string, req models.AttackRequest) (*models.AttackResult, error)
	PublishAttackEvent(result *models.AttackResult)
}

type samuraiService struct {
	repo      repository.SamuraiRepository
	publisher *messaging.AttackEventPublisher
}

func NewSamuraiService(repo repository.SamuraiRepository, publisher *messaging.AttackEventPublisher) SamuraiService {
	return &samuraiService{
		repo:      repo,
		publisher: publisher,
	}
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

func (s *samuraiService) Attack(samuraiID string, req models.AttackRequest) (*models.AttackResult, error) {
	// Create a new random source with current time
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// %50 kazanma şansı
	outcome := "victory"
	if r.Intn(2) == 0 {
		outcome = "defeat"
	}

	result := &models.AttackResult{
		SamuraiID: samuraiID,
		EnemyName: req.EnemyName,
		Location:  req.Location,
		Result:    outcome,
	}

	return result, nil
}

func (s *samuraiService) PublishAttackEvent(result *models.AttackResult) {
	if err := s.publisher.PublishAttackEvent(result); err != nil {
		slog.Error("Failed to publish attack event",
			"enemy", result.EnemyName,
			"result", result.Result,
			"error", err)
		return
	}

	slog.Info("Published attack event",
		"enemy", result.EnemyName,
		"result", result.Result)
}
