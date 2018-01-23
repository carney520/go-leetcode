package rmelm

// RemoveElement 从数组/切片中原地删除一个元素
// 在go中最常使用的是slice, 所以这里使用slice作为示例. 原理删除元素, 不影响slide的
// 底层数组
func RemoveElement(arr []interface{}, elm interface{}) []interface{} {
	var j = 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == elm {
			continue
		}
		arr[j] = arr[i]
		j++
	}
	return arr[:j]
}
