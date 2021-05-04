package sports

import (
	"time"

	"github.com/match-a-player/server/vendor/go.mongodb.org/mongo-driver/bson/primitive"
)

type Sport struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	Name      string             `bson:"string"`
}
