package contact_logs

import (
	"errors"

	"my/crm-golang/internal/models/contact_log"
	"my/crm-golang/internal/storage/postgres/contacts"
)

type Service struct {
	repository Repository
}

func (s *Service) GetById(id int) (*contact_log.ContactLog, error) {
	return s.repository.GetById(id)
}

func (s *Service) Create(contactLog *contact_log.ContactLog) error {
	contactLog.SetTimestampToNow()

	return s.repository.Create(contactLog)
}

func (s *Service) Update(id int, contactUpdateData *contact_log.ContactLogUpdateData) error {
	contactLogModel, err := s.GetById(id)
	if errors.Is(err, contacts.ContactLogNotFoundErr) {
		return NotFoundErr
	}
	if err != nil {
		return err
	}

	return s.repository.Update(contactLogModel, contactUpdateData)
}

func (s *Service) DeleteById(id int) error {
	return s.repository.DeleteById(id)
}
