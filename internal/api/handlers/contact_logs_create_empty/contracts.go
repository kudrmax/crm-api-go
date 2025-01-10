package contact_logs_create_empty

import "my/crm-golang/internal/models/contact_log"

type ContactLogService interface {
	Create(contact *contact_log.ContactLog) error
}

type ContactService interface {
	GetIdByName(name string) (int, error)
}
