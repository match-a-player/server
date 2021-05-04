package user

import (
	"context"
	"fmt"
	"time"

	"github.com/match-a-player/server/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getUserbyID(id string, timeout uint, db database.DB) ([]User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	users := []User{}
	userCollection := db.GetCollection("users")
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	val, err := userCollection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("getUserbyID: %v", err)
	}
	if err := val.All(ctx, &users); err != nil {
		return nil, fmt.Errorf("getUserbyID: %v", err)
	}
	return users, nil
}

func getAllUsers(timeout uint, db database.DB) ([]User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	users := []User{}
	allUsers, err := db.GetCollection("users").Find(ctx, bson.D{}, nil)
	if err != nil {
		return nil, fmt.Errorf("getAllUsers: %v", err)
	}
	if err := allUsers.All(ctx, &users); err != nil {
		return nil, fmt.Errorf("getAllUsers: %v", err)
	}
	return users, nil
}
