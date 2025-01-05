package contacts

import "errors"

var (
	ContactNotFoundErr = errors.New("contact not found")
	NameAlreadyUsedErr = errors.New("contact name already used")
)
