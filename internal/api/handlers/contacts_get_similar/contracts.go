package contacts_get_similar

type Service interface {
	GetSimilarNames(name string) ([]string, error)
}
