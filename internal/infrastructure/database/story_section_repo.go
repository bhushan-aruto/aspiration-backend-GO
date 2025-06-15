package database

import (
	"context"
	"time"

	"github.com/bhushan-aruto/aspiration-matters-backend/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *MongoDBdatabase) CheckStorySectionExists() (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	collection := db.Collection("story-section")

	_, err := collection.FindOne(ctx, bson.M{}).Raw()

	if err == mongo.ErrNoDocuments {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (db *MongoDBdatabase) CreateStorySection(storysection *entity.StorySection) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	collection := db.Collection("story-section")

	_, err := collection.InsertOne(ctx, storysection)

	return err
}

func (db *MongoDBdatabase) UpdateStorySection(storySection *entity.StorySection) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	collection := db.Collection("story-section")

	_, err := collection.UpdateOne(ctx, bson.M{},
		bson.M{
			"$set": bson.M{
				"image1_url": storySection.Image1Url,
				"image2_url": storySection.Image2Url,
				"image3_url": storySection.Image3url,
				"image4_url": storySection.Image4url,
			},
		},
	)
	return err

}

func (db *MongoDBdatabase) GetstorySection() (*entity.StorySection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	collection := db.Collection("story-section")

	storySection := new(entity.StorySection)

	err := collection.FindOne(
		ctx,
		bson.M{},
	).Decode(storySection)

	return storySection, err
}
