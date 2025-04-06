package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"samurai_api/handlers"
	"samurai_api/messaging"
	"samurai_api/models"
	"samurai_api/repository"
	"samurai_api/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
	"github.com/testcontainers/testcontainers-go/modules/rabbitmq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log/slog"
)

// MongoDB container setup
func setupTestMongo(ctx context.Context, t *testing.T) *mongo.Client {
	container, err := mongodb.Run(ctx, "mongo:6")
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		container.Terminate(ctx)
	})

	connStr, err := container.ConnectionString(ctx)
	if err != nil {
		t.Fatal(err)
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connStr))
	if err != nil {
		t.Fatal(err)
	}

	return client
}

// RabbitMQ container setup
func setupTestRabbitMQ(ctx context.Context, t *testing.T) *messaging.AttackEventPublisher {
	rmq, err := rabbitmq.RunContainer(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		rmq.Terminate(ctx)
	})

	host, err := rmq.Host(ctx)
	if err != nil {
		t.Fatal(err)
	}
	port, err := rmq.MappedPort(ctx, "5672")
	if err != nil {
		t.Fatal(err)
	}

	amqpURI := fmt.Sprintf("amqp://guest:guest@%s:%s/", host, port.Port())

	rawPublisher := messaging.NewRabbitMQPublisher(amqpURI)
	if err != nil {
		t.Fatal(err)
	}

	publisher := messaging.NewAttackEventPublisher(rawPublisher, "samurai.attacked")

	t.Cleanup(func() {
		rawPublisher.Close()
	})

	return publisher
}

func TestSamuraiHandler_TableDriven(t *testing.T) {
	ctx := context.Background()
	client := setupTestMongo(ctx, t)
	publisher := setupTestRabbitMQ(ctx, t)

	repo := repository.NewSamuraiRepository(client)
	svc := service.NewSamuraiService(repo, publisher)
	handler := handlers.SamuraiHandler(svc)

	tests := []struct {
		name       string
		samurai    models.Samurai
		wantStatus int
	}{
		{
			name: "valid samurai 1",
			samurai: models.Samurai{
				Name:   "Musashi",
				Rank:   "Ronin",
				ClanID: "Kurogane",
			},
			wantStatus: http.StatusCreated,
		},
		{
			name: "missing name",
			samurai: models.Samurai{
				Rank:   "Ronin",
				ClanID: "Kurogane",
			},
			wantStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			body, _ := json.Marshal(tc.samurai)

			req := httptest.NewRequest(http.MethodPost, "/samurais", bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			handler.ServeHTTP(rec, req)

			slog.Info("Test case", "name", tc.name, "status", rec.Code)
			assert.Equal(t, tc.wantStatus, rec.Code)
		})
	}
}
