package yago

// StringListContains returns true if the list containts the passed value,
// false otherwise
func StringListContains(list []string, value string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}
