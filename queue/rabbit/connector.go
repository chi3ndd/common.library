package rabbit

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/x-cray/logrus-prefixed-formatter"
)

type (
	Connector struct {
		Addr       string
		Username   string
		Password   string
		connection *amqp.Connection
		channel    *amqp.Channel
		Logger     *logrus.Logger
	}
)

func (con *Connector) Initiation() error {
	// Initiation logger
	con.Logger = &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.DebugLevel,
		Formatter: &prefixed.TextFormatter{
			DisableColors:   false,
			TimestampFormat: time.RFC3339,
			FullTimestamp:   true,
			ForceFormatting: true,
		},
	}
	// Initiation Rabbit Connection
	conn, err := amqp.DialConfig(fmt.Sprintf("amqp://%s", con.Addr), amqp.Config{
		SASL: []amqp.Authentication{
			&amqp.PlainAuth{
				Username: con.Username,
				Password: con.Password,
			},
		},
	})
	con.Logger.Infof("[RabbitMQ] Initializing connection to [%s] - User (%s)", con.Addr, con.Username)
	if err != nil {
		return err
	}
	con.connection = conn
	con.channel, err = conn.Channel()
	// Success
	return err
}

func (con *Connector) Disconnect() error {
	if con.connection != nil && !con.connection.IsClosed() {
		err := con.channel.Close()
		if err != nil {
			con.Logger.Errorf("[RabbitMQ] Close channel error!. Reason: %v", err)
			return err
		}
		err = con.connection.Close()
		if err != nil {
			con.Logger.Errorf("[RabbitMQ] Close connection error!. Reason: %v", err)
			return err
		}
	}
	con.Logger.Info("[RabbitMQ] Close connection ... Success!")
	// Success
	return nil
}
