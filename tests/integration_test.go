package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"samurai_api/handlers"
	"samurai_api/models"
	"samurai_api/repository"
	"samurai_api/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	mongotc "github.com/testcontainers/testcontainers-go/modules/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log/slog"
)

func setupTestMongo(ctx context.Context, t *testing.T) *mongo.Client {
	mongoC, err := mongotc.Run(ctx, "mongo:6")
	if err != nil {
		t.Fatalf("failed to start container: %v", err)
	}

	t.Cleanup(func() {
		if err := testcontainers.TerminateContainer(mongoC); err != nil {
			t.Logf("failed to terminate container: %v", err)
		}
	})

	endpoint, err := mongoC.ConnectionString(ctx)
	if err != nil {
		t.Fatalf("failed to get connection string: %v", err)
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(endpoint))
	if err != nil {
		t.Fatalf("failed to connect to mongo: %v", err)
	}

	return client
}

func TestIntegration_SamuraiAndClan(t *testing.T) {
	ctx := context.Background()
	client := setupTestMongo(ctx, t)

	clanRepo := repository.NewClanRepository(client)
	samuraiRepo := repository.NewSamuraiRepository(client)

	clanService := service.NewClanService(clanRepo)
	samuraiService := service.NewSamuraiService(samuraiRepo)

	clanHandler := handlers.ClanHandler(clanService)
	samuraiHandler := handlers.SamuraiHandler(samuraiService)

	tests := []struct {
		name       string
		handler    http.Handler
		method     string
		url        string
		body       interface{}
		wantStatus int
	}{
		{
			name:       "create clan",
			handler:    clanHandler,
			method:     http.MethodPost,
			url:        "/clans",
			body:       models.Clan{Name: "Takeda"},
			wantStatus: http.StatusCreated,
		},
		{
			name:       "create samurai valid",
			handler:    samuraiHandler,
			method:     http.MethodPost,
			url:        "/samurais",
			body:       models.Samurai{Name: "Kenshin", Rank: "Ronin", ClanID: "Takeda"},
			wantStatus: http.StatusCreated,
		},
		{
			name:       "create samurai missing name",
			handler:    samuraiHandler,
			method:     http.MethodPost,
			url:        "/samurais",
			body:       models.Samurai{Rank: "Ronin", ClanID: "Takeda"},
			wantStatus: http.StatusInternalServerError,
		},
		{
			name:       "create samurai missing clan",
			handler:    samuraiHandler,
			method:     http.MethodPost,
			url:        "/samurais",
			body:       models.Samurai{Name: "Jubei", Rank: "Ronin"},
			wantStatus: http.StatusInternalServerError,
		},
		{
			name:       "get clans",
			handler:    clanHandler,
			method:     http.MethodGet,
			url:        "/clans",
			body:       nil,
			wantStatus: http.StatusOK,
		},
		{
			name:       "get samurais",
			handler:    samuraiHandler,
			method:     http.MethodGet,
			url:        "/samurais",
			body:       nil,
			wantStatus: http.StatusOK,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var req *http.Request
			if tc.body != nil {
				jsonBody, _ := json.Marshal(tc.body)
				req = httptest.NewRequest(tc.method, tc.url, bytes.NewReader(jsonBody))
				req.Header.Set("Content-Type", "application/json")
			} else {
				req = httptest.NewRequest(tc.method, tc.url, nil)
			}

			rec := httptest.NewRecorder()
			tc.handler.ServeHTTP(rec, req)

			slog.Info("Integration test", "case", tc.name, "status", rec.Code)
			assert.Equal(t, tc.wantStatus, rec.Code)
		})
	}
}
