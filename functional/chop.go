package functional

// Chop function will take a target integer and a list of ordered values,
// and return the index of the target in the slice if it exists.
//
// Otherwise, it returns -1
func Chop(target int, values []int) (index int) {
	slice, carry := reduceR(target, 0, values)

	return match(target, carry, slice)
}

// reduce function will take in a target integer and an ordered list of integers,
// applying a binary search technique to match in which slice should the target value
// be present.
//
// It will return a new list of integers (the smaller or the bigger half) as well as
// and index carry, to account for skipped values as it traverses through the list.
func reduce(t int, v []int) (s []int, c int) {
	var small = []int{}
	var big = []int{}

	// split the slice in two
	small = v[0:(len(v) / 2)]
	big = v[(len(v) / 2):]

	if big[0] > t {
		return small, 0
	}
	return big, len(small)
}

// reduceR function will recursively apply the reduce() function to the input values,
// for as long as the list's length is greater than one single value.
//
// It returns a slice of integers that contains either zero or one elements, and the
// index carry value to consider for the skipped indexes.
func reduceR(t, c int, v []int) (s []int, i int) {
	if len(v) > 1 {
		s, i := reduce(t, v)
		return reduceR(t, i+c, s)
	}
	return v, c
}

// match function will take in a target integer value, an index carry value and a
// list of ints (that may contain zero or one elements), and match the target value
// to the list's contents.
//
// If there is a match, the index carry value is returned. If the slice is zero-length
// or its value does not match the target, it returns -1.
func match(t, c int, v []int) int {
	if len(v) == 0 {
		return -1
	}

	if v[0] == t {
		return c
	}

	return -1
}
