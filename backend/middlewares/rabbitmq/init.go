package rabbitmq

import (
	"fmt"
	"go/constant"
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	QueueName string
	Exchange  string
	Key       string
	MQurl     string
}

func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	rabbitmq := &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, MQurl: constant.MQURL}
	var err error
	rabbitmq.conn, err = amqp.Dial(rabbitmq.MQurl)
	rabbitmq.failOnError(err, "Failed to connect to RabbitMQ")
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnError(err, "Failed to open a channel")
	return rabbitmq
}

func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}

func (r *RabbitMQ) failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
