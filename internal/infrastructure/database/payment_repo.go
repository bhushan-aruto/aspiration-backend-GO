package database

import (
	"context"
	"log"
	"time"

	"github.com/bhushan-aruto/aspiration-matters-backend/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
)

func (db *MongoDBdatabase) GetCourseAmountByIds(ids []string) (int32, error) {

	log.Println(ids)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	collection := db.Collection("courses")

	cursor, err := collection.Find(ctx, bson.M{
		"_id": bson.M{"$in": ids},
	})

	if err != nil {
		return 0, err
	}

	defer cursor.Close(ctx)

	var totalPrice int32
	for cursor.Next(ctx) {
		course := new(entity.Course)

		if err := cursor.Decode(course); err != nil {
			return 0, err
		}

		totalPrice += int32(course.Price)
	}

	return totalPrice, nil
}

func (db *MongoDBdatabase) AppendPurchasedCoursesIds(userID string, courseIDs []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	collection := db.Collection("my-learning")

	filter := bson.M{"user_id": userID}
	update := bson.M{
		"$addToSet": bson.M{
			"course_ids": bson.M{"$each": courseIDs},
		},
	}
	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}

func (db *MongoDBdatabase) SavePurchaseHistory(purchase *entity.PurchaseHistory) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	collection := db.Collection("purchase-History")
	_, err := collection.InsertOne(ctx, purchase)

	return err
}

func (db *MongoDBdatabase) GetSingleCoursePriceById(id string) (int32, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	collection := db.Collection("courses")

	var course entity.Course
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&course)
	if err != nil {
		return 0, err
	}

	return int32(course.Price), nil
}

func (db *MongoDBdatabase) DeleteCartCourseafterPayment(userID, courseID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	Collection := db.Collection("cart")

	_, err := Collection.DeleteOne(ctx, bson.M{"user_id": userID, "course_id": courseID})
	return err
}
