package validator

import (
	"strconv"
	"time"

	"github.com/tksasha/calculator"
)

const (
	required = "is required"
	invalid  = "is invalid"
)

type Validator struct {
	Errors Errors
}

func New() *Validator {
	return &Validator{
		Errors: Errors{},
	}
}

func (v *Validator) HasErrors() bool {
	return v.Errors.exists()
}

func (v *Validator) Set(attribute, value string) {
	v.Errors.Set(attribute, value)
}

func (v *Validator) Presence(attribute, value string) string {
	if value == "" {
		v.Errors.Set(attribute, required)
	}

	return value
}

func (v *Validator) Integer(attribute, value string) int {
	if value == "" {
		return 0
	}

	digit, err := strconv.Atoi(value)
	if err != nil {
		v.Errors.Set(attribute, invalid)

		return 0
	}

	return digit
}

func (v *Validator) Formula(attribute, formula string) (string, float64) {
	if formula == "" {
		v.Errors.Set(attribute, required)

		return formula, 0.0
	}

	sum, err := calculator.Calculate(formula)
	if err != nil {
		v.Errors.Set(attribute, invalid)

		return formula, 0.0
	}

	return formula, sum
}

func (v *Validator) Boolean(attribute, value string) bool {
	switch value {
	case "true":
		return true
	case "false", "":
		return false
	default:
		v.Errors.Set(attribute, invalid)

		return false
	}
}

func (v *Validator) Date(attribute, value string) time.Time {
	if value == "" {
		v.Errors.Set(attribute, required)

		return time.Time{}
	}

	date, err := time.Parse(time.DateOnly, value)
	if err != nil {
		v.Errors.Set(attribute, invalid)

		return time.Time{}
	}

	return date
}
