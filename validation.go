package validation

import (
	"strconv"
	"time"

	"github.com/shopspring/decimal"
	"github.com/tksasha/calculator"
)

const (
	required = "is required"
	invalid  = "is invalid"
)

type Validation struct {
	Errors Errors
}

func New() *Validation {
	return &Validation{
		Errors: Errors{},
	}
}

func (v *Validation) Presence(attribute, value string) string {
	if value == "" {
		v.Errors.Set(attribute, required)
	}

	return value
}

func (v *Validation) Integer(attribute, value string) int {
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

func (v *Validation) Formula(attribute, formula string) (string, decimal.Decimal) {
	if formula == "" {
		v.Errors.Set(attribute, required)

		return formula, decimal.NewFromInt(0)
	}

	sum, err := calculator.Calculate(formula)
	if err != nil {
		v.Errors.Set(attribute, invalid)

		return formula, decimal.NewFromInt(0)
	}

	return formula, sum
}

func (v *Validation) Boolean(attribute, value string) bool {
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

func (v *Validation) Date(attribute, value string) time.Time {
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
