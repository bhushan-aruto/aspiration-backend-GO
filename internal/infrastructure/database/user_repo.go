package database

import (
	"context"
	"log"
	"time"

	"github.com/bhushan-aruto/aspiration-matters-backend/internal/entity"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/infrastructure/server/models"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *MongoDBdatabase) CheckUserByEmail(email string) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	Collection := db.Collection("users")

	var user entity.User
	err := Collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {

			return nil, nil
		}
		log.Println("Error finding user by email:", err)
		return nil, err
	}

	return &user, nil
}

func (db *MongoDBdatabase) SaveUser(user *entity.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	collection := db.Collection("users")
	_, err := collection.InsertOne(ctx, user)

	return err
}

func (db *MongoDBdatabase) StoreOTP(email, otp string, expiry time.Time) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	otpData := models.OTP{
		Email:     email,
		Otp:       otp,
		ExpiresAt: expiry,
	}
	collection := db.Collection("otps")
	_, err := collection.InsertOne(ctx, otpData)
	if err != nil {
		return err
	}

	go func(email, otp string) {
		time.Sleep(time.Minute * 5)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		collection := db.Collection("otps")
		filter := bson.M{
			"email": email, "otp": otp,
		}
		_, err := collection.DeleteOne(ctx, filter)
		if err != nil {
			log.Println("error occured while deleteing the otp ", err)
		}
	}(email, otp)

	return nil
}

func (db *MongoDBdatabase) VerifyOTP(email, otp string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var otpRecord models.OTP

	collection := db.Collection("otps")

	if err := collection.FindOne(ctx, bson.M{"email": email, "otp": otp}).Decode(&otpRecord); err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}
	if otpRecord.ExpiresAt.Before(time.Now()) {
		return false, nil
	}
	return true, nil
}

func (db *MongoDBdatabase) GetUserByID(id string) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.Collection("users")
	var user entity.User
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *MongoDBdatabase) MyLearningCollection(myLearing *entity.MyLearning) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	collection := db.Collection("my-learning")

	_, err := collection.InsertOne(ctx, myLearing)
	return err
}

func (db *MongoDBdatabase) UpdatePassword(email, newHashedPassword string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.Collection("users")

	hashedPassword, err := utils.HashPassword(newHashedPassword)
	if err != nil {
		return err
	}

	filter := bson.M{"email": email}
	update := bson.M{"$set": bson.M{"password": hashedPassword}}

	log.Println("Updating password for:", email)
	log.Println("New hashed password:", hashedPassword)

	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println("Error updating password:", err)
		return err
	}

	log.Println("Matched:", res.MatchedCount, "Modified:", res.ModifiedCount)

	return nil
}
