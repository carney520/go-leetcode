package array

// 这个文件主要实现了从数组中删除重复的元素

// RmDupElms 从"有序"的数组中删除重复的元素, 要求不分配新的数组
func RmDupElms(arr []interface{}) []interface{} {
	if len(arr) < 2 {
		return arr
	}

	var i, j int
	for i, j = 1, 0; i < len(arr); i++ {
		if arr[j] != arr[i] {
			j++
			arr[j] = arr[i]
		}
	}
	return arr[:j+1]
}

// RmLimitDupElms 同RmDupElms, 只不过可以限定重复的次数
func RmLimitDupElms(arr []interface{}, limit int) []interface{} {
	if len(arr) < 2 {
		return arr
	}

	var i, j, num int
	for i, j = 1, 0; i < len(arr); i++ {
		if arr[j] != arr[i] {
			num = 0
			j++
			arr[j] = arr[i]
		} else {
			num++
			if num < limit {
				j++
				arr[j] = arr[i]
			}
		}
	}
	return arr[:j+1]
}
