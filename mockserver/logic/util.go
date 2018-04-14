package logic

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

func ValidateObjectID(id string) error {
	if !bson.IsObjectIdHex(id) {
		return fmt.Errorf("id <%s> is invalid", id)
	}
	return nil
}
