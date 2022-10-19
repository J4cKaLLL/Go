package middleware

import "gopkg.in/go-playground/validator.v9"

// Validator dependencia de la librería gopkg.in/go-playground/validator.v9
type Validator struct {
	validator *validator.Validate
}

// Validate realiza la validación
func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

// NewValidator realiza la implementación del validador
func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}
