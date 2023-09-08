package errors

import "errors"

// GetEmptyError get all empty errors
func GetEmptyError(resource string) error {
	return errors.New("The " + resource + " variable is empty")
}
