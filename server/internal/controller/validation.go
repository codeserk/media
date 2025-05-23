package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
)

func ValidateRequest[E any](w http.ResponseWriter, r *http.Request, v *validator.Validate, input E) error {
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		ValidationError(w, fmt.Errorf("invalid input: %v", err))
		return err
	}
	err = v.Struct(input)
	if err != nil {
		ValidationError(w, err)
		return err
	}

	return nil
}

var decoder = schema.NewDecoder()

func ValidateQuery[TInput any](w http.ResponseWriter, r *http.Request, v *validator.Validate, input TInput) error {
	if err := decoder.Decode(input, r.URL.Query()); err != nil {
		ValidationError(w, fmt.Errorf("decode: %v (%v)", err, r.URL.Query()))
		return err
	}
	if err := v.Struct(input); err != nil {
		ValidationError(w, err)
		return err
	}

	return nil
}
