package repositories

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// func UpdateService(db string, id string, service entities.Service) (bool, error) {
// 	collection := Repo.Client.Database(db).Collection("company")
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

// 	mongoId, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		cancel()
// 		return false, err
// 	}

// 	filter := bson.M{
// 		"services._id": mongoId,
// 	}
// 	arrayFilter := bson.M{"elem._id": mongoId}
// 	update := bson.M{
// 		"$set": bson.M{
// 			"services.$[elem]": service,
// 		},
// 	}
// 	updateOpts := options.Update().SetArrayFilters(options.ArrayFilters{
// 		Filters: []interface{}{arrayFilter},
// 	})

// 	result, err := collection.UpdateOne(ctx, filter, update, updateOpts)
// 	if err != nil {
// 		cancel()
// 		return false, err
// 	}
// 	if result.MatchedCount == 0 {
// 		cancel()
// 		return false, nil
// 	}

// 	defer cancel()
// 	return true, nil
// }

func DeleteService(db string, id string) (bool, error) {
	collection := Repo.Client.Database(db).Collection("company")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	mongoId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		cancel()
		return false, err
	}

	// filter := bson.M{
	// 	"services._id": mongoId,
	// }
	// arrayFilter := bson.M{"elem._id": mongoId}
	update := bson.M{
		"$pull": bson.M{
			"services": bson.M{"_id": mongoId},
		},
	}
	// updateOpts := options.Update().SetArrayFilters(options.ArrayFilters{
	// 	Filters: []interface{}{arrayFilter},
	// })

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
