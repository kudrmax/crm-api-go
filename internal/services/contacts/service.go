package contacts

import (
	"errors"

	"my/crm-golang/internal/models/contact"
	"my/crm-golang/internal/my_errors"
	"my/crm-golang/internal/services/search"
)

type Service struct {
	repository   Repository
	searchEngine search.SearchEngine
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetByName(name string) (*contact.Contact, error) {
	return s.repository.GetByName(name)
}

func (s *Service) GetAll() ([]*contact.Contact, error) {
	return s.repository.GetAll()
}

func (s *Service) Create(contact *contact.Contact) error {
	if contact.Name == "" {
		return my_errors.NameAlreadyUsedErr
	}
	return s.repository.Create(contact)
}

func (s *Service) Update(name string, contactUpdateData *contact.ContactUpdateData) error {
	contactModel, err := s.GetByName(name)
	if errors.Is(err, my_errors.ContactNotFoundErr) {
		return my_errors.ContactNotFoundErr
	}
	if err != nil {
		return err
	}

	return s.repository.Update(contactModel, contactUpdateData)
}

func (s *Service) DeleteByName(name string) error {
	return s.repository.DeleteByName(name)
}

func (s *Service) GetSimilarNames(name string) ([]string, error) {
	contactModels, err := s.GetAll()
	if err != nil {
		return []string{}, err
	}

	nameArray := make([]string, 0, len(contactModels))
	for _, contactModel := range contactModels {
		nameArray = append(nameArray, contactModel.Name)
	}

	return s.searchEngine.Search(name, nameArray), nil
}
