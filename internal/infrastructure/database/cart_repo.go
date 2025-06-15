package database

import (
	"context"
	"fmt"
	"time"

	"github.com/bhushan-aruto/aspiration-matters-backend/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
)

func (db *MongoDBdatabase) AddToCart(item *entity.CartItem) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	collection := db.Collection("cart")
	_, err := collection.InsertOne(ctx, item)
	return err
}

func (db *MongoDBdatabase) GetCartCourse(userID string) ([]*entity.Course, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	cartCollection := db.Collection("cart")
	courseCollection := db.Collection("courses")

	cursor, err := cartCollection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var courseIDs []string
	for cursor.Next(ctx) {
		var item entity.CartItem
		if err := cursor.Decode(&item); err != nil {
			fmt.Println("Error decoding cart item:", err)
			continue
		}

		courseIDs = append(courseIDs, item.CourseID)
	}

	if len(courseIDs) == 0 {
		return []*entity.Course{}, nil
	}

	var courses []*entity.Course
	courseCursor, err := courseCollection.Find(ctx, bson.M{"_id": bson.M{"$in": courseIDs}})
	if err != nil {
		return nil, err
	}
	defer courseCursor.Close(ctx)

	for courseCursor.Next(ctx) {
		var c entity.Course
		if err := courseCursor.Decode(&c); err != nil {
			fmt.Println("Error decoding course:", err)
			continue
		}
		courses = append(courses, &c)
	}

	return courses, nil
}

func (db *MongoDBdatabase) DeleteCartCourse(userID, courseID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	Collection := db.Collection("cart")

	_, err := Collection.DeleteOne(ctx, bson.M{"user_id": userID, "course_id": courseID})
	return err
}
