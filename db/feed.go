package db

import (
	"context"
	"os"
	"time"

	"github.com/wgarcia1309/go-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetFeed(userID string, page int) ([]models.FollowerTweetRetrieved, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(os.Getenv("DATABASENAME"))
	col := db.Collection(os.Getenv("RELATION_COLLECTION"))

	skip := (page - 1) * 20

	conditions := make([]bson.M, 0)
	conditions = append(conditions, bson.M{"$match": bson.M{"userID": userID}})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         os.Getenv("TWEETS_COLLECTION"),
			"localField":   os.Getenv("LOCAL_FIELD"),
			"foreignField": os.Getenv("FOREIGN_FIELD"),
			"as":           "tweet",
		}})
	conditions = append(conditions, bson.M{"$unwind": "$tweet"})
	conditions = append(conditions, bson.M{"$sort": bson.M{"tweet.Date": -1}})
	conditions = append(conditions, bson.M{"$skip": skip})
	conditions = append(conditions, bson.M{"$limit": 20})

	ptr, err := col.Aggregate(ctx, conditions)
	if err != nil {
		return nil, false
	}
	var results []models.FollowerTweetRetrieved
	err = ptr.All(ctx, &results)
	if err != nil {
		return results, false
	}
	return results, true
}
