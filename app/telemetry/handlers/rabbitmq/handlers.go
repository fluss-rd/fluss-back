package handlers

import (
	"context"
	"encoding/json"
	"log"

	"github.com/flussrd/fluss-back/app/shared/rabbit"
	"github.com/flussrd/fluss-back/app/telemetry/models"
	"github.com/flussrd/fluss-back/app/telemetry/service"
)

type Handler interface {
	HandleMessages(ctx context.Context) error
}

type rabbitMQHandler struct {
	client  rabbit.RabbitClient
	service service.Service
}

func NewRabbitHandler(client rabbit.RabbitClient, service service.Service) Handler {
	return rabbitMQHandler{
		client:  client,
		service: service,
	}
}

func (handler rabbitMQHandler) HandleMessages(ctx context.Context) error {
	ch, err := handler.client.Consume(ctx, "modules-messages") // TODO: no magic strings
	if err != nil {
		return err
	}

	log.Println("Listening for messages...")
	for message := range ch {

		log.Println("routing key: ", message.RoutingKey)

		// TODO: the logic should be in another function
		moduleMessage := models.Message{}
		err = json.Unmarshal(message.Body, &moduleMessage)
		if err != nil {
			log.Println("unmarshalling_message_failed", err)
			log.Println(string(message.Body))
		}

		// TODO: validate in another function
		if moduleMessage.PhoneNumber == "" {
			log.Println("ERROR: missing phone number in message")
			continue
		}

		err = handler.service.SaveMeasurement(ctx, moduleMessage)
		if err != nil {
			log.Println("saving measurement failed: ", err.Error())
		}
	}

	// TODO: we should be listening to a cancelling signal
	return nil
}
