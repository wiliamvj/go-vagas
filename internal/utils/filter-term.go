package utils

import "strings"

var (
	terms = []string{"#1123", "#1124"}
)

func FilterTerms(text string) bool {
	for _, term := range terms {
		if strings.Contains(strings.ToLower(text), strings.ToLower(term)) {
			return true
		}
	}
	return false
}
