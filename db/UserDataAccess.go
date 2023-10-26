package db

import (
	"context"

	"github.com/brayanzuritadev/citas/models"
	"github.com/brayanzuritadev/citas/tools"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ReviewExistUser(email string) (models.User, bool, string) {
	ctx := context.TODO()

	db := MongoCN.Database(DatabaseName)

	col := db.Collection("User")

	condition := bson.M{"email": email}

	var user models.User
	err := col.FindOne(ctx, condition).Decode(user)

	ID := user.ID.Hex()
	if err != nil {
		return user, false, ID
	}

	return user, true, ID
}

func InsertUser(u models.User) (string, bool, error) {
	ctx := context.TODO()

	db := MongoCN.Database(DatabaseName)

	col := db.Collection("User")

	u.Password, _ = tools.PasswordEncrypt(u.Password)

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.String(), true, nil
}
