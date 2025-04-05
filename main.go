package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"samurai_api/db"
	"samurai_api/handlers"
	"samurai_api/repository"
	"samurai_api/service"
)

const port = 1600

func printBanner() {
	banner, err := os.ReadFile("internal/banner/ascii.txt")
	if err == nil {
		fmt.Print(string(banner))
	}
}

func main() {
	printBanner()

	ctx := context.Background()

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
		slog.Info("Using default Mongo URI", "uri", mongoURI)
	} else {
		slog.Info("Using Mongo URI from environment", "uri", mongoURI)
	}

	client, err := db.ConnectMongo(ctx, mongoURI)
	if err != nil {
		slog.Error("failed to connect to Mongo", "error", err)
		os.Exit(1)
	}

	// Layered Architecture
	samuraiRepo := repository.NewSamuraiRepository(client)
	samuraiService := service.NewSamuraiService(samuraiRepo)
	samuraiHandler := handlers.SamuraiHandler(samuraiService)

	clanRepo := repository.NewClanRepository(client)
	clanService := service.NewClanService(clanRepo)
	clanHandler := handlers.ClanHandler(clanService)

	// Router setup
	mux := http.NewServeMux()
	mux.Handle("/samurais", samuraiHandler)
	mux.Handle("/clans", clanHandler)

	addr := fmt.Sprintf(":%d", port)
	slog.Info("Samurai API is running on port 1600 ⚔️")
	err = http.ListenAndServe(addr, mux)
	if err != nil {
		slog.Error("server failed", "error", err)
	}
}
