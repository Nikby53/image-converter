package rabbitMQ

import "github.com/streadway/amqp"

type Broker interface {
	Publish(name, key string) error
	QueueDeclare(name string) (*Client, error)
}

func (c *Client) Publish(name, key string) error {
	err := c.channel.Publish(name, key, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Hello World"),
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) QueueDeclare(name string) (*Client, error) {
	q, err := c.channel.QueueDeclare(name, false, false, false, false, nil)
	if err != nil {
		panic(err)

	}
	return &Client{queue: q}, nil
}
