package repository

import (
	"context"
	"samurai_api/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type ClanRepository interface {
	Create(ctx context.Context, clan models.Clan) error
	GetAll(ctx context.Context) ([]models.Clan, error)
}

type clanRepository struct {
	collection *mongo.Collection
}

func NewClanRepository(db *mongo.Client) ClanRepository {
	return &clanRepository{
		collection: db.Database("samurai").Collection("clans"),
	}
}

func (r *clanRepository) Create(ctx context.Context, clan models.Clan) error {
	_, err := r.collection.InsertOne(ctx, clan)
	return err
}

func (r *clanRepository) GetAll(ctx context.Context) ([]models.Clan, error) {
	cursor, err := r.collection.Find(ctx, map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var clans []models.Clan
	if err := cursor.All(ctx, &clans); err != nil {
		return nil, err
	}
	return clans, nil
}
