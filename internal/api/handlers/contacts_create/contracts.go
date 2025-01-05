package contacts_create

import "my/crm-golang/internal/models/contact"

type Service interface {
	Create(contact *contact.Contact) error
}
