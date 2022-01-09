package db

import (
	"context"
	"os"
	"time"

	"github.com/wgarcia1309/go-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindUsers(userID string, page int64, search string, tipo string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(os.Getenv("DATABASENAME"))
	col := db.Collection(os.Getenv("USER_COLLECTION"))

	var results []*models.User

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}

	ptr, err := col.Find(ctx, query, findOptions)
	if err != nil {
		return results, false
	}

	for ptr.Next(ctx) {
		var s models.User
		err := ptr.Decode(&s)
		if err != nil {
			return results, false
		}

		var r models.Relation
		r.UserID = userID
		r.UserRelationID = s.ID.Hex()

		include := false

		areRelated := GetRelation(r)
		switch tipo {
		case "N":
			if areRelated != nil {
				include = true
			}
		case "F":
			if areRelated == nil {
				include = true
			}
		case "A":
			include = true
		default:
			return nil, false
		}
		if r.UserRelationID == userID {
			include = false
		}

		if include {
			s.Password = ""
			s.Bio = ""
			s.Website = ""
			s.Location = ""
			s.Banner = ""
			s.Email = ""
			results = append(results, &s)
		}
	}

	err = ptr.Err()
	if err != nil {
		return results, false
	}
	ptr.Close(ctx)
	return results, true
}
