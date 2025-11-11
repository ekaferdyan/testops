package utils

import (
	"regexp"
	"strings"
)

// Regex ini mengizinkan huruf, titik, DAN SPASI.
var containsInvalidNameCharRegex = regexp.MustCompile("[^a-zA-Z. ]")

// "input HARUS terdiri dari satu atau lebih digit SAJA dari awal sampai akhir".
var onlyDigits = regexp.MustCompile(`^\d+$`) // atau `^[0-9]+$`

// Pola [0-9] berarti "cari satu digit dari 0 sampai 9".
var containsDigit = regexp.MustCompile("[0-9]")

// "input MENGANDUNG setidaknya satu huruf di MANA SAJA".
var containsLetters = regexp.MustCompile("[a-zA-Z]") //

// "input HARUS terdiri dari satu atau lebih huruf SAJA dari awal sampai akhir".
var onlyLetters = regexp.MustCompile(`^[a-zA-Z]+$`)

// Pola Regex untuk Nomor HP Indonesia yang menerima: 08..., +628..., 628..., dan 8...
// Ini mengizinkan '+' HANYA di awal, diikuti oleh pola yang valid, dan diakhiri ($).
var idPhoneRegex = regexp.MustCompile(`^(?:0|\+62|62)?8[1-9]\d{6,11}$`)

func ContainsSpecialCharacters(input string) bool {
	return containsInvalidNameCharRegex.MatchString(input)
}

func OnlyDigits(input string) bool {
	return onlyDigits.MatchString(input)
}

func ContainsDigits(input string) bool {
	return containsDigit.MatchString(input)
}

func ContainsLetters(input string) bool {
	return containsLetters.MatchString(input)
}

func OnlyLetters(input string) bool {
	return onlyLetters.MatchString(input)
}

func IdRegex(input string) bool {
	return idPhoneRegex.MatchString(input)
}

func NormalizePhone(input string) string {
	// 1. Bersihkan semua karakter non-digit kecuali tanda plus (+)
	re := regexp.MustCompile(`[^\d\+]`)
	normalized := re.ReplaceAllString(input, "")

	// 2. Jika dimulai dengan '08', ganti menjadi '+628'
	if strings.HasPrefix(input, "08") {
		return "+62" + strings.TrimPrefix(normalized, "0")
	}

	//2. Jika dimulai dengan '8', ganti menjadi '+628'
	if strings.HasPrefix(input, "8") {
		return "+62" + normalized
	}

	//3. Jika dimulai dengan '62', ganti menjadi '+628'
	if strings.HasPrefix(input, "62") {
		return "+" + normalized
	}

	return normalized // fallback

}
