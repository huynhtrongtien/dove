package http_parser

import (
	"time"

	"github.com/beego/beego/validation"
)

func SetupValidation() {
	validation.AddCustomFunc("DateFormat", DateFormat)
	validation.AddCustomFunc("HourFormat", HourFormat)
}

func DateFormat(v *validation.Validation, obj interface{}, key string) {
	dateValue, ok := obj.(string)
	if !ok {
		_ = v.SetError(key, "Can't parse to string")
	}

	_, err := time.Parse("2006-01-02", dateValue)
	if err != nil {
		_ = v.SetError(key, "Must be date value with format YYYY-MM-DD")
	}
}

func HourFormat(v *validation.Validation, obj interface{}, key string) {
	dateValue, ok := obj.(string)
	if !ok {
		_ = v.SetError(key, "Can't parse to string")
	}

	_, err := time.Parse("15:06", dateValue)
	if err != nil {
		_ = v.SetError(key, "Must be date value with format HH:mm")
	}
}
