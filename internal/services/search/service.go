package search

import (
	"strings"

	fuzz "github.com/paul-mannino/go-fuzzywuzzy"
)

type SearchEngine struct{}

func (s *SearchEngine) Search(query string, stringsArray []string) []string {
	fuzzySearchArray := s.fuzzySearch(query, stringsArray)
	prefixSearchArray := s.prefixSearch(s.getPrefix(fuzzySearchArray), stringsArray)

	resultList := append(prefixSearchArray, fuzzySearchArray...)

	resultSet := map[string]bool{}
	for _, str := range resultList {
		resultSet[str] = true
	}

	resultList = make([]string, 0, len(resultSet))
	for str := range resultSet {
		resultList = append(resultList, str)
	}

	return resultList
}

func (s *SearchEngine) fuzzySearch(query string, stringsArray []string) []string {
	query = strings.ToLower(query)
	result := make([]string, 0, len(stringsArray))
	for _, str := range stringsArray {
		if fuzz.Ratio(query, strings.ToLower(str)) > 80 {
			result = append(result, str)
		}
	}

	return result
}

func (s *SearchEngine) prefixSearch(prefix string, stringsArray []string) []string {
	if prefix == "" {
		return []string{}
	}

	result := make([]string, 0, len(stringsArray))
	for _, str := range stringsArray {
		if strings.HasPrefix(strings.ToLower(str), strings.ToLower(prefix)) {
			result = append(result, str)
		}
	}

	return result
}

func (s *SearchEngine) getPrefix(fuzzySearchArray []string) string {
	if len(fuzzySearchArray) == 0 {
		return ""
	}

	prefix := fuzzySearchArray[0]
	for _, str := range fuzzySearchArray {
		if len(str) < len(prefix) {
			prefix = str
		}
	}

	return prefix
}
