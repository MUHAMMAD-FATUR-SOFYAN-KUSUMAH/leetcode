package leetcode

func reverseWords(s string) string {
	var arr []string
	var temp string = ""
	for i := 0; i < len(s); i++ {
		if s[i] != 32 {
			temp = temp + string(s[i])
			if i+1 > len(s)-1 {
				arr = append(arr, temp)
				temp = ""
				break
			}
			if s[i+1] == 32 {
				arr = append(arr, temp)
				temp = ""
			}
		}
	}

	for i := len(arr) - 1; i >= 0; i-- {
		temp = temp + arr[i]
		if i == 0 {
			break
		}
		temp = temp + " "
	}
	return temp
}
