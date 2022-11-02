package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Template struct {
	ID       primitive.ObjectID `bson:"_id"`
	BodyPage string             `bson:"body_page" json:"body_page"`
}
