package contacts_get_one

import "my/crm-golang/internal/models/contact"

type Service interface {
	GetByName(name string) (*contact.Contact, error)
}
