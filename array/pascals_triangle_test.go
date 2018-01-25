package array

import (
	"testing"
)

func TestPascalsTriangle(t *testing.T) {
	res := PascalsTriangle(5)
	expected := [][]int{
		[]int{1},
		[]int{1, 1},
		[]int{1, 2, 1},
		[]int{1, 3, 3, 1},
		[]int{1, 4, 6, 4, 1},
	}

	if len(res) != 5 {
		t.Errorf("PascalTriangle返回的行数有误, 期望5, 实际上是:%d\n", len(res))
	}

	if !NestSliceEqual(res, expected) {
		t.Errorf("PascalsTriangle(5) 期望得到的是 %v, 但是得到 %v\n", expected, res)
	}
}

func NestSliceEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if !IntSliceEqual(a[i], b[i]) {
			return false
		}
	}
	return true
}
