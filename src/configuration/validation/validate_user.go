package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	r "github.com/rozzettimatheus/crud-go/src/configuration/response_error"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

// start with the app
func init() {
	// cast
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		un := ut.New(en, en)
		transl, _ = un.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserErr(validation_err error) *r.ResponseError {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	// json unmarshall error
	if errors.As(validation_err, &jsonErr) {
		return r.BadRequestErr("Invalid field type")
	}
	// validation error
	if errors.As(validation_err, &jsonValidationError) {
		errorCauses := []r.Cause{}
		// validation error as string => cast to validationErrors
		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := r.Cause{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}
			errorCauses = append(errorCauses, cause)
		}
		return r.ValidationErr("Some fields are invalid", errorCauses)
	}
	return r.BadRequestErr("Error trying to parse fields")
}
