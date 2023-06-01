package util

func InArray(str string, array []string) bool {
	for _, element := range array {
		if str == element {
			return true
		}
	}

	return false
}

func RemoveItem(str string, array []string) []string {
	result := []string{}
	for _, element := range array {
		if element != str {
			result = append(result, element)
		}
	}

	return result
}
