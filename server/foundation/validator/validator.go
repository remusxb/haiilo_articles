package validator

import (
	"reflect"
	"strings"
	"sync"

	"github.com/go-playground/validator/v10"
)

type (
	// AppValidatorInterface is a wrapper over go validator
	AppValidatorInterface interface {
		Validate(i interface{}) error
		RegisterValidation(tagValidator TagValidatorInterface, callValidationEvenIfNull ...bool) error
		RegisterTagNameFunc(fn validator.TagNameFunc)
	}

	// TagValidatorInterface is used for custom validators
	TagValidatorInterface interface {
		Tag() string
		Validate(fl validator.FieldLevel) bool
	}

	AppValidator struct {
		validate validator.Validate
	}
)

var (
	once   sync.Once
	client *AppValidator
)

func GetInstance() *AppValidator {
	if client != nil {
		return client
	}

	once.Do(func() {
		client = New()
	})

	return client
}

func New() *AppValidator {
	newValidator := validator.New()

	// Use json tag names for errors instead of go struct names.
	newValidator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return &AppValidator{validate: *newValidator}
}

func (validator AppValidator) Validate(i interface{}) error {
	return validator.validate.Struct(i)
}

func (validator AppValidator) RegisterValidation(tagValidator TagValidatorInterface, callValidationEvenIfNull ...bool) error {
	return validator.validate.RegisterValidation(tagValidator.Tag(), tagValidator.Validate, callValidationEvenIfNull...)
}

func (validator AppValidator) RegisterTagNameFunc(fn validator.TagNameFunc) {
	validator.validate.RegisterTagNameFunc(fn)
}
