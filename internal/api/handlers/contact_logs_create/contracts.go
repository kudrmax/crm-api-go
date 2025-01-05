package contact_logs_create

import "my/crm-golang/internal/models/contact_log"

type Service interface {
	Create(contact *contact_log.ContactLog) error
}
