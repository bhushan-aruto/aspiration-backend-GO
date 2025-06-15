package database

import (
	"context"
	"time"

	"github.com/bhushan-aruto/aspiration-matters-backend/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (db *MongoDBdatabase) GetPurchaseHistoryByUserId(userID string) ([]*entity.PurchaseHistory, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := db.Collection("purchase-History")
	opts := options.Find().SetSort(bson.D{{Key: "date", Value: -1}})

	cursor, err := collection.Find(ctx, bson.M{
		"user_id": userID,
	}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var purchaseHistory []*entity.PurchaseHistory
	
	for cursor.Next(ctx) {
		var purchase entity.PurchaseHistory
		if err := cursor.Decode(&purchase); err != nil {
			return nil, err
		}
		purchaseHistory = append(purchaseHistory, &purchase)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return purchaseHistory, nil
}
