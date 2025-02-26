package validator_test

import (
	"testing"

	"github.com/tksasha/validator"
	"gotest.tools/v3/assert"
)

func TestGet(t *testing.T) {
	actual := build(t).Get("name")

	expected := "name: is required, must be unique"

	assert.Equal(t, actual, expected)
}

func TestError(t *testing.T) {
	err := build(t)

	assert.Error(t, err, "age: must be greater than 18; name: is required, must be unique")
}

func TestHas(t *testing.T) {
	errors := build(t)

	assert.Assert(t, errors.Has("age"))
	assert.Assert(t, errors.Has("name"))
	assert.Assert(t, !errors.Has("amount"))
}

func build(t *testing.T) validator.Errors {
	t.Helper()

	errors := validator.Errors{}

	errors.Set("age", "must be greater than 18")

	errors.Set("name", "is required")
	errors.Set("name", "must be unique")

	return errors
}
