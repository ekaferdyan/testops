package utils

import "regexp"

// Pola Regex akan di-compile sekali saat program dimulai.
var containsInvalidCharRegex = regexp.MustCompile("[^a-zA-Z0-9_.]")

func ContainsSpecialCharacters(username string) bool {
	return containsInvalidCharRegex.MatchString(username)
}
