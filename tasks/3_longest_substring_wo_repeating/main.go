package main

func longestSubstr(s string) int {
	occur := make(map[rune]struct{}, 0)
	max := 0
	arr := []rune(s)

	if len(arr) == 1 {
		return 1
	}

	for i := range arr {
		for j := i; j < len(arr); j++ {
			if _, exist := occur[arr[j]]; !exist {
				occur[arr[j]] = struct{}{}
			} else {
				length := len(occur)
				if length > max {
					max = length
				}
				occur = make(map[rune]struct{}, 0)
				break
			}
		}
	}

	return max
}
