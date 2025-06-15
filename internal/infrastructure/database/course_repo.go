package database

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/bhushan-aruto/aspiration-matters-backend/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *MongoDBdatabase) AddCourse(course *entity.Course) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	collection := db.Collection("courses")
	_, err := collection.InsertOne(ctx, course)
	return err
}

func (db *MongoDBdatabase) GetAllTheCourses() ([]*entity.Course, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	collection := db.Collection("courses")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var courses []*entity.Course

	for cursor.Next(ctx) {
		var c entity.Course
		if err := cursor.Decode(&c); err == nil {
			courses = append(courses, &c)
		}
	}
	return courses, nil

}

func (db *MongoDBdatabase) DeleteCourseById(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	collection := db.Collection("courses")

	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})

	return err
}

func (db *MongoDBdatabase) GetCourseByID(id string) (*entity.Course, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("called my ame single course")
	collection := db.Collection("courses")

	var course entity.Course
	log.Println("Querying for course with ID:", id)

	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&course)
	if err != nil {
		return nil, err
	}

	return &course, nil
}

func (db *MongoDBdatabase) GetPurchasedCoursesByUserId(userID string) ([]*entity.Course, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	myLearningColl := db.Collection("my-learning")
	var learning struct {
		CourseIDs []string `bson:"course_ids"`
	}
	// err := myLearningColl.FindOne(ctx, bson.M{"user_id": userID}).Decode(&learning)
	// if err != nil {
	// 	return nil, fmt.Errorf("error fetching user learning data: %v", err)
	// }
	err := myLearningColl.FindOne(ctx, bson.M{"user_id": userID}).Decode(&learning)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {

			return []*entity.Course{}, nil
		}
		return nil, fmt.Errorf("error fetching user learning data: %v", err)
	}

	if len(learning.CourseIDs) == 0 {
		return []*entity.Course{}, nil
	}

	coursesColl := db.Collection("courses")
	filter := bson.M{"_id": bson.M{"$in": learning.CourseIDs}}

	cursor, err := coursesColl.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("error fetching courses: %v", err)
	}
	defer cursor.Close(ctx)

	var courses []*entity.Course
	for cursor.Next(ctx) {
		var course entity.Course
		if err := cursor.Decode(&course); err != nil {
			return nil, fmt.Errorf("error decoding course: %v", err)
		}
		courses = append(courses, &course)
	}

	return courses, nil
}

func (db *MongoDBdatabase) GetCoursesExcludingUserPurchased(userID string) ([]*entity.Course, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	myLearningCollection := db.Collection("my-learning")

	var myLearning struct {
		CourseIDs []string `bson:"course_ids"`
	}
	err := myLearningCollection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&myLearning)
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, err
	}

	courseCollection := db.Collection("courses")

	filter := bson.M{}
	if len(myLearning.CourseIDs) > 0 {
		filter = bson.M{
			"_id": bson.M{
				"$nin": myLearning.CourseIDs,
			},
		}
	}

	cursor, err := courseCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var courses []*entity.Course
	for cursor.Next(ctx) {
		var course entity.Course
		if err := cursor.Decode(&course); err == nil {
			courses = append(courses, &course)
		}
	}

	return courses, nil
}
