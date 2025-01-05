package contacts_search

type Service interface {
	//GetAll() ([]*contact.Contact, error)
	GetSimilarNames(name string) ([]string, error)
}
