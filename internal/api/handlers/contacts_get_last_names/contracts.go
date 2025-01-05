package contacts_get_last_names

type Service interface {
	GetLastContactsNames(count uint) ([]string, error)
}
