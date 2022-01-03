package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DevulevoTweets struct {
	ID      primitive.ObjectID `bson:"_id," json:"_id,omitempty"`
	UserID  string             `bson:"userId" json:"userId,omitempty"`
	Mensaje string             `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha   time.Time          `bson:"fecha" json:"fecha,omitempty"`
}
