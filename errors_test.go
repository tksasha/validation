package validation_test

import (
	"testing"

	"github.com/tksasha/validation"
	"gotest.tools/v3/assert"
)

func TestGet(t *testing.T) {
	t.Run("when error exists", func(t *testing.T) {
		actual := build(t).Get("name")

		expected := "name: is required, must be unique"

		assert.Equal(t, *actual, expected)
	})

	t.Run("when error does not exist", func(t *testing.T) {
		message := build(t).Get("amount")

		assert.Assert(t, message == nil)
	})
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

func TestExists(t *testing.T) {
	t.Run("when empty", func(t *testing.T) {
		errors := validation.Errors{}

		assert.Assert(t, !errors.Exists())
	})

	t.Run("when not empty", func(t *testing.T) {
		errors := validation.Errors{}

		errors.Set("name", "is required")

		assert.Assert(t, errors.Exists())
	})
}

func build(t *testing.T) validation.Errors {
	t.Helper()

	errors := validation.Errors{}

	errors.Set("age", "must be greater than 18")

	errors.Set("name", "is required")
	errors.Set("name", "must be unique")

	return errors
}
