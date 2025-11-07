package utils

import "regexp"

// Pola Regex akan di-compile sekali saat program dimulai.
var containsInvalidCharRegex = regexp.MustCompile("[^a-zA-Z0-9_.]")

func ContainsSpecialCharacters(Email string) bool {
	return containsInvalidCharRegex.MatchString(Email)
}
