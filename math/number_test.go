package math

import (
	"testing"
)

func TestLcm(t *testing.T) {
	testCases := []struct {
		name string
		a    int
		b    int
		exp  int
	}{
		{
			name: "Pos_Pos",
			a:    10,
			b:    6,
			exp:  30,
		},
		{
			name: "Zero_Pos",
			a:    0,
			b:    100,
			exp:  0,
		},
		{
			name: "Zero_Zero",
			a:    0,
			b:    0,
			exp:  0,
		},
		{
			name: "Neg_Neg",
			a:    -100,
			b:    -7,
			exp:  700,
		},
		{
			name: "Neg_Pos",
			a:    -7,
			b:    5,
			exp:  35,
		},
		{
			name: "Zero_neg",
			a:    0,
			b:    -129,
			exp:  0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out := Lcm(tc.a, tc.b)
			if out != tc.exp {
				t.Errorf("Expected %d, got %d instead\n", tc.exp, out)
			}
		})
	}
}
