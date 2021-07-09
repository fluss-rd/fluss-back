package rabbit

import (
	"context"
	"encoding/json"

	"github.com/streadway/amqp"
)

type RabbitClient interface {
	Publish(ctx context.Context, exchangeName string, routingKey string, message interface{}) error
	Finish()
}

type rabbitClient struct {
	conn *amqp.Connection
	ch   *amqp.Channel // The best practice is to use one channel per tread. Hence, this is not "thread safe". The consumer of this library can use the same client for multiple threads.
}

func InitRabbitClient(connectionString string) (RabbitClient, error) {
	conn, err := amqp.Dial(connectionString)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &rabbitClient{
		conn: conn,
		ch:   ch,
	}, nil

}

func (client *rabbitClient) Publish(ctx context.Context, exchangeName string, routingKey string, message interface{}) error {
	marshalledBody, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return client.ch.Publish(
		exchangeName,
		routingKey,
		false, // mandatory
		false, // inmediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        marshalledBody,
		},
	)
}

func (client *rabbitClient) Consume(ctx context.Context, queueName string) (<-chan amqp.Delivery, error) {
	return client.ch.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
}

func (client *rabbitClient) Finish() {
	client.conn.Close()
}
