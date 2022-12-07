package utils

import "testing"

func TestMin(t *testing.T) {
	testCases := []struct {
		name string
		args []int
		exp  int
	}{
		{
			name: "first",
			args: []int{1, 2, 3, 4, 5},
			exp:  1,
		},
		{
			name: "last",
			args: []int{1, 2, 3, 4, -5},
			exp:  -5,
		},
		{
			name: "middle",
			args: []int{-10, 1, -100, 100, 10000},
			exp:  -100,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out := IntMin(tc.args...)
			if out != tc.exp {
				t.Errorf("Expected %d, got %d instead\n", tc.exp, out)
			}
		})
	}
}
