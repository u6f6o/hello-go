package hello_go

func AreAnagrams(one, another string) bool {
	if len([]rune(one)) < 1 || len([]rune(one)) != len([]rune(another)) {
		return false
	}
	runeCount := make(map[rune]int)

	for _, runeVal := range one {
		if prevVal, present := runeCount[runeVal]; present {
			runeCount[runeVal] = prevVal + 1
		} else {
			runeCount[runeVal] = 1
		}
	}
	for _, runeVal := range another {
		if prevVal, present := runeCount[runeVal]; present {
			if prevVal > 0 {
				runeCount[runeVal] = prevVal - 1
			} else {
				return false
			}
		} else {
			return false
		}
	}
	return true
}