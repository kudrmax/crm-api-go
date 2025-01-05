package contacts

import (
	"gorm.io/gorm"

	"my/crm-golang/internal/models/contact"
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
	if err := r.db.Where("name = ?", name).First(&contactModel).Error; err != nil {
		return nil, err
	}
	return &contactModel, nil
}

func (r *Repository) Create(contact *contact.Contact) error {
	if err := r.db.Create(contact).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteByName(name string) error {
	if err := r.db.Where("name = ?", name).Delete(&contact.Contact{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) Update(contactModel *contact.Contact, contactUpdateData *contact.ContactUpdateData) error {
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

	if err := r.db.Model(contactModel).Updates(updateFields).Error; err != nil {
		return err
	}
	return nil
}
