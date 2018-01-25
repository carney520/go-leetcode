package array

import "testing"

type ThreeSumTestData struct {
	input  []int
	expect [][]int
}

func TestThreeSum(t *testing.T) {
	data := []ThreeSumTestData{
		ThreeSumTestData{[]int{-1, 0, 1, 2, -1, -4}, [][]int{
			[]int{-1, -1, 2},
			[]int{-1, 0, 1},
		}},
	}

	for _, v := range data {
		act := ThreeSum(CopyIntSlice(v.input))
		if !NestSliceEqual(act, v.expect) {
			t.Errorf("ThreeSum(%v)计算有误， 期望： %v, 实际上是： %v\n", v.input, v.expect, act)
		}
	}
}

func CopyIntSlice(slice []int) []int {
	r := make([]int, len(slice))
	copy(r, slice)
	return r
}
