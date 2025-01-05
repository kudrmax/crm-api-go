package contacts

import (
	"errors"

	"my/crm-golang/internal/models/contact"
	"my/crm-golang/internal/storage/postgres/contacts"
)

type Service struct {
	repository Repository
}

func (s *Service) GetByName(name string) (*contact.Contact, error) {
	return s.repository.GetByName(name)
}

func (s *Service) Create(contact *contact.Contact) error {
	contactModel, err := s.GetByName(contact.Name)
	if err != nil {
		return err
	}

	if contactModel != nil {
		return NameAlreadyUsedErr
	}

	return s.repository.Create(contact)
}

func (s *Service) Update(name string, contactUpdateData *contact.ContactUpdateData) error {
	contactModel, err := s.GetByName(name)
	if errors.Is(err, contacts.ContactNotFoundErr) {
		return NotFoundErr
	}
	if err != nil {
		return err
	}

	return s.repository.Update(contactModel, contactUpdateData)
}

func (s *Service) DeleteByName(name string) error {
	return s.repository.DeleteByName(name)
}
