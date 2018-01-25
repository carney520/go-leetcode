package array

import "testing"

func TestMergeSortedArray(t *testing.T) {
	a := make([]int, 8)
	copy(a, []int{1, 2, 3, 4})
	b := []int{2, 4}
	res := MergeSortedArray(a, b, 4)
	expect := []int{1, 2, 2, 3, 4, 4}

	if !IntSliceEqual(res, expect) {
		t.Errorf("MergeSortedArray 结果错误，期望: %v, 实际上为: %v ", expect, res)
	}

	a = make([]int, 8)
	copy(a, []int{3, 4})
	b = []int{1, 3, 7, 9}
	res = MergeSortedArray(a, b, 2)
	expect = []int{1, 3, 3, 4, 7, 9}

	if !IntSliceEqual(res, expect) {
		t.Errorf("MergeSortedArray 结果错误，期望: %v, 实际上为: %v ", expect, res)
	}
}
