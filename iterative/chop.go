package iterative

// Chop function will take a target integer and a list of ordered values,
// and return the index of the target in the slice if it exists.
//
// Otherwise, it returns -1
func Chop(target int, values []int) (index int) {

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

	var i uint
	var pairs uint = uint(len(values) / 2)

	// make a copy of the input to work with, safely
	input := values

	// initialize a counter to make up for the values "skipped"
	var indexCounter int = 0

	// iterate through all potential pairs
	for i = 0; i < pairs; i++ {
		var small = []int{}
		var big = []int{}

		// split the slice in two
		small = input[0:(len(input) / 2)]
		big = input[(len(input) / 2):]

		// check if each big / small slice is not just a single value
		if len(small) > 1 && len(big) > 1 {

			// check if the target value is greater than the biggest "small" value
			if target > small[len(small)-1] {

				// make up for the skipped values by adding this slice's length to
				// the counter; then continue working on the "big" slice
				indexCounter = indexCounter + len(small)
				input = big
			} else {

				// doesn't need to make up for in the index, continue working on the
				// "small" slice
				input = small
			}
		} else {
			// either or both slices only contains one value; check if it matches
			// the target
			for idx, val := range small {
				if val == target {
					return idx + indexCounter
				}
			}

			// if it doesn't, increment the counter once with the "small" slice's length
			indexCounter = indexCounter + len(small)

			for idx, val := range big {
				if val == target {
					return idx + indexCounter
				}
			}
		}
	}

	// value wasn't found in the input slice; return -1
	return -1
}
