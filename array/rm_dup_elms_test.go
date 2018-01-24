package array

import "testing"

// 测试从数组中删除重复元素

func TestRmDupElms(t *testing.T) {
	arr := [...]interface{}{1, 2, 2, 3, 4}
	res := RmDupElms(arr[:])
	expected := []interface{}{1, 2, 3, 4}
	if len(res) != 4 {
		t.Error("RmDupElms后长度应该为4, 实际上是:", len(res))
	}

	if !SliceEqual(res, expected) {
		t.Errorf("RmDupElms后不匹配, 期望%v, 实际上是 %v", expected, res)
	}
}

func TestRmLimitDupElms(t *testing.T) {
	arr := [...]interface{}{1, 2, 2, 2, 2, 3, 4}
	res := RmLimitDupElms(arr[:], 2)
	expected := []interface{}{1, 2, 2, 3, 4}
	if len(res) != 5 {
		t.Error("RmDupElms后长度应该为5, 实际上是:", len(res))
	}

	if !SliceEqual(res, expected) {
		t.Errorf("RmLimitDupElms 不匹配, 期望%v, 实际上是 %v", expected, res)
	}
}
