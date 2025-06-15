package server

import (
	"github.com/bhushan-aruto/aspiration-matters-backend/config"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/infrastructure/database"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/infrastructure/messaging"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/infrastructure/server/routes"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/infrastructure/storage"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/repository"
	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

func StartServer(cfg *config.Config, dbConn *database.MongoDBdatabase, producer *messaging.RabbitMQProducer, s3Conn *storage.S3Connection, paymentGatewayRepo repository.PaymentGateWayRepo) {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{echo.HeaderContentType, echo.HeaderAuthorization},
	}))

	routes.InitRoutes(cfg, e, dbConn, producer, s3Conn, paymentGatewayRepo)

	e.Logger.Fatal(e.Start(cfg.ServerAddress))
}
