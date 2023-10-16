package httpserver

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
)

type ValidationErrDetail struct {
	Field   string      `json:"field"`
	Value   interface{} `json:"value"`
	Message string      `json:"message"`
}

func ValidationErrorDetails(obj interface{}, tag string, errs validator.ValidationErrors) []*ValidationErrDetail {
	if len(errs) == 0 {
		return []*ValidationErrDetail{}
	}
	var errors []*ValidationErrDetail
	e := reflect.TypeOf(obj).Elem()
	for _, err := range errs {
		f, _ := e.FieldByName(err.Field())
		tagName, _ := f.Tag.Lookup(tag)
		val := err.Value()
		var message string

		switch err.ActualTag() {
		case "required":
			message = fmt.Sprintf("required tag='%s'", tagName)
		case "min":
			message = fmt.Sprintf("%s required at least %s length", tagName, err.Param())
		case "max":
			message = fmt.Sprintf("%s required at no longer %s length", tagName, err.Param())
		case "len":
			message = fmt.Sprintf("%s required exactly %s length", tagName, err.Param())
		case "uuid4":
			message = fmt.Sprintf("%s type must be uuid", tagName)
		default:
			message = fmt.Sprintf("invalid %s", tagName)
		}

		errors = append(errors, &ValidationErrDetail{
			Field:   tagName,
			Value:   val,
			Message: message,
		})
	}
	return errors
}

func NewValidationErrorDetails(field, message string, value interface{}) []*ValidationErrDetail {
	return []*ValidationErrDetail{
		{
			Field:   field,
			Value:   value,
			Message: message,
		},
	}
}
