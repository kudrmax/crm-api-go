package contacts

import (
	"errors"
	"strings"

	"gorm.io/gorm"

	"my/crm-golang/internal/models/contact"
	errors2 "my/crm-golang/internal/my_errors"
)

type Repository struct {
	//log internal.Log
	db *gorm.DB
}

func New(
	//log internal.Log,
	db *gorm.DB,
) *Repository {
	return &Repository{
		db: db,
		//log: log,
	}
}

func (r *Repository) GetByName(name string) (*contact.Contact, error) {
	var contactModel contact.Contact
	err := r.db.Where("name = ?", name).First(&contactModel).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors2.ContactNotFoundErr
	}

	return &contactModel, nil
}

func (r *Repository) GetAll() ([]*contact.Contact, error) {
	contactModels := make([]*contact.Contact, 0)
	err := r.db.Find(&contactModels).Error
	if err != nil {
		return nil, err
	}
	return contactModels, nil
}

func (r *Repository) Create(contact *contact.Contact) error {
	err := r.db.Create(contact).Error
	if err != nil && strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
		return errors2.NameAlreadyUsedErr
	}
	return err
}

func (r *Repository) DeleteByName(name string) error {
	if err := r.db.Where("name = ?", name).Delete(&contact.Contact{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) Update(contactModel *contact.Contact, contactUpdateData *contact.ContactUpdateData) error {
	updateFields := r.getUpdateFields(contactUpdateData)

	err := r.db.Model(contactModel).Updates(updateFields).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors2.ContactNotFoundErr
	}
	return err
}

func (r *Repository) GetLastContacts(count uint) ([]*contact.Contact, error) {
	models, err := r.GetAll()
	if err != nil {
		return nil, err
	}
	return models[:count], nil
}

func (r *Repository) getUpdateFields(contactUpdateData *contact.ContactUpdateData) map[string]interface{} {
	updateFields := map[string]interface{}{}

	if contactUpdateData.Name != "" {
		updateFields["name"] = contactUpdateData.Name
	}
	if contactUpdateData.Phone != "" {
		updateFields["phone"] = contactUpdateData.Phone
	}
	if contactUpdateData.Telegram != "" {
		updateFields["telegram"] = contactUpdateData.Telegram
	}
	if contactUpdateData.Birthday != "" {
		updateFields["birthday"] = contactUpdateData.Birthday
	}

	return updateFields
}
