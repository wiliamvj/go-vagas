package utils

import "strings"

var (
	terms = []string{"#govagas", "#golangvagas", "#vagasgolang", "#vagasgo", "#gojobs"}
)

func FilterTerms(text string) bool {
	for _, term := range terms {
		if strings.Contains(strings.ToLower(text), strings.ToLower(term)) {
			return true
		}
	}
	return false
}
