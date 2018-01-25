package array

import (
	"fmt"
	"testing"
)

func TestPlusOne(t *testing.T) {
	data := [][2][]int{}
	data = append(
		data,
		[2][]int{[]int{1, 2, 3, 4}, []int{1, 2, 3, 5}},
		[2][]int{[]int{1, 2, 3, 9}, []int{1, 2, 4, 0}},
	)
	for _, item := range data {
		val, exp := item[0], item[1]
		// 拷贝
		org := make([]int, len(val))
		copy(org, val)
		fmt.Println(len(org))

		res := PlusOne(val)
		if !IntSliceEqual(res, exp) {
			t.Errorf("PlusOne异常, %v的求值结果应该为 %v, 但是得到%v\n", org, exp, res)
		}
	}
}

func IntSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
