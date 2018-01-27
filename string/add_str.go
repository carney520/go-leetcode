package strings

const zero byte = '0'

// AddStr 将a, b两个非负的整型字符串, 返回相加的结果
func AddStr(a, b string) string {
	j, k := len(a), len(b)
	var carry byte
	var res []byte
	for i := max(j, k); i > 0; i-- {
		var v byte
		if j > 0 {
			v += a[j-1] - zero
		}
		if k > 0 {
			v += b[k-1] - zero
		}
		v += carry
		carry = v / 10
		res = append(res, v%10)
		j--
		k--
	}
	if carry != 0 {
		res = append(res, carry)
	}

	rev := make([]byte, 0, len(res))
	for i := len(res); i > 0; i-- {
		rev = append(rev, res[i-1]+zero)
	}

	return string(rev)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
