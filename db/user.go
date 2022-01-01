package db

import (
	"context"
	"os"
	"time"

	"github.com/wgarcia1309/go-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewUser(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database(os.Getenv("DATABASENAME"))
	collection := db.Collection("USER_COLLECTION")

	u.Password, _ = encrypt(u.Password)

	result, err := collection.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	objectID, _ := result.InsertedID.(primitive.ObjectID)
	return objectID.String(), true, nil
}

func EmailExist(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(os.Getenv("DATABASENAME"))
	collection := db.Collection("USER_COLLECTION")
	condition := bson.M{"email": email}

	var user models.User
	err := collection.FindOne(ctx, condition).Decode(&user)
	ID := user.ID.Hex()
	if err != nil {
		return user, false, ID
	}
	return user, true, ID
}

func UsernameExist(username string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(os.Getenv("DATABASENAME"))
	collection := db.Collection("USER_COLLECTION")
	condition := bson.M{"username": username}

	var user models.User
	err := collection.FindOne(ctx, condition).Decode(&user)
	ID := user.ID.Hex()
	if err != nil {
		return user, false, ID
	}
	return user, true, ID
}
