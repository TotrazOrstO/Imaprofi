package rabbitmq

import (
	"fmt"
	"github.com/TotrazOrstO/Imapro/pkg/config"
	"github.com/streadway/amqp"
)

func InitRabbitMQ(cfg config.RabbitMQConfig) (*amqp.Connection, error) {
	conn, err := amqp.Dial(cfg.URL)
	if err != nil {
		return nil, fmt.Errorf("error connecting to RabbitMQ: %v", err)
	}

	return conn, nil
}
