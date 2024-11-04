package validation

import (
	"net/mail"
	"regexp"
)

func IsValidBDMobileNumber(number string) bool {
	rex := regexp.MustCompile(`^(\+88|(0{2})?88)?(01)[3-9]{1}[0-9]{8}$`)
	return rex.MatchString(number)
}

func IsValidEmailAddress(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
