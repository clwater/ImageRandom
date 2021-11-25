package google_earth

import (
	"ImageRandom/db"
	"ImageRandom/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func Random() string {
	var (
		cursor *mongo.Cursor
		err    error
	)
	collection := db.GetCollection()



	pipeline := mongo.Pipeline{
		{
			{"$sample", bson.D{
				{"size", 1},
			}},
		},
	}

	cursor, _ = collection.Aggregate(context.TODO(), pipeline)

	var results []model.GoogleEarth
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
		return ""
	}

	return results[0].Url
}

