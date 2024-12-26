package contacts

import "my/crm-golang/internal/models/contact"

type Repository struct {
	//log internal.Log
	//db  *sqlx.DB
}

//	func New(log internal.Log, db *sqlx.DB) *Repository {
//		return &Repository{
//			log: log,
//			db:  db,
//		}
//	}

func (r *Repository) GetByName(name string) (*contact.Contact, error) {
	return nil, nil
}

func (r *Repository) Create(contact *contact.Contact) error {
	return nil
}

func (r *Repository) DeleteByName(name string) error {
	return nil
}

func (r *Repository) Update(contact *contact.Contact, contactUpdateData *contact.ContactUpdateData) error {
	return nil
}
