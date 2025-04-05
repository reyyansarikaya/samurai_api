package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
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

func setupTestMongoDB(ctx context.Context, t *testing.T) *mongo.Client {
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

func TestClanHandler_WithLayers(t *testing.T) {
	ctx := context.Background()
	client := setupTestMongoDB(ctx, t)

	// Build layers manually
	clanRepo := repository.NewClanRepository(client)
	clanService := service.NewClanService(clanRepo)
	handler := handlers.ClanHandler(clanService)

	// 1. POST /clans
	clan := models.Clan{
		Name:   "Raikage",
		Region: "Kyushu",
		Leader: "Hoshigaki",
	}
	body, _ := json.Marshal(clan)

	req := httptest.NewRequest(http.MethodPost, "/clans", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)

	// 2. GET /clans
	reqGet := httptest.NewRequest(http.MethodGet, "/clans", nil)
	recGet := httptest.NewRecorder()
	handler.ServeHTTP(recGet, reqGet)

	assert.Equal(t, http.StatusOK, recGet.Code)

	data, _ := io.ReadAll(recGet.Body)
	var result []models.Clan
	json.Unmarshal(data, &result)

	assert.Equal(t, 1, len(result))
	assert.Equal(t, "Raikage", result[0].Name)
	assert.Equal(t, "Hoshigaki", result[0].Leader)
}
