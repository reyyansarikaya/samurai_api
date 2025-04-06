package messaging

import (
	"encoding/json"
	"log/slog"
	"samurai_api/models"
)

type AttackEventPublisher struct {
	channelName string
	publisher   Publisher
}

func NewAttackEventPublisher(p Publisher, channelName string) *AttackEventPublisher {
	return &AttackEventPublisher{
		channelName: channelName,
		publisher:   p,
	}
}

func (a *AttackEventPublisher) PublishAttackEvent(result *models.AttackResult) error {
	body, err := json.Marshal(result)
	if err != nil {
		slog.Error("Failed to marshal attack result", "error", err)
		return err
	}

	err = a.publisher.Publish(a.channelName, body)
	if err != nil {
		slog.Error("Failed to publish attack event", "error", err)
		return err
	}

	return nil
}
