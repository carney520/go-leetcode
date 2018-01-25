package array

import "testing"

type twoSumTestData struct {
	target int
	array  []int
	expect [2]int
}

func TestTwoSum(t *testing.T) {
	data := []twoSumTestData{
		twoSumTestData{9, []int{1, 3, 4, 8, 3, 3}, [2]int{1, 4}},
		twoSumTestData{9, []int{1, 1, 4, 8, 2, 1}, [2]int{1, 4}},
		twoSumTestData{9, []int{1, 9, 4, 5, 2, 1}, [2]int{3, 4}},
		twoSumTestData{9, []int{1, 9, 4, 6, 2, 1}, [2]int{0, 0}},
	}

	for _, val := range data {
		res, _ := TwoSum(val.array, val.target)
		if !IntSliceEqual(res[:], val.expect[:]) {
			t.Errorf("TwoSum(%v, %d) 计算错误，期望： %v, 实际上是： %v\n", val.array, val.target, val.expect, res)
		}
	}
}
