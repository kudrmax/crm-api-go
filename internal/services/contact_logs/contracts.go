package contact_logs

import (
	"my/crm-golang/internal/models/contact_log"
)

type Repository interface {
	GetById(id int) (*contact_log.ContactLog, error)
	Create(contactLog *contact_log.ContactLog) error
	DeleteById(id int) error
	Update(contactLog *contact_log.ContactLog, contactLogUpdateData *contact_log.ContactLogUpdateData) error
}
