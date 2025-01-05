package contacts_get_all

import "my/crm-golang/internal/models/contact"

type Service interface {
	GetAll() ([]*contact.Contact, error)
}
