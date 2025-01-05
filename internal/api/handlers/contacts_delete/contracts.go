package contacts_delete

type Service interface {
	DeleteByName(name string) error
}
