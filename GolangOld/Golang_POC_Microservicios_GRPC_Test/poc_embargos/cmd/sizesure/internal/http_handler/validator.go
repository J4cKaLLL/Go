package http_handler

import "gopkg.in/go-playground/validator.v9"

// Validator dependencia de libreria go validator.v9
type Validator struct {
	validator *validator.Validate
}

// Validate realiza la validacion
func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

// NewValidator realiza la implementacion del validador
func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

