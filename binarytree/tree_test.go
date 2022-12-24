package binarytree

import (
	"sort"
	"testing"
)

func TestNewTree(t *testing.T) {
	// Type: int
	t.Run("int", func(t *testing.T) {
		defer func() {
			err := recover()
			if err != nil {
				t.Errorf("Unexpected Error: %v", err)
			}
		}()
		_ = NewTree(func(left, right int) bool {
			return left < right
		})
	})
	// Type: float64
	t.Run("float64", func(t *testing.T) {
		defer func() {
			err := recover()
			if err != nil {
				t.Errorf("Unexpected Error: %v", err)
			}
		}()
		_ = NewTree(func(left, right float64) bool {
			return left < right
		})
	})
	// Type: string
	t.Run("string", func(t *testing.T) {
		defer func() {
			err := recover()
			if err != nil {
				t.Errorf("Unexpected Error: %v", err)
			}
		}()
		_ = NewTree(func(left, right string) bool {
			return left < right
		})
	})
}

func TestContains(t *testing.T) {
	testCases := []struct {
		name   string
		args   []int // Elements inserted
		noArgs []int // Elements not found
	}{
		{
			name:   "AllSame",
			args:   []int{1, 1, 1, 1, 1},
			noArgs: []int{2, 4, 6, 8, 10},
		},
		{
			name:   "AllUnique",
			args:   []int{1, 3, 5, 7, 9},
			noArgs: []int{2, 4, 6, 8, 10},
		},
		{
			name:   "NoElement",
			noArgs: []int{2, 4, 6, 8, 10},
		},
		{
			name:   "WideRange",
			args:   []int{1 << 0, 1 << 3, 1 << 6, 1 << 9, 1 << 12, 1 << 15, 1 << 18, 1 << 21, 1 << 24, 1 << 27},
			noArgs: []int{-1 << 0, -1 << 3, -1 << 6, -1 << 9, -1 << 12, -1 << 15, -1 << 18, -1 << 21, -1 << 24, -1 << 27},
		},
		{
			name:   "Random",
			args:   []int{-74046, -109457, 397844, -886408, 733142, 971690, 680540, -22490, -509136, -172753},
			noArgs: []int{-430885, 369708, 642851, 676203, 600807, -20428, -956524, -325873, -724936, 196515},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				err := recover()
				if err != nil {
					t.Errorf("Unexpected Error: %v", err)
				}
			}()

			tree := NewTree(func(left, right int) bool {
				return left < right
			})

			for _, v := range tc.args {
				tree.Push(v)
			}

			// check args
			for _, v := range tc.args {
				if found := tree.Contains(v); !found {
					t.Fatalf("%v should be contained, but not found.", v)
				}
			}

			// check noArgs
			for _, v := range tc.noArgs {
				if found := tree.Contains(v); found {
					t.Fatalf("%v should NOT be contained, but not found.", v)
				}
			}
		})
	}
}

func TestPrevNext(t *testing.T) {
	testCases := []struct {
		name string
		args []int
	}{
		{
			name: "AllSame",
			args: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			name: "AllUnique",
			args: []int{3, 1, 4, 5, 9, 2, 6, 8, 7, 0},
		},
		{
			name: "NoElement",
		},
		{
			name: "WideRange",
			args: []int{1 << 0, 1 << 3, 1 << 6, 1 << 9, 1 << 12, 1 << 15, 1 << 18, 1 << 21, 1 << 24, 1 << 27},
		},
		{name: "Random",
			args: []int{390841, -276234, -58866, 279117, -377391, -507712, 95271, 853932, 582680, -539932},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				err := recover()
				if err != nil {
					t.Errorf("Unexpected Error: %v", err)
				}
			}()

			tree := NewTree(func(left, right int) bool {
				return left < right
			})

			// Prepare Sorted Arguments
			sortedArgs := sort.IntSlice(tc.args)
			sort.Sort(sortedArgs)

			// Compress sorted Args
			// ex.
			// [1,1,1,2,2,2,3] -> [1,2,3]
			// [1,1,1,1,1,1,1] -> [1]
			// []              -> []
			// [1,2,3,4,5,6,7] -> [1,2,3,4,5,6,7]
			for i, j := 1, 0; i <= len(sortedArgs); i++ {
				if i == len(sortedArgs) {
					sortedArgs = sortedArgs[:j+1]
					break
				}

				if sortedArgs[i] != sortedArgs[j] {
					j++
					sortedArgs[j] = sortedArgs[i]
				}
			}

			for _, v := range tc.args {
				tree.Push(v)
			}

			for i, v := range sortedArgs {
				if ok := tree.Contains(v); !ok {
					t.Fatalf("%d should be contained, but not found.", v)
				}
				if i == 0 {
					//Min case
					if _, err := tree.Prev(v); err != ErrNotFound {
						t.Fatalf("Expected %v, got %v", ErrNotFound, err)
					}

				} else if i == len(sortedArgs)-1 {
					//Max case
					if _, err := tree.Next(v); err != ErrNotFound {
						t.Fatalf("Expected %v, got %v", ErrNotFound, err)
					}
				} else {
					if prev, err := tree.Prev(v); err != nil {
						t.Fatalf("Unexpected Error: %v target = %d, %v", err, v, sortedArgs)
					} else if prev != sortedArgs[i-1] {
						t.Fatalf("Expected %v, got %v", sortedArgs[i-1], v)
					}

					if next, err := tree.Next(v); err != nil {
						t.Fatalf("Unexpected Error: %v", err)
					} else if next != sortedArgs[i+1] {
						t.Fatalf("Expected %v, got %v", sortedArgs[i+1], v)
					}
				}
			}
		})
	}
}
func TestPushPop(t *testing.T) {
	testCases := []struct {
		name string
		args []int
	}{
		{
			name: "AllSame",
			args: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			name: "AllUnique",
			args: []int{3, 1, 4, 5, 9, 2, 6, 8, 7, 0},
		},
		{
			name: "NoElement",
		},
		{
			name: "WideRange",
			args: []int{1 << 0, 1 << 3, 1 << 6, 1 << 9, 1 << 12, 1 << 15, 1 << 18, 1 << 21, 1 << 24, 1 << 27},
		},
		{name: "Random",
			args: []int{390841, -276234, -58866, 279117, -377391, -507712, 95271, 853932, 582680, -539932},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				err := recover()
				if err != nil {
					t.Errorf("Unexpected Error: %v", err)
				}
			}()

			tree := NewTree(func(left, right int) bool {
				return left < right
			})

			// Push values
			for _, v := range tc.args {
				tree.Push(v)
			}

			// Pop values
			for _, v := range tc.args {
				if err := tree.Pop(v); err != nil {
					t.Fatal(err)
				}
			}

			// Tree must be empty
			if tree.Len() != 0 {
				t.Errorf("Expected %v, got %v instead.", 0, tree.Len())
			}
		})
	}
}

