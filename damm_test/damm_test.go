package damm_test

import (
	"testing"

	"github.com/umahmood/damm"
)

var tests = []struct {
	in      int
	out     int
	isValid bool
}{
	{7, 76, true},
	{42, 427, true},
	{572, 5724, true},
	{7823, 78234, true},
	{82915, 829150, true},
	{917342, 9173427, true},
}

func TestDamm(t *testing.T) {
	for _, v := range tests {
		out := damm.Calc(v.in)
		if out != v.out {
			t.Errorf("fail calc: test %d got %d want %d", v.in, out, v.out)
		}
		valid := damm.Validate(out)
		if valid != v.isValid {
			t.Errorf("fail validate: test %d got %t want %t", v.out, valid, v.isValid)
		}
	}
}
