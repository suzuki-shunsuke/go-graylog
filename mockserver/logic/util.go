package logic

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"

	log "github.com/sirupsen/logrus"
)

// ValidateObjectID validates ObjectID.
func ValidateObjectID(id string) error {
	if !bson.IsObjectIdHex(id) {
		return fmt.Errorf("id <%s> is invalid", id)
	}
	return nil
}

// LogWE outputs outputs log at Warn or Error level.
func LogWE(sc int, entry *log.Entry, msg string) {
	if sc >= 500 {
		entry.Error(msg)
	} else {
		entry.Warn(msg)
	}
}
