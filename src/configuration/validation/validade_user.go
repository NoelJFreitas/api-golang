package validation

import (
	"encoding/json"
	"errors"

	"github.com/NoelJFreitas/api-golang/src/configuration/rest_errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	translator ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		translator, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, translator)
	}
}

func ValidateUseError(validation_error error) *rest_errors.RestError {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_error, &jsonErr) {
		// caso o erro seja de Unmarshal
		return rest_errors.NewBadRequestError("Invalid field type")
	} else if errors.As(validation_error, &jsonValidationError) {
		// caso o erro seja de validação de campos
		errorCauses := []rest_errors.Causes{}

		for _, e := range validation_error.(validator.ValidationErrors) {
			cause := rest_errors.Causes{
				Message: e.Translate(translator),
				Field:   e.Field(),
			}
			errorCauses = append(errorCauses, cause)
		}
		return rest_errors.NewBadRequestValidationError("Some fields are invalid", errorCauses)
	}
	// caso o erro seja desconhecido
	return rest_errors.NewBadRequestError("Error trying to convert fields")
}
