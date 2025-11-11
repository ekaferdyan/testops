package validator

import (
	"github.com/go-playground/validator/v10"
	// Import semua custom validation function yang Anda butuhkan
	"sambel-ulek/backend/utils"
)

// Validate adalah "Mesin" global kita yang dibuat 1x dan dipakai ulang
var Validate *validator.Validate

func init() {
	// 1. Inisiasi instance validator
	Validate = validator.New()

	// Validate Phone Number Indonesia
	Validate.RegisterValidation("id_phone_not_valid", func(fl validator.FieldLevel) bool {
		// 1. Ambil nilainya (sebagai string)
		//    fl.Field().Interface() mengambil nilai sebagai interface{}
		//    .(string) mengubah nilai itu ke string
		value, ok := fl.Field().Interface().(string)

		// 2. Cek apakah tipenya benar-benar string
		//    'ok' akan 'false' jika field-nya bukan string (misal: int, atau nil)
		if !ok {
			return false
		}

		// 3. Jika tipenya string, teruskan nilainya ke fungsi util Anda
		//    (termasuk jika string-nya kosong, regex akan return false)
		return utils.IdRegex(value)
	})

	//Validate Name Special Character
	Validate.RegisterValidation("name_special_character", func(fl validator.FieldLevel) bool {
		value, ok := fl.Field().Interface().(string)

		if !ok {
			return false
		}

		return !utils.ContainsSpecialCharacters(value)
	})

	//Validate Name Contains Digits
	Validate.RegisterValidation("name_contains_digits", func(fl validator.FieldLevel) bool {
		value, ok := fl.Field().Interface().(string)

		if !ok {
			return false
		}
		return !utils.ContainsDigits(value)
	})

}
