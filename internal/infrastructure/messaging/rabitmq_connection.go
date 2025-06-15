package messaging

import (
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQConnection struct {
	Conn *amqp.Connection
	Chan *amqp.Channel
}

func NewRabitMQConnection(rabitmqurl string) *RabbitMQConnection {

	conn, err := amqp.Dial(rabitmqurl)
	if err != nil {
		log.Fatal("failed to connect the rabbitmq broker :", err.Error())
	}

	cha, err := conn.Channel()
	if err != nil {
		log.Fatalln("failed to open the rabitmt chanel", err.Error())
	}

	log.Println("connected to rabitMQ")

	return &RabbitMQConnection{
		Conn: conn,
		Chan: cha,
	}

}
