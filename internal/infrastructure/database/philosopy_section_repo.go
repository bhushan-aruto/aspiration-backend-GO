package database

import (
	"context"
	"time"

	"github.com/bhushan-aruto/aspiration-matters-backend/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *MongoDBdatabase) CheckPhilosopySectionExists() (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	collection := db.Collection("philospy-section")

	_, err := collection.FindOne(ctx, bson.M{}).Raw()

	if err == mongo.ErrNoDocuments {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (db *MongoDBdatabase) CreatePhilosopySection(philosopySection *entity.PhilosopySection) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	collection := db.Collection("philospy-section")
	_, err := collection.InsertOne(ctx, philosopySection)
	return err
}

func (db *MongoDBdatabase) UpdatePhilosopySection(philosopySection *entity.PhilosopySection) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	collection := db.Collection("philospy-section")

	_, err := collection.UpdateOne(
		ctx,
		bson.M{},
		bson.M{
			"$set": bson.M{
				"image1_url": philosopySection.Image1Url,
				"image2_url": philosopySection.Image2Url,
			},
		},
	)

	return err
}

func (db *MongoDBdatabase) GetPhilospySection() (*entity.PhilosopySection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	collection := db.Collection("philospy-section")

	philosopySection := new(entity.PhilosopySection)

	err := collection.FindOne(
		ctx,
		bson.M{},
	).Decode(philosopySection)

	return philosopySection, err
}
