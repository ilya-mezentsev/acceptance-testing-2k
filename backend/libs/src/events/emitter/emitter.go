package emitter

import (
	"encoding/json"
	"events/events_meta/env"
	"events/events_meta/types"
	"github.com/streadway/amqp"
	"logger"
	"os"
	"time"
)

type Emitter struct {
	Emit     types.Events
	messages chan types.Message
}

var e Emitter

func init() {
	e.messages = make(chan types.Message)

	e.Emit.System = System{e.messages}
	e.Emit.Admin = Admin{e.messages}
}

func Get() Emitter {
	return e
}

func (e Emitter) Init(connection string) {
	go initQueue(connection, e.messages)
}

func initQueue(connection string, messages chan types.Message) {
	conn, err := amqp.Dial(connection)
	if err != nil {
		logger.ErrorF("Failed to connect to RabbitMQ: %v", err)
		time.Sleep(env.QueueAwaitTimeout)
		initQueue(connection, messages)
	}
	defer func() {
		_ = conn.Close()
	}()

	ch, err := conn.Channel()
	if err != nil {
		logger.ErrorF("Failed to open a channel: %v", err)
		os.Exit(1)
	}
	defer func() {
		_ = ch.Close()
	}()

	q, err := ch.QueueDeclare(
		env.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logger.ErrorF("Failed to declare a queue: %v", err)
		os.Exit(1)
	}

	logger.Info("Events emitter is ready")
	for m := range messages {
		sendMessage(ch, q, m)
	}
}

func sendMessage(ch *amqp.Channel, q amqp.Queue, message types.Message) {
	data, _ := json.Marshal(message)
	err := ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        data,
		},
	)

	if err != nil {
		logger.ErrorF("Failed to publish a message: %v", err)
	}
}
