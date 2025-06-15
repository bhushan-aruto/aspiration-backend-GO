package database

import (
	"context"
	"log"
	"time"

	"github.com/bhushan-aruto/aspiration-matters-backend/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
)

func (db *MongoDBdatabase) AddEventGalleryImage(image *entity.EventGallary) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	collection := db.Collection("event-gallery")
	_, err := collection.InsertOne(ctx, image)
	return err
}

func (db *MongoDBdatabase) GetAllEventGalleryImages() ([]*entity.EventGallary, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	collection := db.Collection("event-gallery")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var images []*entity.EventGallary

	for cursor.Next(ctx) {
		var img entity.EventGallary
		if err := cursor.Decode(&img); err != nil {
			log.Printf("Error decoding image: %v", err)
			continue
		}
		images = append(images, &img)

	}
	return images, nil

}

func (db *MongoDBdatabase) DeleteEventImagebyFileName(fileName string) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	collection := db.Collection("event-gallery")
	_, err := collection.DeleteOne(ctx, bson.M{
		"file_name": fileName,
	})

	return err

}
