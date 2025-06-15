package database

import (
	"context"
	"log"
	"time"

	"github.com/bhushan-aruto/aspiration-matters-backend/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (db *MongoDBdatabase) AddBlog(blog *entity.Blog) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	collection := db.Collection("blogs-section")
	_, err := collection.InsertOne(ctx, blog)
	return err
}

// func (db *MongoDBdatabase) GetAllBlogs() ([]*entity.Blog, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
// 	defer cancel()

// 	collection := db.Collection("blogs-section")

// 	cursor, err := collection.Find(ctx, bson.M{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer cursor.Close(ctx)

// 	var blogs []*entity.Blog

// 	for cursor.Next(ctx) {
// 		var blog entity.Blog
// 		if err := cursor.Decode(&blog); err != nil {
// 			log.Printf("Error decoding image: %v", err)
// 			continue
// 		}
// 		blogs = append(blogs, &blog)

// 	}
// 	return blogs, nil

// }

func (db *MongoDBdatabase) GetAllBlogs() ([]*entity.Blog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := db.Collection("blogs-section")

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var blogs []*entity.Blog
	for cursor.Next(ctx) {
		var blog entity.Blog
		if err := cursor.Decode(&blog); err != nil {
			log.Printf("Error decoding blog: %v", err)
			continue
		}
		blogs = append(blogs, &blog)
	}

	return blogs, nil
}

func (db *MongoDBdatabase) GetBlogById(id string) (*entity.Blog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	collection := db.Collection("blogs-section")
	var blog *entity.Blog
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&blog)
	return blog, err

}

func (db *MongoDBdatabase) DeleteBlogByFileName(fileName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	collection := db.Collection("blogs-section")
	_, err := collection.DeleteOne(ctx, bson.M{
		"file_name": fileName,
	})
	return err
}
