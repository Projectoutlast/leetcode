package arrays

func findMaxConsecutiveOnes(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	previousNum := nums[0]

	if len(nums) == 1 {
		return previousNum
	}

    var maxCons, currentCons int = 1, 1

	for _, num := range nums[2:] {
		if num == previousNum {
			currentCons++
		} else {
			if currentCons > maxCons {
				maxCons = currentCons
				currentCons = 1
			}
		}
	}

	return maxCons
}