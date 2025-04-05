package repository

import (
	"context"
	"samurai_api/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type SamuraiRepository interface {
	Create(ctx context.Context, samurai models.Samurai) error
	GetAll(ctx context.Context) ([]models.Samurai, error)
<<<<<<< HEAD
	FindByName(ctx context.Context, name string) (*models.Samurai, error)
=======
>>>>>>> origin/main
}

type samuraiRepository struct {
	collection *mongo.Collection
}

func NewSamuraiRepository(db *mongo.Client) SamuraiRepository {
	return &samuraiRepository{
		collection: db.Database("samurai").Collection("samurais"),
	}
}

func (r *samuraiRepository) Create(ctx context.Context, samurai models.Samurai) error {
	_, err := r.collection.InsertOne(ctx, samurai)
	return err
}

func (r *samuraiRepository) GetAll(ctx context.Context) ([]models.Samurai, error) {
	cursor, err := r.collection.Find(ctx, map[string]interface{}{})
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
<<<<<<< HEAD

func (r *samuraiRepository) FindByName(ctx context.Context, name string) (*models.Samurai, error) {
	var samurai models.Samurai
	err := r.collection.FindOne(ctx, map[string]interface{}{"name": name}).Decode(&samurai)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &samurai, nil
}
=======
>>>>>>> origin/main
