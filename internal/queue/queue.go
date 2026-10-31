package queue

import (
	"fmt"
	"log"
)

type QueueType int

const (
	RabbitMQ QueueType = iota // 0, 1, 2, for validation
)

type Queue struct {
	qc QueueConnection // a message broker
	cfg any // the respective message broker cfg
}

type QueueConnection interface {
	Publish(msg []byte) error
	Consume() error
}

func (q *Queue) Publish(msg []byte) error {
	return q.qc.Publish(msg)
}

func (q *Queue) Consume() error {
	return q.qc.Consume()
}

func New(qt QueueType, cfg any) *Queue {

	q := new(Queue)

	switch qt {
		case RabbitMQ:
			fmt.Println("to be implemented")
		default:
			log.Fatal("type not implemented")
	}

	return q
}