package handlers

import (
	"context"
	b64 "encoding/base64"
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

		// decodedBody, err := decodeBody(message.Body)
		// if err != nil {
		// 	log.Println("decoding body failed: ", err.Error())
		// 	log.Println(string(message.Body))
		// 	continue
		// }

		// TODO: the logic should be in another function
		moduleMessage := models.Message{}
		err = json.Unmarshal(message.Body, &moduleMessage)
		if err != nil {
			log.Println("unmarshalling_message_failed", err)
			log.Println(string(message.Body))
		}

		err = handler.service.SaveMeasurement(ctx, moduleMessage)
		if err != nil {
			log.Println("saving measurement failed: ", err.Error())
		}
	}

	// TODO: we should be listening to a cancelling signal
	return nil
}

func decodeBody(body []byte) ([]byte, error) {
	var output []byte
	encoding := b64.StdEncoding.WithPadding(b64.NoPadding)

	_, err := encoding.Decode(output, body)
	if err != nil {
		return nil, err
	}

	return output, nil
}
