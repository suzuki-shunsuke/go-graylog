package validator

import (
	"regexp"

	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/mgo.v2/bson"

	log "github.com/sirupsen/logrus"
)

var (
	// CreateValidator validates parameters of Create APIs.
	CreateValidator *validator.Validate
	// UpdateValidator validates parameters of Update APIs.
	UpdateValidator   *validator.Validate
	indexPrefixRegexp *regexp.Regexp
)

func init() {
	indexPrefixRegexp = regexp.MustCompile(`^[a-z0-9][a-z0-9_+-]*$`)
	CreateValidator = validator.New()
	CreateValidator.SetTagName("v-create")
	UpdateValidator = validator.New()
	UpdateValidator.SetTagName("v-update")
	validators := map[string]validator.Func{
		"indexprefixregexp": ValidateIndexPrefixRegexp,
		"objectid":          ValidateObjectID,
	}
	for k, v := range validators {
		if err := CreateValidator.RegisterValidation(k, v); err != nil {
			log.Fatal(err)
		}
		if err := UpdateValidator.RegisterValidation(k, v); err != nil {
			log.Fatal(err)
		}
	}
}

// ValidateIndexPrefixRegexp validates index prefix's pattern.
func ValidateIndexPrefixRegexp(lf validator.FieldLevel) bool {
	return indexPrefixRegexp.MatchString(lf.Field().String())
}

// ValidateObjectID validates objectID
// https://docs.mongodb.com/manual/reference/bson-types/#objectid
func ValidateObjectID(lf validator.FieldLevel) bool {
	return bson.IsObjectIdHex(lf.Field().String())
}
