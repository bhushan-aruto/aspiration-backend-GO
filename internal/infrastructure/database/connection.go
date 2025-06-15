package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBdatabase struct {
	Client   *mongo.Client
	database *mongo.Database
}

func NewDatabase(databaseUrl string) *MongoDBdatabase {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(databaseUrl))
	if err != nil {
		log.Fatalln("errror while connectinbg the database", err.Error())
	}
	database := os.Getenv("DATABASE_NAME")
	if database == "" {
		log.Fatalln("database name is missing  or empty")
	}
	db := mongoClient.Database(database)

	if err := mongoClient.Ping(ctx, nil); err != nil {
		log.Println("Error occured with mongoDB ", err.Error())
	}

	log.Println("Conencted to the mongo db")

	return &MongoDBdatabase{
		Client:   mongoClient,
		database: db,
	}
}

func (db *MongoDBdatabase) Collection(name string) *mongo.Collection {
	return db.database.Collection(name)

}
