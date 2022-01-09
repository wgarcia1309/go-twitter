package db

import (
	"context"
	"os"
	"time"

	"github.com/wgarcia1309/go-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewUser(u models.User) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Hour)
	defer cancel()
	db := MongoCN.Database(os.Getenv("DATABASENAME"))
	collection := db.Collection(os.Getenv("USER_COLLECTION"))

	u.Password, _ = encrypt(u.Password)

	result, err := collection.InsertOne(ctx, u)
	if err != nil {
		return "", err
	}

	objectID, _ := result.InsertedID.(primitive.ObjectID)
	return objectID.String(), nil
}

func EmailExist(email string) (models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(os.Getenv("DATABASENAME"))
	collection := db.Collection(os.Getenv("USER_COLLECTION"))
	condition := bson.M{"email": email}

	var user models.User
	err := collection.FindOne(ctx, condition).Decode(&user)
	if err != nil {
		return user, false
	}
	return user, true
}

func UsernameExist(username string) (models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Hour)
	defer cancel()

	db := MongoCN.Database(os.Getenv("DATABASENAME"))
	collection := db.Collection(os.Getenv("USER_COLLECTION"))
	condition := bson.M{"username": username}

	var user models.User
	err := collection.FindOne(ctx, condition).Decode(&user)
	if err != nil {
		return user, false
	}
	return user, true
}

func GetUserProfile(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database(os.Getenv("DATABASENAME"))
	col := db.Collection(os.Getenv("USER_COLLECTION"))

	var UserProfile models.User
	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{"_id": objID}

	err := col.FindOne(ctx, condition).Decode(&UserProfile)
	UserProfile.Password = ""
	if err != nil {
		return UserProfile, err
	}
	return UserProfile, nil
}

func UpdateUser(u models.User, ID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(os.Getenv("DATABASENAME"))
	col := db.Collection(os.Getenv("USER_COLLECTION"))

	registro := make(map[string]interface{})
	if len(u.Name) > 0 {
		registro["name"] = u.Name
	}
	if len(u.Lastname) > 0 {
		registro["lastname"] = u.Lastname
	}
	if len(u.Birthdate.String()) > 0 {
		registro["birthdate"] = u.Birthdate
	}
	if len(u.Avatar) > 0 {
		registro["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		registro["banner"] = u.Banner
	}
	if len(u.Bio) > 0 {
		registro["bio"] = u.Bio
	}
	if len(u.Location) > 0 {
		registro["location"] = u.Location
	}
	if len(u.Website) > 0 {
		registro["website"] = u.Website
	}

	updtString := bson.M{
		"$set": registro,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filtro := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filtro, updtString)
	if err != nil {
		return err
	}

	return nil
}
