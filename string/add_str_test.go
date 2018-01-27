package strings

import (
	"testing"
)

func TestAddStr(t *testing.T) {
	data := []struct {
		a, b, e string
	}{
		{"0", "2", "2"},
		{"7", "8", "15"},
		{"70", "8", "78"},
		{"99", "5", "104"},
		{"34", "102", "136"},
		{"34000000", "102", "34000102"},
	}

	for _, val := range data {
		o := AddStr(val.a, val.b)
		if o != val.e {
			t.Errorf("AddStr(%s, %s)计算错误, 期望%s, 实际上是: %s", val.a, val.b, val.e, o)
		}
	}
}
