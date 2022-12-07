package utils

func IntMax(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	var max int = nums[0]
	for _, e := range nums {
		if max < e {
			max = e
		}
	}
	return max
}

func IntMin(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	var min int = nums[0]
	for _, e := range nums {
		if min > e {
			min = e
		}
	}
	return min
}
