package contacts_update

import "my/crm-golang/internal/models/contact"

type Service interface {
	Update(name string, contactUpdateData *contact.ContactUpdateData) error
}
