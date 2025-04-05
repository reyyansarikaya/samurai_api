package repository

import (
	"context"
	"errors"
	"samurai_api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SamuraiRepository interface {
	Create(ctx context.Context, s *models.Samurai) error
	FindAll(ctx context.Context) ([]models.Samurai, error)
	FindByName(ctx context.Context, name string) (*models.Samurai, error)
}

type samuraiRepo struct {
	collection *mongo.Collection
}

func NewSamuraiRepository(client *mongo.Client) SamuraiRepository {
	return &samuraiRepo{
		collection: client.Database("samurai").Collection("samurais"),
	}
}

func (r *samuraiRepo) Create(ctx context.Context, s *models.Samurai) error {
	_, err := r.collection.InsertOne(ctx, s)
	return err
}

func (r *samuraiRepo) FindAll(ctx context.Context) ([]models.Samurai, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var samurais []models.Samurai
	if err := cursor.All(ctx, &samurais); err != nil {
		return nil, err
	}
	return samurais, nil
}

func (r *samuraiRepo) FindByName(ctx context.Context, name string) (*models.Samurai, error) {
	filter := bson.M{"name": name}
	var samurai models.Samurai
	err := r.collection.FindOne(ctx, filter).Decode(&samurai)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &samurai, nil
}
