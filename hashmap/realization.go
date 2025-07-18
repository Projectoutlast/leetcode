package hashmap

func canConstruct(ransomNote string, magazine string) bool {
	setArrange := make(map[rune]int, len(magazine))

	for _, letter := range magazine {
		if _, ok := setArrange[letter]; ok {
			setArrange[letter]++
			continue
		} else {
			setArrange[letter] = 1
		}
	}

	for _, letter := range ransomNote {
		if _, ok := setArrange[letter]; ok {
			setArrange[letter]--
			if setArrange[letter] == 0 {
				delete(setArrange, letter)
			}
		} else {
			return false
		}
	}

	return true
}
