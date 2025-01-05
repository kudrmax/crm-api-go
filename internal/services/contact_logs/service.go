package contact_logs

import (
	"my/crm-golang/internal/models/contact_log"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}
func (s *Service) GetById(id int) ([]*contact_log.ContactLog, error) {
	return s.repository.GetById(id)
}

func (s *Service) Create(contactLog *contact_log.ContactLog) error {
	contactLog.SetTimestampToNow()
	return s.repository.Create(contactLog)
}

//func (s *Service) Update(id int, contactUpdateData *contact_log.ContactLogUpdateData) error {
//	contactLogModel, err := s.GetById(id)
//	if errors.Is(err, my_errors.ContactLogNotFoundErr) {
//		return my_errors.ContactLogNotFoundErr
//	}
//	if err != nil {
//		return err
//	}
//
//	return s.repository.Update(contactLogModel, contactUpdateData)
//}

func (s *Service) DeleteById(id int) error {
	return s.repository.DeleteById(id)
}
