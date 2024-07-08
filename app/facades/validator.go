package facades

import (
  "fmt"

  "github.com/go-playground/validator/v10"
)

var validate = validator.New()

func Validate(data any) (messages []string) {
  if errs := validate.Struct(data); errs != nil {
    for _, err := range errs.(validator.ValidationErrors) {
      message := fmt.Sprintf("[%s]: '%v' | Needs to implement '%s'", err.Field(), err.Value(), err.Tag())
      messages = append(messages, message)
    }
  }

  return
}
