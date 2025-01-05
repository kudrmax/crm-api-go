package my_errors

import "errors"

var (
	ContactNameIsRequiredErr = errors.New("contact name is required")
	ContactNotFoundErr       = errors.New("contact not found")
	NameAlreadyUsedErr       = errors.New("contact name already used")
)
