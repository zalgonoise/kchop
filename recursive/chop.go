package recursive

// Chop function will take a target integer and a list of ordered values,
// and return the index of the target in the slice if it exists.
//
// Otherwise, it returns -1
//
// It does so by calling a recursive `chop()` function, which will perform
// a binary search on the slice by splitting it in half, and working with the
// one that matches the target value the closest.
func Chop(target int, values []int) int {
	// short-circuit on zero-length input slice
	if len(values) == 0 {
		return -1
	}

	// short-circuit on one-length  input slice
	if len(values) == 1 {
		if target == values[0] {
			return 0
		}
		return -1
	}

	return chop(target, values, 0)
}

// chop function will recursively be called with different `v []int` slices.
//
// This function will also take an extra parameter (for the carrying index) which
// makes up for the skipped indexes when breaking down larger arrays / slices.
//
// It always returns the index of the match or -1 if no matches are found, so
// `Chop()` will simply call this function with its input, and a carrying index of 0.
func chop(t int, v []int, c int) int {
	var small = []int{}
	var big = []int{}

	// split the slice in two
	small = v[0:(len(v) / 2)]
	big = v[(len(v) / 2):]

	// check if target should be in the smaller values' slice
	if len(big) > 0 && t < big[0] {

		// rinse and repeat if slice length is 2 or more
		if len(small) > 1 {
			return chop(t, small, c)
		}

		// if value matches, return the carrying index value
		if len(small) == 1 && small[0] == t {
			return c
		}

		// no matches found
		return -1
	}

	// target should be in the big values' slice, or it's a mismatch
	// rinse and repeat if slice length is 2 or more
	// increment the carrying index with the length of the "small" slice
	if len(big) > 1 {
		return chop(t, big, c+len(small))
	}

	// if value matches, return the carrying index value, summing the
	// length of the "small" slice
	if len(big) == 1 && big[0] == t {
		return c + len(small)
	}

	// no matches found
	return -1

}
