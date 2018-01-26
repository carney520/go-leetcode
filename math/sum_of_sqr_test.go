package math

import "testing"

func TestSumOfSqr(t *testing.T) {
	data := []struct {
		i int
		e bool
	}{
		{5, true},
		{6, false},
		{3, false},
		{9, true},
	}

	for _, val := range data {
		o := SumOfSqr(val.i)
		if o != val.e {
			t.Errorf("SumOfSqr(%d) 计算失败，期望: %t, 实际上为: %t", val.i, val.e, o)
		}
	}
}
