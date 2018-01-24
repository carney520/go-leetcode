package array

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestRemoveElement(t *testing.T) {
	arr := [...]interface{}{1, 2, 3, 4, 2}
	res := RemoveElement(arr[:], 2)
	expected := []interface{}{1, 3, 4}

	// 长度
	if len(res) != 3 {
		t.Error("RemoveElement 长度应该为3, 当时得到", len(res))
	}

	// 比较元素
	if !SliceEqual(res, expected) {
		t.Error("RemoveElement 没有删除输出预定元素", expected, res)
	}

	// 获取slice底层数组
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&res))
	arrPtr := uintptr(unsafe.Pointer(&arr))
	if sliceHeader.Data != arrPtr {
		t.Errorf("底层数组不一致, expected: %x, actual: %x\n", arrPtr, sliceHeader)
	}
}
