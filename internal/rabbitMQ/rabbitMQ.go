package rabbitMQ

import (
	"github.com/Nikby53/image-converter/internal/configs"
	"github.com/Nikby53/image-converter/internal/logs"
	"github.com/streadway/amqp"
)

var logger = logs.NewLogger()

type Client struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queue      amqp.Queue
	logger     *logs.StandardLogger
}

func NewRabbitMQ(conf *configs.RabbitMQConfig) (*Client, error) {
	conn, err := amqp.Dial(conf.RabbitURL)
	if err != nil {
		logger.Fatalf("can't connect to AMQP: %s", err)
		return nil, err
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		logger.Fatalf("can't open the channel: %s", err)
	}
	defer ch.Close()
	q, err := ch.QueueDeclare("", false, false, false, false, nil)
	if err != nil {
		logger.Fatalf("can't declare the queue: %s", err)
	}
	return &Client{connection: conn, channel: ch, queue: q}, nil
}
