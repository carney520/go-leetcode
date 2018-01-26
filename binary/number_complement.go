package binary

// NumberComplement 求一个二进制非0开头的整型的补码， 返回补码后的数字
// 比如 5 二进制形式为101，他的补码为010，所以他的返回值为2
func NumberComplement(num int) int {
	mask := ^0 // 111111
	for mask&num != 0 {
		mask = mask << 1 // 111000
	}
	return ^mask & ^num // 000101
}
