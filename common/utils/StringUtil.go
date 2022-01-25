package utils

func IsEmpty(value string) bool {

	if value == "" || len(value) < 1{
		return true
	}
	return false
}
