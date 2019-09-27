package common

import "fmt"

// Errors is a collection of errors.
type Errors []error

func (errs Errors) Error() string {
	msgs := make([]string, len(errs))
	for i, err := range errs {
		msgs[i] = err.Error()
	}
	return fmt.Sprint(msgs)
}
