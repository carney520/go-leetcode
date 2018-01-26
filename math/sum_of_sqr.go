package math

import (
	m "math"
)

// SumOfSqr 给定一个正整数c，判断是否有两个数a，b，使得a^2 + b^2 = c
func SumOfSqr(num int) bool {
	if num < 0 {
		return false
	}
	i, j := 0, int(m.Sqrt(float64(num)))
	for i <= j {
		r := i*i + j*j
		if r == num {
			return true
		} else if r < num {
			i++
		} else {
			j++
		}
	}
	return false
}
