package store

import (
	"gopkg.in/mgo.v2/bson"
)

func NewObjectID() string {
	return bson.NewObjectId().Hex()
}
