package contacts

import "my/crm-golang/internal/models/contact"

type Repository interface {
	GetByName(name string) (*contact.Contact, error)
	GetAll() ([]*contact.Contact, error)
	GetLastContacts(count uint) ([]*contact.Contact, error)
	Create(contact *contact.Contact) error
	DeleteByName(name string) error
	Update(contact *contact.Contact, contactUpdateData *contact.ContactUpdateData) error
}
