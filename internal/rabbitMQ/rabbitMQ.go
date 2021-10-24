package rabbitMQ

import (
	"fmt"

	"github.com/Nikby53/image-converter/internal/configs"
	"github.com/Nikby53/image-converter/internal/logs"
	"github.com/streadway/amqp"
)

var logger = logs.NewLogger()

// Client struct contains.
type Client struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queue      amqp.Queue
}

// NewRabbitMQ is set up rabbitMQ connection.
func NewRabbitMQ(conf *configs.RabbitMQConfig) (*Client, error) {
	conn, err := amqp.Dial(conf.RabbitURL)
	if err != nil {
		return nil, fmt.Errorf("can't connect to AMQP: %w", err)
	}
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {
			logger.Errorf("can't close connection %v", err)
		}
	}(conn)
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("can't create an AMQP channel: %w", err)
	}
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {
			logger.Errorf("can't close channel %v", err)
		}
	}(ch)
	q, err := ch.QueueDeclare("", false, false, false, false, nil)
	if err != nil {
		return nil, fmt.Errorf("can't declare queue %w", err)
	}
	return &Client{connection: conn, channel: ch, queue: q}, nil
}
