package db

import (
	"context"
	"os"
	"time"

	"github.com/wgarcia1309/go-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewTweet(t models.Tweet) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(os.Getenv("DATABASENAME"))
	col := db.Collection(os.Getenv("TWEETS_COLLECTION"))

	data := bson.M{
		"userID":  t.UserID,
		"message": t.Message,
		"date":    t.Date,
	}
	_, err := col.InsertOne(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func GetTweets(userID string, page int64) ([]*models.TweetRetrieved, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database(os.Getenv("DATABASENAME"))
	col := db.Collection(os.Getenv("TWEETS_COLLECTION"))

	var tweets []*models.TweetRetrieved

	condition := bson.M{
		"userID": userID,
	}
	var pageItems int64 = 20
	optionsMongo := options.Find()
	optionsMongo.SetLimit(pageItems)
	optionsMongo.SetSort(bson.D{{Key: "date", Value: -1}})
	optionsMongo.SetSkip((page - 1) * pageItems)

	ptr, err := col.Find(ctx, condition, optionsMongo)
	if err != nil {
		return tweets, false
	}

	for ptr.Next(context.TODO()) {

		var record models.TweetRetrieved
		err := ptr.Decode(&record)
		if err != nil {
			return tweets, false
		}
		tweets = append(tweets, &record)
	}
	return tweets, true
}

func DeleteTweet(tweetID string, userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database(os.Getenv("DATABASENAME"))
	col := db.Collection(os.Getenv("TWEETS_COLLECTION"))

	objID, _ := primitive.ObjectIDFromHex(tweetID)

	condition := bson.M{
		"_id":    objID,
		"userID": userID,
	}
	_, err := col.DeleteOne(ctx, condition)
	return err
}
