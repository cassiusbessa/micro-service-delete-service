package repositories

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteService(db string, id string) (bool, error) {
	collection := Repo.Client.Database(db).Collection("company")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	mongoId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		cancel()
		return false, err
	}
	update := bson.M{
		"$pull": bson.M{
			"services": bson.M{"_id": mongoId},
		},
	}
	result, err := collection.UpdateOne(ctx, bson.M{}, update)
	if err != nil {
		cancel()
		return false, err
	}
	if result.ModifiedCount == 0 {
		cancel()
		return false, nil
	}
	defer cancel()
	return true, nil
}
