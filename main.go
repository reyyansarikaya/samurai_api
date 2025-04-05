package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"samurai_api/db"
	"samurai_api/handlers"
	"samurai_api/repository"
	"samurai_api/service"
)

func loadBanner() {
	bannerPath := "internal/banner/ascii.txt"
	data, err := os.ReadFile(bannerPath)
	if err != nil {
		slog.Warn("failed to load banner", "error", err)
		return
	}
	fmt.Println(string(data))
}

func main() {
	// Logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	// Banner
	loadBanner()

	// Mongo URI
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		uri = "mongodb://localhost:27017"
		slog.Warn("MONGO_URI not set, defaulting to localhost")
	} else {
		slog.Info("Using Mongo URI from environment", "uri", uri)
	}

	// MongoDB client
	client, err := db.ConnectMongo(uri)
	if err != nil {
		slog.Error("Failed to connect to MongoDB", "error", err)
		os.Exit(1)
	}

	// Samurai Layers
	samuraiRepo := repository.NewSamuraiRepository(client)
	samuraiService := service.NewSamuraiService(samuraiRepo)
	samuraiHandler := handlers.SamuraiHandler(samuraiService)

	// Clan Layers
	clanRepo := repository.NewClanRepository(client)
	clanService := service.NewClanService(clanRepo)
	clanHandler := handlers.ClanHandler(clanService)

	// Routes
	mux := http.NewServeMux()
	mux.HandleFunc("/samurais", samuraiHandler)
	mux.HandleFunc("/clans", clanHandler)

	slog.Info("Samurai API is running on port 1600 ⚔️")
	err = http.ListenAndServe(":1600", mux)
	if err != nil {
		slog.Error("Server failed", "error", err)
		os.Exit(1)
	}
}
