package array

// PlusOne 实现十进制进位, 在末尾元素+1
func PlusOne(arr []int) []int {
	one := 1
	for i := len(arr) - 1; i >= 0; i-- {
		value := arr[i] + one
		one = value / 10
		arr[i] = value % 10
	}
	if one > 0 {
		newArr := make([]int, len(arr)+1)
		newArr = append(newArr, one)
		newArr = append(newArr, arr...)
		return newArr
	}
	return arr
}
