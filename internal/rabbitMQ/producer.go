package rabbitMQ

import "github.com/streadway/amqp"

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
