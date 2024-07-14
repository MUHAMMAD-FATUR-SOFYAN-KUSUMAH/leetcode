package leetcode

func recursive_search(str string, length int, word string, length_word int) any {
	if length >= len(str) {
		if length_word >= len(word) {
			return true
		}
		return false
	}

	if str[length] == word[length_word] {
		if length_word >= len(word)-1 {
			return true
		}
		return recursive_search(str, length+1, word, length_word+1)
	}

	return recursive_search(str, length+1, word, length_word)
}

func exist(board [][]byte, word string) bool {
	temp := ""

	for _, slice := range board {
		for _, value := range slice {
			temp = temp + string(value)
		}
	}

	r := recursive_search(temp, 0, word, 0)

	bo := r.(bool)

	return bo
}
