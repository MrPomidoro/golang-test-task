package validate

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/golang-test-task/pkg/common/errors"
	"reflect"
	"strings"
	"sync"
)

var invalidEmail = errors.New("invalid email")

func ValidateEmail(email string) (string, error) {
	if !IsEmail(email) {
		return "", invalidEmail
	}
	email = strings.ToLower(email)
	return email, nil
}

type DefaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

var _ binding.StructValidator = &DefaultValidator{}

// ValidateStruct validates the input struct using the defaultValidator.
//
// obj: the struct to validate.
//
// error: an error indicating validation failure. Nil if validation succeeds.
func (v *DefaultValidator) ValidateStruct(obj interface{}) error {

	if kindOfData(obj) == reflect.Struct {

		v.lazyinit()

		if err := v.validate.Struct(obj); err != nil {
			return err
		}
	}

	return nil
}

func (v *DefaultValidator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

func (v *DefaultValidator) lazyinit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("binding")
	})
}

func kindOfData(data interface{}) reflect.Kind {

	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}
