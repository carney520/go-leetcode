package array

// TwoSum 给定一个值a，求数组中两个值得和等于a的索引。
// 要求索引是有序的，索引从1开始
func TwoSum(arr []int, target int) (res [2]int, ok bool) {
	cache := map[int]int{}
	for i, value := range arr {
		if value >= target {
			continue
		}

		if j, exists := cache[target-value]; exists {
			res[0] = j + 1
			res[1] = i + 1
			return res, true
		}

		if _, exists := cache[value]; !exists {
			cache[value] = i
		}
	}
	return
}
