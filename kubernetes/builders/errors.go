package builders

import "errors"

var (
	errorExternalNameEmpty = errors.New("The ExternalName variable is empty")
	errorSelectorEmpty     = errors.New("The Selector variable is empty")
	errorPortsEmpty        = errors.New("The Ports variable is empty")
)
