package utils

import (
	math "project/hiden2000/acgo/lib/math"
)

// IntMaxは math.Intsのスライスnumsにおける最大値を返す
// Time: O(N)
func IntMax[T math.Ints](nums ...T) T {
	var max T
	if len(nums) == 0 {
		return max
	}
	max = nums[0]
	for _, e := range nums {
		if max < e {
			max = e
		}
	}
	return max
}

// IntMinは math.Intsのスライスnumsにおける最小値を返す
// Time: O(N)
func IntMin[T math.Ints](nums ...T) T {
	var min T
	if len(nums) == 0 {
		return min
	}
	min = nums[0]
	for _, e := range nums {
		if min > e {
			min = e
		}
	}
	return min
}
