package validator

import "github.com/go-playground/validator/v10"

// Validate adalah "Mesin" global kita yang dibuat 1x dan dipakai ulang
var Validate = validator.New()
