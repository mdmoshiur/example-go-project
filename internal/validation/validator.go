package validation

import (
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

// use a single instance , it caches struct info
var (
	uni        *ut.UniversalTranslator
	Validator  *validator.Validate
	Translator ut.Translator
)

// InitValidator initialize the validator package
func InitValidator() error {
	eng := en.New()
	uni = ut.New(eng, eng)

	Translator, _ = uni.GetTranslator("en")
	Validator = validator.New(validator.WithRequiredStructEnabled())

	Validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		// skip if tag key says it should be ignored
		if name == "-" {
			return ""
		}

		return name
	})

	// override translation

	// override required_unless translation
	err := Validator.RegisterTranslation("required_unless", Translator, func(ut ut.Translator) error {
		return ut.Add("required_unless", "{0} is a required field", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required_unless", fe.Field())

		return t
	})
	if err != nil {
		return err
	}

	// override required_without_all translation
	err = Validator.RegisterTranslation("required_without_all", Translator, func(ut ut.Translator) error {
		return ut.Add("required_without_all", "{0} is a required field without dependent fields", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required_without_all", fe.Field())

		return t
	})
	if err != nil {
		return err
	}

	err = en_translations.RegisterDefaultTranslations(Validator, Translator)
	if err != nil {
		return err
	}

	return nil
}

// RemovePrefixStructName removes errors struct name from keys
func RemovePrefixStructName(fieldErrs map[string]string) map[string]string {
	errs := make(map[string]string, len(fieldErrs))
	for field, err := range fieldErrs {
		errs[field[strings.Index(field, ".")+1:]] = err
	}

	return errs
}
