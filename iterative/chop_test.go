package iterative

import (
	"reflect"
	"testing"
)

func TestChop(t *testing.T) {
	module := "Iterative"
	funcname := "Chop(int, []int) int"

	type test struct {
		target int
		values []int
		wants  int
	}

	var tests = []test{
		// idx: 0
		{
			target: 3,
			values: []int{},
			wants:  -1,
		},
		{
			target: 3,
			values: []int{1},
			wants:  -1,
		},
		{
			target: 1,
			values: []int{1},
			wants:  0,
		},
		// idx: 3
		{
			target: 1,
			values: []int{1, 3, 5},
			wants:  0,
		},
		{
			target: 3,
			values: []int{1, 3, 5},
			wants:  1,
		},
		{
			target: 5,
			values: []int{1, 3, 5},
			wants:  2,
		},
		// idx: 6
		{
			target: 0,
			values: []int{1, 3, 5},
			wants:  -1,
		},
		{
			target: 2,
			values: []int{1, 3, 5},
			wants:  -1,
		},
		{
			target: 4,
			values: []int{1, 3, 5},
			wants:  -1,
		},
		// idx: 9
		{
			target: 6,
			values: []int{1, 3, 5},
			wants:  -1,
		},
		{
			target: 1,
			values: []int{1, 3, 5, 7},
			wants:  0,
		},
		{
			target: 3,
			values: []int{1, 3, 5, 7},
			wants:  1,
		},
		// idx: 12
		{
			target: 5,
			values: []int{1, 3, 5, 7},
			wants:  2,
		},
		{
			target: 7,
			values: []int{1, 3, 5, 7},
			wants:  3,
		},
		{
			target: 0,
			values: []int{1, 3, 5, 7},
			wants:  -1,
		},
		// idx: 15
		{
			target: 2,
			values: []int{1, 3, 5, 7},
			wants:  -1,
		},
		{
			target: 4,
			values: []int{1, 3, 5, 7},
			wants:  -1,
		},
		{
			target: 6,
			values: []int{1, 3, 5, 7},
			wants:  -1,
		},
		// idx: 18
		{
			target: 8,
			values: []int{1, 3, 5, 7},
			wants:  -1,
		},
		{
			target: 1,
			values: []int{1, 3, 5, 7, 9},
			wants:  0,
		},
		{
			target: 3,
			values: []int{1, 3, 5, 7, 9},
			wants:  1,
		},
		// idx: 21
		{
			target: 5,
			values: []int{1, 3, 5, 7, 9},
			wants:  2,
		},
		{
			target: 7,
			values: []int{1, 3, 5, 7, 9},
			wants:  3,
		},
		{
			target: 9,
			values: []int{1, 3, 5, 7, 9},
			wants:  4,
		},
		// idx: 24
		{
			target: 0,
			values: []int{1, 3, 5, 7, 9},
			wants:  -1,
		},
		{
			target: 2,
			values: []int{1, 3, 5, 7, 9},
			wants:  -1,
		},
		{
			target: 4,
			values: []int{1, 3, 5, 7, 9},
			wants:  -1,
		},
		// idx: 27
		{
			target: 6,
			values: []int{1, 3, 5, 7, 9},
			wants:  -1,
		},
		{
			target: 8,
			values: []int{1, 3, 5, 7, 9},
			wants:  -1,
		},
		{
			target: 10,
			values: []int{1, 3, 5, 7, 9},
			wants:  -1,
		},
	}

	for idx, test := range tests {
		result := Chop(test.target, test.values)

		if !reflect.DeepEqual(test.wants, result) {
			t.Errorf(
				"#%v -- FAILED -- [%s] [%s] output mismatch error: wanted %v ; got %v",
				idx,
				module,
				funcname,
				test.wants,
				result,
			)
		}
	}
}
