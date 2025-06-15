package database

import (
	"context"
	"time"

	"github.com/bhushan-aruto/aspiration-matters-backend/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *MongoDBdatabase) CheckAboutSectionExists() (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	collection := db.Collection("AboutSection")

	_, err := collection.FindOne(ctx, bson.M{}).Raw()

	if err == mongo.ErrNoDocuments {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (db *MongoDBdatabase) CreateAboutSection(aboutSection *entity.AboutUsSection) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	collection := db.Collection("AboutSection")
	_, err := collection.InsertOne(ctx, aboutSection)
	return err
}

func (db *MongoDBdatabase) UpdateAboutSection(aboutSection *entity.AboutUsSection) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	collection := db.Collection("AboutSection")

	_, err := collection.UpdateOne(
		ctx,
		bson.M{},
		bson.M{
			"$set": bson.M{
				"image1_url": aboutSection.Image1Url,
				"image2_url": aboutSection.Image2Url,
			},
		},
	)

	return err
}

func (db *MongoDBdatabase) GetAboutSection() (*entity.AboutUsSection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	collection := db.Collection("AboutSection")

	aboutSection := new(entity.AboutUsSection)

	err := collection.FindOne(
		ctx,
		bson.M{},
	).Decode(aboutSection)

	return aboutSection, err
}
