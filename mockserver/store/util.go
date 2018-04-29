package store

import (
	"gopkg.in/mgo.v2/bson"
)

// NewObjectID returns a new ObjectId.
func NewObjectID() string {
	return bson.NewObjectId().Hex()
}
