package models

type Samurai struct {
	ID     string `json:"id" bson:"_id,omitempty"` // MongoDB document ID
	Name   string `json:"name" bson:"name"`        // Name of the samurai
	Rank   string `json:"rank" bson:"rank"`        // Rank (e.g., Ronin, Hatamoto, Ashigaru)
	ClanID string `json:"clan_id" bson:"clan_id"`  // ID of the clan the samurai belongs to
}
