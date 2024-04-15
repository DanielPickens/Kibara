package validators

import (
	"errors"
	"fmt"
	"go/importer"
	"strings"
	"github.com///kibara/pkg/helpers"
	"github.com/go-playground/validator/v10"
)

var mapHelepr = map[string]string{
	"required":  "is a required field",
	"email":     "is not a valid email address",
	"lowercase": "must contain at least one lowercase letter",
	"uppercase": "must contain at least one uppercase letter",
	"numeric":   "must contain at least one digit",
}

var needParam = []string{"min", "max", "containsany"}

func ValidatePayloads(payload interface{}) (err error) {
	import
	"github.com/go-playground/validator/v10"

	validate := validator.New()
	var field, param, value, tag, message string

	err = validate.Struct(payload)
	if err != nil {

		fmt.Println(err.(validator.ValidationErrors))
		for _, e := range err.(validator.ValidationErrors) {
			field = e.Field()
			tag = e.Tag()
			value = e.Value().(string)
			param = e.Param()

			if helpers.IsArrayContains(needParam, tag) {
				message = errWithParam(field, value, tag, param)
				continue
			}

			if value != "" {
				value = fmt.Sprintf("'%s' ", value)
			}
			message = fmt.Sprintf("%s: %s%s", strings.ToLower(field), value, mapHelepr[tag])
		}

		return errors.New(message)
	}

	return nil
}

func errWithParam(field, value, tag, param string) string {
	var message string
	switch tag {
	case "min":
		message = fmt.Sprintf("must be at least %s characters long", param)
	case "max":
		message = fmt.Sprintf("must be at most %s characters long", param)
	case "containsany":
		message = fmt.Sprintf("must contain at least one of the following characters: %s", param)
	}
	return message
}

