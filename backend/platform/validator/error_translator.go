package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// TranslateError menerjemahkan error teknis (yang dikembalikan oleh package validator)
// menjadi []ErrorResponse yang rapi, berisi Field dan Message yang user-friendly.
func TranslateError(err error) []ErrorResponse {
	// Inisialisasi slice kosong untuk menampung semua pesan error.
	var errors []ErrorResponse

	// 1. Cek Tipe Error: Memastikan error yang masuk adalah benar-benar
	//    error validasi (validator.ValidationErrors).
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		// 2. Iterasi Error: Melakukan loop untuk setiap field yang gagal validasi.
		for _, fe := range validationErrors {
			// 3. Append Error: Menambahkan objek ErrorResponse baru ke slice 'errors'.
			errors = append(errors, ErrorResponse{
				// Mengubah nama field menjadi huruf kecil (misal: "Email" menjadi "Email")
				// untuk konsistensi dalam response JSON.
				Field: strings.ToLower(fe.Field()),
				// Memanggil fungsi kamus penerjemah untuk mendapatkan pesan Bahasa Indonesia.
				Message: getErrorMessage(fe),
			})
		}
	}

	// Mengembalikan slice berisi daftar semua error yang sudah diterjemahkan.
	return errors
}

// getErrorMessage adalah "kamus" penerjemah kita
func getErrorMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "Field tidak boleh kosong"
	case "min":
		return fmt.Sprintf("Minimal harus %s karakter", fe.Param())
	case "max":
		return fmt.Sprintf("Maksimal %s karakter", fe.Param())
	case "email":
		return "Format tidak valid. Pastikan formatnya benar."
	case "id_phone_not_valid":
		return "Format tidak valid. Gunakan format yang sesuai, contoh: 08123456789 atau +628123456789."
	case "name_special_character":
		return "Nama tidak boleh mengandung spesial karakter"
	case "name_contains_digits":
		return "Nama tidak boleh mengandung angka"
	default:
		return "Input tidak valid"
	}
}
