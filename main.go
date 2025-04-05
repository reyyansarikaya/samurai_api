package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"samurai_api/db"
	"samurai_api/handlers"
	"samurai_api/repository"
	"samurai_api/service"
)

func main() {
	// 🥷 Show banner
	data, err := os.ReadFile("internal/banner/ascii.txt")
	if err != nil {
		fmt.Println("Samurai API - 義は我が道")
	} else {
		fmt.Println(string(data))
	}

	// MongoDB connection
	client, err := db.ConnectMongoDB()
	if err != nil {
		log.Fatal(err)
	}

	// Clan setup
	clanRepo := repository.NewClanRepository(client)
	clanService := service.NewClanService(clanRepo)
	http.HandleFunc("/clans", handlers.ClanHandler(clanService))

	// Samurai setup
	samuraiRepo := repository.NewSamuraiRepository(client)
	samuraiService := service.NewSamuraiService(samuraiRepo)
	http.HandleFunc("/samurais", handlers.SamuraiHandler(samuraiService))

	// Start server
	log.Println("⚔️ Listening on http://localhost:1600")
	http.ListenAndServe(":1600", nil)
}
