package contact_logs

import (
	"strings"
	"time"

	"gorm.io/gorm"

	"my/crm-golang/internal/models/contact_log"
	"my/crm-golang/internal/my_errors"
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

func (r *Repository) GetById(id int) (*contact_log.ContactLog, error) {
	var contactLogModel contact_log.ContactLog
	if err := r.db.Where("id = ?", id).First(&contactLogModel).Error; err != nil {
		return nil, err
	}
	return &contactLogModel, nil
}

func (r *Repository) Create(contact *contact_log.ContactLog) error {
	err := r.db.Create(contact).Error
	if err != nil && strings.Contains(err.Error(), "violates foreign key constraint") {
		return my_errors.ContactIdNotFoundErr
	}
	return err
}

func (r *Repository) DeleteById(id int) error {
	if err := r.db.Where("id = ?", id).Delete(&contact_log.ContactLog{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) Update(contactLogModel *contact_log.ContactLog, contactUpdateData *contact_log.ContactLogUpdateData) error {
	updateFields := r.getUpdateFields(contactUpdateData)
	if err := r.db.Model(contactLogModel).Updates(updateFields).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) getUpdateFields(contactUpdateData *contact_log.ContactLogUpdateData) map[string]interface{} {
	updateFields := map[string]interface{}{}

	if contactUpdateData.ContactId != 0 {
		updateFields["contact_id"] = contactUpdateData.ContactId
	}
	if contactUpdateData.LogMessage != "" {
		updateFields["log_message"] = contactUpdateData.LogMessage
	}
	if !time.Time.IsZero(contactUpdateData.Datetime) {
		updateFields["datetime"] = contactUpdateData.Datetime
	}

	return updateFields
}
