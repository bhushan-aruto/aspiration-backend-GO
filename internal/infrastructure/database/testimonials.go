package database

import (
	"context"
	"errors"
	"time"

	"github.com/bhushan-aruto/aspiration-matters-backend/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (db *MongoDBdatabase) Addtestimonials(testimonials *entity.Testimonial) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	collection := db.Collection("testimonials")
	_, err := collection.InsertOne(ctx, testimonials)
	return err
}

// func (db *MongoDBdatabase) GetVerifiedTestimonials() ([]*entity.Testimonial, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	collection := db.Collection("testimonials")
// 	cursor, err := collection.Find(ctx, bson.M{"verified": true})
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer cursor.Close(ctx)

// 	var testimonials []*entity.Testimonial
// 	for cursor.Next(ctx) {
// 		var testimonial entity.Testimonial
// 		if err := cursor.Decode(&testimonial); err != nil {
// 			continue
// 		}
// 		testimonials = append(testimonials, &testimonial)
// 	}
// 	return testimonials, nil
// }

func (db *MongoDBdatabase) GetVerifiedTestimonials() ([]*entity.Testimonial, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := db.Collection("testimonials")

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := collection.Find(ctx, bson.M{"verified": true}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var testimonials []*entity.Testimonial
	for cursor.Next(ctx) {
		var testimonial entity.Testimonial
		if err := cursor.Decode(&testimonial); err != nil {
			continue
		}
		testimonials = append(testimonials, &testimonial)
	}
	return testimonials, nil
}

func (db *MongoDBdatabase) GetUnverifiedTestimonials() ([]*entity.Testimonial, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := db.Collection("testimonials")
	cursor, err := collection.Find(ctx, bson.M{"verified": false})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var testimonials []*entity.Testimonial
	for cursor.Next(ctx) {
		var testimonial entity.Testimonial
		if err := cursor.Decode(&testimonial); err != nil {
			continue
		}
		testimonials = append(testimonials, &testimonial)
	}
	return testimonials, nil

}

// func (db *MongoDBdatabase) VerifyTestimonial(id string) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	collection := db.Collection("testimonials")
// 	objId, _ := primitive.ObjectIDFromHex(id)

// 	_, err := collection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": bson.M{"verified": true}})
// 	return err
// }

func (db *MongoDBdatabase) VerifyTestimonial(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := db.Collection("testimonials")

	// No ObjectIDFromHex conversion
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"verified": true}}

	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}

func (db *MongoDBdatabase) DeleteTestimonialByFileName(fileName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := db.Collection("testimonials")
	res, err := collection.DeleteOne(ctx, bson.M{"file_name": fileName})

	if res.DeletedCount == 0 {
		return errors.New("no testimonial found with given filename")
	}

	return err
}
