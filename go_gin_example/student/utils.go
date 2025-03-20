package student

import (
	"time"

	"github.com/go-playground/validator/v10"
)

func ValidateDate(fl validator.FieldLevel) bool {
	const layout = "2006-01-02"
	dateStr := fl.Field().String()
	_, err := time.Parse(layout, dateStr)
	return err == nil
}
