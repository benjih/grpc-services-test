package model

// Assuming we would want to validate this down the line
func IsValidEmailAddress(emailAddress string) bool {
	return len(emailAddress) > 0
}

// Assuming we would want to validate this down the line
func IsValidTelephoneNumber(telephoneNumber string) bool {
	return len(telephoneNumber) > 0
}
