package utils

import (
	"github.com/HideN2000/go_ds/math"
)

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
