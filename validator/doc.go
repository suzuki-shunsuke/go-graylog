/*
Package validator provides validators for graylog's create and update APIs.
validator uses gopkg.in/go-playground/validator.v9 .

https://godoc.org/gopkg.in/go-playground/validator.v9

  role := &graylog.Role{}
  if err := validator.CreateValidator.Struct(role); err != nil {
  	return "", nil, err
  }
*/
package validator
