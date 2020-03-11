package erroraggregator

import (
	"errors"
	"strings"
)

const defaultDivider = `; `

// Aggregator handle an errors list
type Aggregator struct {
	errors []string
}

// New create a new Aggregator struct instance
func New() Aggregator {
	return Aggregator{}
}

// Append a new error to Aggregator
func (a *Aggregator) Append(err error) {
	a.errors = append(a.errors, err.Error())
}

// GetErrorMessages return all errors formatted
func (a *Aggregator) GetErrorMessages() error {

	if !a.GotErrors() {
		return nil
	}

	errorMessages := strings.Join(a.errors[:], defaultDivider)

	return errors.New(errorMessages)
}

// GotErrors return true for errors existence
func (a *Aggregator) GotErrors() bool {
	return len(a.errors) > 0
}
