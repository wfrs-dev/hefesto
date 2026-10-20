package tools

import (
	"github.com/go-playground/validator/v10"
)

type PlaygroundValidator struct {
	validator *validator.Validate
}

func (v *PlaygroundValidator) Validate(i any) error {
	return v.validator.Struct(i)
}
