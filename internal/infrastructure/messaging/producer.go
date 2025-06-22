package messaging

import (
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/streadway/amqp"
)

type EmailMesg struct {
	To        string            `json:"to"`
	Subject   string            `json:"subject"`
	EmailType string            `json:"email_type"`
	Data      map[string]string `json:"data"`
}

type RabbitMQProducer struct {
	Channel *amqp.Channel
}

func NewRabbitMQProducer(channel *amqp.Channel) *RabbitMQProducer {
	return &RabbitMQProducer{
		Channel: channel,
	}
}

func (p *RabbitMQProducer) SendMail(message EmailMesg, priority int) error {
	queueName := os.Getenv("QUEUE_NAME")
	if queueName == "" {
		return errors.New("missing or empty env varibale QUEUE_NAME ")
	}

	args := amqp.Table{
		"x-max-priority": int32(10),
	}

	_, err := p.Channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		args,
	)
	if err != nil {
		log.Println("Queue declare failed")
		return err
	}

	body, err := json.Marshal(message)
	if err != nil {
		log.Println("failed to marshal the message")
		return err
	}

	err = p.Channel.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp.Persistent,
			Priority:     uint8(priority),
			Body:         body,
		},
	)
	if err != nil {
		log.Println("failed to publish the messsage")
	}
	log.Println("message published to queue :", queueName)

	return nil

}

func (p *RabbitMQProducer) SendWelcomeMail(to, name string) error {
	msg := EmailMesg{
		To:        to,
		Subject:   "Welcome to Aspiration Matters 🎉",
		EmailType: "welcome",
		Data: map[string]string{
			"name": name,
		},
	}
	return p.SendMail(msg, 5)
}

func (p *RabbitMQProducer) SendOTP(to, otp string) error {
	otpmail := EmailMesg{
		To:        to,
		Subject:   "Your OTP Code !",
		EmailType: "otp",
		Data: map[string]string{
			"otp": otp,
		},
	}

	return p.SendMail(otpmail, 9)
}
