package models

type Clan struct {
	ID     string `json:"id" bson:"_id,omitempty"` // MongoDB document ID
	Name   string `json:"name" bson:"name"`        // Name of the clan
	Region string `json:"region" bson:"region"`    // Region where the clan is based
	Leader string `json:"leader" bson:"leader"`    // Leader of the clan
}
