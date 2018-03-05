package graylog

import (
	"gopkg.in/go-playground/validator.v9"
)

var CreateValidator *validator.Validate
var UpdateValidator *validator.Validate

func init() {
	CreateValidator = validator.New()
	CreateValidator.SetTagName("v-create")
	UpdateValidator = validator.New()
	UpdateValidator.SetTagName("v-update")
}
