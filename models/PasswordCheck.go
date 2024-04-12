package models

func doesContainDigits(password string) bool {
	for _, char := range password {
		if char >= 48 && char <= 57 {
			return true
		}
	}
	return false
}

func doesHaveCapsAndMinis(password string) bool {
	caps := 0
	minis := 0
	for _, char := range password {
		if char >= 97 && char <= 122 {
			caps = 1
		}
	}
	for _, char := range password {
		if char >= 65 && char <= 90 {
			minis = 1
		}
	}
	return caps == 1 && minis == 1
}

func doesNotContainsForbidden(password string) bool {
	forbidden := "'\"`"
	for _, char := range password {
		for _, char2 := range forbidden {
			if char == char2 {
				return false
			}
		}
	}
	return true
}

func doesContainsSpecialCharacter(password string) bool {
	for _, char := range password {
		if char > 32 && char < 48 && char > 57 && char < 65 && char > 90 && char < 97 && char > 122 {
			return true
		}
	}
	return false
}

func doesNotContainUnsupported(password string) bool {
	for _, char := range password {
		if char <= 32 {
			return false
		}
	}
	return true
}

func PasswordChecker(password string) bool {
	return doesContainDigits(password) && doesContainsSpecialCharacter(password) && doesNotContainsForbidden(
		password) && doesNotContainUnsupported(password) && doesHaveCapsAndMinis(password)
}
