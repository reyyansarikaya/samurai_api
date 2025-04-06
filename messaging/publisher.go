package messaging

import (
	"github.com/rabbitmq/amqp091-go"
)

type Publisher interface {
	Publish(queue string, body []byte) error
}

type RabbitMQPublisher struct {
	ch *amqp091.Channel
}

func NewRabbitMQPublisher(ch *amqp091.Channel) *RabbitMQPublisher {
	return &RabbitMQPublisher{ch: ch}
}

func (p *RabbitMQPublisher) Publish(queue string, body []byte) error {
	return p.ch.Publish(
		"",    // default exchange
		queue, // routing key
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
