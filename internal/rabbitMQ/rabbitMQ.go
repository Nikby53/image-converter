package rabbitMQ

import (
	"github.com/Nikby53/image-converter/internal/configs"
	"github.com/streadway/amqp"
)

type Client struct {
	conf *configs.RabbitMQConfig
}

func NewRabbitMQ(conf *configs.RabbitMQConfig) error {
	conn, err := amqp.Dial(conf.RabbitURL)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	return nil
}
