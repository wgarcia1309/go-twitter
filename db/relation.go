package db

import (
	"context"
	"os"
	"time"

	"github.com/wgarcia1309/go-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

func Follow(t models.Relation) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(os.Getenv("DATABASENAME"))
	col := db.Collection(os.Getenv("RELATION_COLLECTION"))

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return err
	}

	return nil
}

func Unfollow(t models.Relation) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(os.Getenv("DATABASENAME"))
	col := db.Collection(os.Getenv("RELATION_COLLECTION"))

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return err
	}
	return nil
}

func GetRelation(t models.Relation) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(os.Getenv("DATABASENAME"))
	col := db.Collection(os.Getenv("RELATION_COLLECTION"))

	condicion := bson.M{
		"userID":         t.UserID,
		"userRelationID": t.UserRelationID,
	}

	var result models.Relation
	err := col.FindOne(ctx, condicion).Decode(&result)
	if err != nil {
		return err
	}
	return nil
}
