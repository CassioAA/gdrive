package queue

import (
	"time"

	ampq "github.com/rabbitmq/amqp091-go"
)

type RabbitmqConfig struct {
	URL       string
	TopicName string
	Timeout   time.Duration
}

type RabbitmqConnection struct {
	cfg  RabbitmqConfig
	conn *ampq.Connection
}

func newRabbitmqConnection(cfg RabbitmqConfig) (rc RabbitmqConnection, err error) {
	rc.cfg = cfg
	rc.conn, err = ampq.Dial(cfg.URL)

	return
}