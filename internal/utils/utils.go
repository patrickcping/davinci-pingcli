package utils

func RemoveTrailingS(word string) string {
	if len(word) > 0 && word[len(word)-1] == 's' {
		return word[:len(word)-1]
	}
	return word
}
