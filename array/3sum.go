package array

import (
	"sort"
)

// ThreeSum 给定一个有序数组arr， 找出所有a + b + c = 0的组合
func ThreeSum(arr []int) (res [][]int) {
	sort.Sort(sort.IntSlice(arr))
	for i := 0; i < len(arr)-2; i++ {
		// 跳过连续的数值
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}

		j := i + 1
		k := len(arr) - 1
		target := -arr[i]
		for j < k {
			if arr[j]+arr[k] == target {
				res = append(res, []int{arr[i], arr[j], arr[k]})
				j++
				k--
				// 跳过连续的数值
				for j < k && arr[j] == arr[j+1] {
					j++
				}

				for j < k && arr[k] == arr[k-1] {
					k--
				}
			} else if arr[j]+arr[k] > target {
				// 递减
				k--
			} else {
				// 递增
				j++
			}
		}
	}
	return
}
