package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID   `bson:"_id"`
	CreatedAt time.Time            `bson:"created_at"`
	UpdatedAt time.Time            `bson:"updated_at"`
	FirstName string               `bson:"first_name"`
	LastName  string               `bson:"last_name"`
	UserName  string               `bson:"user_name"`
	Location  string               `bson:"location"`
	Age       int                  `bson:"age"`
	Sports    []primitive.ObjectID `bson:"sports"`
}
