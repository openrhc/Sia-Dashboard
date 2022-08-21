package utils

func IsDigit(str string) bool {
	for _, code := range str {
		if code < 48 || code > 57 {
			return false
		}
	}
	return true
}
