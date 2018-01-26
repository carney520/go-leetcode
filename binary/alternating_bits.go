package binary

// AlternatingBits 判断一个正整数的二进制位是否交替
func AlternatingBits(num int) bool {
	mask := ^0 << 1
	i := 1 << 1
	last := 1 & num

	for mask&num != 0 {
		j := i & num
		if j != 0 {
			if last != 0 {
				return false
			}
			last = 1
		} else {
			if last != 1 {
				return false
			}
			last = 0
		}

		mask <<= 1
		i <<= 1
	}
	return true
}

// AlternatingBits2 实现2
// 错开进行异或比对
func AlternatingBits2(num int) bool {
	// 000101010
	// >> 2
	// 000001010
	// ^
	// 000100000  如果是交替的，结果一定是2的幂
	// -1
	// 000011111 如果是2的幂，减1，后面的位都是1
	// &
	// 000000000 & 之后应该等于0
	num ^= num >> 2
	return num&(num-1) == 0
}
