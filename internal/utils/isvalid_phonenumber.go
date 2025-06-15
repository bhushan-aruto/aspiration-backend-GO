package utils

import (
	"strings"

	"github.com/nyaruka/phonenumbers"
)

func IsValidPhonenumber(phone string) bool {
	num, err := phonenumbers.Parse(phone, "")
	if err != nil {
		return false
	}

	if !phonenumbers.IsValidNumber(num) || !strings.HasPrefix(phone, "+") {
		return false
	}

	return phonenumbers.IsValidNumber(num)

}
