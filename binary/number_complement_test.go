package binary

import "testing"

func TestNumberComplement(t *testing.T) {
	data := []struct {
		i, e int
	}{
		{5, 2},
		{1, 0},
		{0, 0},
	}
	for _, val := range data {
		o := NumberComplement(val.i)
		if o != val.e {
			t.Errorf("NumberComplement(%d) 计算错误: 期望 %d, 实际为 %d\n", val.i, val.e, o)
		}
	}
}
