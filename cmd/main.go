package main

import (
	"github.com/bhushan-aruto/aspiration-matters-backend/config"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/infrastructure/database"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/infrastructure/messaging"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/infrastructure/razorpay"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/infrastructure/server"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/infrastructure/storage"
)

func main() {
	config := config.LoadConfig()

	dbConn := database.NewDatabase(config.DatabaseUrl)
	rabitMQconn := messaging.NewRabitMQConnection(config.RabbitMQUrl)
	producer := messaging.NewRabbitMQProducer(rabitMQconn.Chan)
	s3Conn := storage.NewS3Connection(
		config.S3Region,
		config.AwsAccessKeyId,
		config.AwsSecretAccessKey,
	)

	razorpayPayemntGateWay := razorpay.NewRazorPayRepo(config.PaymentGatewayUrl, config.PaymentGateWayKeyId, config.PaymentGateWayKeySecrete)

	server.StartServer(
		config,
		dbConn,
		producer,
		s3Conn,
		razorpayPayemntGateWay,
	)
}
