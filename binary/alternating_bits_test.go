package binary

import "testing"

func TestAlternatingBits(t *testing.T) {
	data := []struct {
		i int
		e bool
	}{
		{5, true},
		{7, false},
		{11, false},
		{10, true},
	}

	for _, val := range data {
		o := AlternatingBits(val.i)
		if o != val.e {
			t.Errorf("AlternatingBits(%d) 计算错误, 期望：%t, 实际上是: %t", val.i, val.e, o)
		}

		o2 := AlternatingBits2(val.i)
		if o2 != val.e {
			t.Errorf("AlternatingBits2(%d) 计算错误, 期望：%t, 实际上是: %t", val.i, val.e, o2)
		}
	}
}
