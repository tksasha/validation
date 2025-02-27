package validation_test

import (
	"testing"
	"time"

	"github.com/tksasha/validation"
	"gotest.tools/v3/assert"
)

func TestValidationPresence(t *testing.T) {
	t.Run("add error for blank value", func(t *testing.T) {
		validation := validation.New()

		_ = validation.Presence("name", "")

		assert.Error(t, validation.Errors, "name: is required")
	})

	t.Run("passes when value is present", func(t *testing.T) {
		validation := validation.New()

		_ = validation.Presence("name", "Bruce Wayne")

		assert.Assert(t, !validation.Errors.Exists())
	})
}

func TestValidationInteger(t *testing.T) {
	t.Run("add error for invalid value", func(t *testing.T) {
		validation := validation.New()

		res := validation.Integer("age", "abc")

		assert.Equal(t, res, 0)
		assert.Error(t, validation.Errors, "age: is invalid")
	})

	t.Run("converts string to integer when value is valid", func(t *testing.T) {
		validation := validation.New()

		res := validation.Integer("age", "33")

		assert.Equal(t, res, 33)
		assert.Assert(t, !validation.Errors.Exists())
	})
}

func TestValidationFormula(t *testing.T) {
	t.Run("add error for blank value", func(t *testing.T) {
		validation := validation.New()

		formula, res := validation.Formula("formula", "")

		assert.Equal(t, formula, "")
		assert.Equal(t, res, 0.0)
		assert.Error(t, validation.Errors, "formula: is required")
	})

	t.Run("add error for invalid value", func(t *testing.T) {
		validation := validation.New()

		formula, res := validation.Formula("formula", "abc")

		assert.Equal(t, formula, "abc")
		assert.Equal(t, res, 0.0)
		assert.Error(t, validation.Errors, "formula: is invalid")
	})

	t.Run("calculate result when value is valid", func(t *testing.T) {
		validation := validation.New()

		formula, res := validation.Formula("formula", "2+3")

		assert.Equal(t, formula, "2+3")
		assert.Equal(t, res, 5.0)
		assert.Assert(t, !validation.Errors.Exists())
	})
}

func TestValidationBoolean(t *testing.T) {
	t.Run("returns true for 'true' string", func(t *testing.T) {
		validation := validation.New()

		res := validation.Boolean("visible", "true")

		assert.Equal(t, res, true)
		assert.Assert(t, !validation.Errors.Exists())
	})

	t.Run("returns false for 'false' string", func(t *testing.T) {
		validation := validation.New()

		res := validation.Boolean("visible", "false")

		assert.Equal(t, res, false)
		assert.Assert(t, !validation.Errors.Exists())
	})

	t.Run("returns false for empty string", func(t *testing.T) {
		validation := validation.New()

		res := validation.Boolean("visible", "")

		assert.Equal(t, res, false)
		assert.Assert(t, !validation.Errors.Exists())
	})

	t.Run("adds error for invalid value", func(t *testing.T) {
		validation := validation.New()

		res := validation.Boolean("visible", "xxx")

		assert.Equal(t, res, false)
		assert.Error(t, validation.Errors, "visible: is invalid")
	})
}

func TestValidationDate(t *testing.T) {
	t.Run("returns error when value is blank", func(t *testing.T) {
		validation := validation.New()

		res := validation.Date("date", "")

		assert.Assert(t, res.IsZero())
		assert.Error(t, validation.Errors, "date: is required")
	})

	t.Run("returns error when value is invalid", func(t *testing.T) {
		validation := validation.New()

		res := validation.Date("date", "abc")

		assert.Assert(t, res.IsZero())
		assert.Error(t, validation.Errors, "date: is invalid")
	})

	t.Run("parse date when value is valid", func(t *testing.T) {
		validation := validation.New()

		res := validation.Date("date", "2025-02-04")

		assert.Equal(t, res, date(t, "2025-02-04"))
		assert.Assert(t, !validation.Errors.Exists())
	})
}

func date(t *testing.T, value string) time.Time {
	t.Helper()

	date, err := time.Parse(time.DateOnly, value)
	if err != nil {
		t.Fatalf("failed to parse date: %v", err)
	}

	return date
}
