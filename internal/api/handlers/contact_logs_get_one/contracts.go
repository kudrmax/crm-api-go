package contact_logs_get_one

import "my/crm-golang/internal/models/contact_log"

type ContactLogService interface {
	GetByLogId(id int) (*contact_log.ContactLog, error)
}

type ContactService interface {
	GetIdByName(name string) (int, error)
}
