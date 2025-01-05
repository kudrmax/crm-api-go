package contacts

import (
	"errors"
)

var (
	NotFoundErr        = errors.New("contact not found")
	NameAlreadyUsedErr = errors.New("contact name already used")
)
