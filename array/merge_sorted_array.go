package array

// MergeSortedArray 合并两个有序的数组a， b， 假设a有足够的空间容纳b，无须分配新的数组
// m 为a的有效数据存储长度
func MergeSortedArray(a, b []int, m int) []int {
	total := m + len(b)
	j, k := m-1, len(b)-1
	for i := total - 1; i >= 0; i-- {
		if j == -1 {
			copy(a, b[:i+1])
			break
		} else if k == -1 {
			break
		}

		if a[j] > b[k] {
			a[i] = a[j]
			j--
		} else {
			a[i] = b[k]
			k--
		}
	}
	return a[:total]
}
