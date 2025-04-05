package tests

import (
	"bytes"
	"context"
	"encoding/json"
<<<<<<< HEAD
	"log/slog"
=======
	"io"
>>>>>>> origin/main
	"net/http"
	"net/http/httptest"
	"samurai_api/handlers"
	"samurai_api/models"
	"samurai_api/repository"
	"samurai_api/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setupTestMongoForSamurai(ctx context.Context, t *testing.T) *mongo.Client {
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

<<<<<<< HEAD
func TestSamuraiHandler_TableDriven(t *testing.T) {
	ctx := context.Background()
	client := setupTestMongoForSamurai(ctx, t)

=======
func TestSamuraiHandler_WithLayers(t *testing.T) {
	ctx := context.Background()
	client := setupTestMongoForSamurai(ctx, t)

	// build layers
>>>>>>> origin/main
	repo := repository.NewSamuraiRepository(client)
	svc := service.NewSamuraiService(repo)
	handler := handlers.SamuraiHandler(svc)

<<<<<<< HEAD
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
			name: "valid samurai 2",
			samurai: models.Samurai{
				Name:   "Kenshin",
				Rank:   "Ashigaru",
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
		{
			name: "missing clan_id",
			samurai: models.Samurai{
				Name: "Jin",
				Rank: "Ronin",
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
=======
	// POST /samurais
	samurai := models.Samurai{
		Name:   "Raiden",
		Rank:   "Ronin",
		ClanID: "Kurogane",
	}
	body, _ := json.Marshal(samurai)

	req := httptest.NewRequest(http.MethodPost, "/samurais", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)

	// GET /samurais
	reqGet := httptest.NewRequest(http.MethodGet, "/samurais", nil)
	recGet := httptest.NewRecorder()
	handler.ServeHTTP(recGet, reqGet)

	assert.Equal(t, http.StatusOK, recGet.Code)

	respBody, _ := io.ReadAll(recGet.Body)
	var samurais []models.Samurai
	json.Unmarshal(respBody, &samurais)

	assert.Equal(t, 1, len(samurais))
	assert.Equal(t, "Raiden", samurais[0].Name)
	assert.Equal(t, "Ronin", samurais[0].Rank)
	assert.Equal(t, "Kurogane", samurais[0].ClanID)
>>>>>>> origin/main
}
