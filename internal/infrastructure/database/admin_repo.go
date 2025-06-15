package database

import (
	"context"
	"log"
	"time"

	"github.com/bhushan-aruto/aspiration-matters-backend/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *MongoDBdatabase) CheckAdminByEmail(email string) *entity.Admin {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	collection := db.Collection("Admin")
	var admin entity.Admin
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&admin)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("No admin found with this email")
			return nil
		}
		log.Println("Error querying the database:", err)
		return nil
	}
	return &admin
}

func (db *MongoDBdatabase) SaveAdmin(admin *entity.Admin) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	collection := db.Collection("Admin")
	_, err := collection.InsertOne(ctx, admin)

	return err
}
