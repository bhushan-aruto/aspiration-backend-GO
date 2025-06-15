package utils

import "regexp"

func IsPhoneNumberValid(phone string) bool {
	re := regexp.MustCompile(`^\d{10}$`)
	return re.MatchString(phone)
}
