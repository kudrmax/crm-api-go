package contact_logs_get_all_list

import "my/crm-golang/internal/models/contact_log"

type ContactLogService interface {
	GetById(id int) ([]*contact_log.ContactLog, error)
}

type ContactService interface {
	GetIdByName(name string) (int, error)
}