func TestGetKthElem(t *testing.T) {

	testCases := []struct {
		name string
		args []int
	}{
		{
			name: "AllSame",
			args: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			name: "AllUnique",
			args: []int{3, 1, 4, 5, 9, 2, 6, 8, 7, 0},
		},
		{
			name: "NoElement",
		},
		{
			name: "WideRange",
			args: []int{1 << 0, 1 << 3, 1 << 6, 1 << 9, 1 << 12, 1 << 15, 1 << 18, 1 << 21, 1 << 24, 1 << 27},
		},
		{name: "Random",
			args: []int{390841, -276234, -58866, 279117, -377391, -507712, 95271, 853932, 582680, -539932},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				err := recover()
				if err != nil {
					t.Errorf("Unexpected Error: %v", err)
				}
			}()

			tree := NewTree(func(left, right int) bool {
				return left < right
			})

			// Prepare Sorted Arguments
			sortedArgs := sort.IntSlice(tc.args)
			sort.Sort(sortedArgs)

			for _, v := range tc.args {
				tree.Push(v)
			}

			for i, v := range sortedArgs {
				if get, err := tree.GetKthElem(i); err != nil {
					t.Fatalf("Unexpected Error: %v", err)
				} else if get != v {
					t.Errorf("Expected %d, got %d instead\n", v, get)
				}
			}
		})
	}
}

func TestMinMax(t *testing.T) {
	testCases := []struct {
		name   string
		args   []int
		expMin int
		expMax int
		expLen int
	}{
		{
			name:   "AllSame",
			args:   []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			expMin: 1,
			expMax: 1,
			expLen: 10,
		},
		{
			name:   "AllUnique",
			args:   []int{3, 1, 4, 5, 9, 2, 6, 8, 7, 0},
			expMin: 0,
			expMax: 9,
			expLen: 10,
		},
		{
			name: "NoElement",
		},
		{
			name:   "WideRange",
			args:   []int{1 << 0, 1 << 3, 1 << 6, 1 << 9, 1 << 12, 1 << 15, 1 << 18, 1 << 21, 1 << 24, 1 << 27},
			expMin: 1 << 0,
			expMax: 1 << 27,
			expLen: 10,
		},
		{
			name:   "Random",
			args:   []int{-481433, 994, -836796, 757068, -862856, 737229, 7080, 778258, -128150, -662491},
			expMin: -862856,
			expMax: 778258,
			expLen: 10,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				err := recover()
				if err != nil {
					t.Errorf("Unexpected Error: %v", err)
				}
			}()

			tree := NewTree(func(left, right int) bool {
				return left < right
			})

			for _, v := range tc.args {
				tree.Push(v)
			}

			if tree.Len() != tc.expLen {
				t.Errorf("Expected %d, got %d instead\n", tc.expLen, tree.Len())
			}

			if len(tc.args) == 0 {
				if _, err := tree.Min(); err != ErrTreeEmpty {
					t.Errorf("Expected %v, got %v instead.", ErrTreeEmpty, err)
				}

				if _, err := tree.Max(); err != ErrTreeEmpty {
					t.Errorf("Expected %v, got %v instead.", ErrTreeEmpty, err)
				}

			} else {

				if min, err := tree.Min(); err != nil {
					t.Fatal(err)
				} else if min != tc.expMin {
					t.Errorf("Min: Expected %d, got %d instead\n", tc.expMin, min)
				}

				if max, err := tree.Max(); err != nil {
					t.Fatal(err)
				} else if max != tc.expMax {
					t.Errorf("Max: Expected %d, got %d instead\n", tc.expMax, max)
				}

			}
		})
	}
}

func BenchmarkPushPop(b *testing.B) {

	const nSize int = 200000

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		tree := NewTree(func(left, right int) bool {
			return left <= right
		})

		for j := 0; j < nSize; j++ {
			tree.Push(j)
		}

		for j := 0; j < nSize; j++ {
			if err := tree.Pop(j); err != nil {
				b.Fatal(err)
			}
		}
	}
}
