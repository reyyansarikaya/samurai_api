package models

type AttackRequest struct {
	EnemyName string `json:"enemy_name"`
	Location  string `json:"location"`
}

type AttackResult struct {
	SamuraiID string `json:"samurai_id"`
	EnemyName string `json:"enemy_name"`
	Location  string `json:"location"`
	Result    string `json:"result"` // "victory", "defeat"
}
