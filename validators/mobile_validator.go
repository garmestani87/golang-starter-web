package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func MobileValidator(fl validator.FieldLevel) bool {
	mobile, ok := fl.Field().Interface().(string)
	if ok {
		ok, _ := regexp.MatchString(`^(00989|\+989|09)(\d{9})$`,mobile)
		return ok
	}
	return false
}
