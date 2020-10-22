package listener

import (
	"encoding/json"
	"events/events_meta/env"
	"events/events_meta/types"
	"github.com/streadway/amqp"
	"logger"
	"os"
	"time"
)

type Listener struct {
	Subscribe types.Channels
	admin     *Admin
	system    *System
}

var l Listener

func init() {
	adminChannel := NewAdminChannel()
	systemChannel := NewSystemChannel()

	l.Subscribe.Admin = adminChannel
	l.Subscribe.System = systemChannel

	l.admin = adminChannel
	l.system = systemChannel
}

func Get() Listener {
	return l
}

func (l Listener) Init(connection string) {
	go initQueue(connection)
}

func initQueue(connection string) {
	conn, err := amqp.Dial(connection)
	if err != nil {
		logger.ErrorF("Failed to connect to RabbitMQ: %v", err)
		time.Sleep(env.QueueAwaitTimeout)
		initQueue(connection)
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

	deliveries, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logger.ErrorF("Failed to register a consumer: %v", err)
		os.Exit(1)
	}

	logger.Info("Events listener is ready")
	for d := range deliveries {
		l.onMessage(d.Body)
	}
}

func (l Listener) onMessage(data []byte) {
	var message types.Message
	err := json.Unmarshal(data, &message)
	if err != nil {
		logger.ErrorF("Unable to unmarshal message data: %v", err)
		return
	}

	switch message.EventName {
	case env.AdminDeleteAccount:
		accountHash, ok := message.Data.(string)
		if !ok {
			logger.ErrorF("Incompatible type of deleted account hash: %v", message.Data)
			break
		}

		l.admin.EmitDeleteAccount(accountHash)

	case env.SystemCleanExpiredAccountHashes, env.SystemCleanExpiredDBConnections:
		seconds, ok := message.Data.(float64)
		if !ok {
			logger.ErrorF("Incompatible type of connection cache duration: %v", message.Data)
			break
		}

		d := time.Duration(seconds) * time.Second
		switch message.EventName {
		case env.SystemCleanExpiredAccountHashes:
			l.system.EmitCleanExpiredAccountHashes(d)
		case env.SystemCleanExpiredDBConnections:
			l.system.EmitCleanExpiredDBConnections(d)
		}
	}
}
