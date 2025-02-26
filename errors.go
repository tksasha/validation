package validator

import (
	"fmt"
	"strings"

	"github.com/tksasha/xstrings"
)

type Errors map[string][]string //nolint:errname

func (e Errors) Set(attribute, message string) {
	attribute = xstrings.ToSnakeCase(attribute)

	e[attribute] = append(e[attribute], message)
}

func (e Errors) Get(attribute string) *string {
	messages := e[attribute]

	if len(messages) == 0 {
		return nil
	}

	message := fmt.Sprintf("%s: %s", attribute, strings.Join(messages, ", "))

	return &message
}

func (e Errors) Has(attribute string) bool {
	return len(e[attribute]) > 0
}

func (e Errors) Error() string {
	errors := []string{}

	for attribute := range e {
		errors = append(errors, *e.Get(attribute))
	}

	return strings.Join(errors, "; ")
}

func (e Errors) exists() bool {
	return len(e) != 0
}
