package inspector

// Range struct defines the basic start and end points of a range, as index values in
// a slice, array or map
type Range struct {
	start int
	end   int
}

// Diff method will return the difference between the start and end points in the range
func (r *Range) Diff() int {
	return (r.end - r.start)
}

// Mid method will return the mid-point in the range
func (r *Range) Mid() int {
	return (r.end - r.start) / 2
}

// Smaller method will return a new range, half the size of the original, returning a
// Range of its lower half. This follows the basic binary search algorithm pattern.
func (r *Range) Smaller() *Range {
	new := &Range{
		start: r.start,
		end:   r.Mid(),
	}
	return new
}

// Larger method will return a new range, half the size of the original, returning a
// Range of its upper half. This follows the basic binary search algorithm pattern.
func (r *Range) Larger() *Range {
	new := &Range{
		start: r.Mid() + 1,
		end:   r.end,
	}
	return new
}

// search function will be called by Search(), however this is a recursive function.
//
// The input pointer to a Range is used for comparison, and if the target value cannot be
// found (but is greater or smaller than the mid value), a new range is built taking into
// account the binary search pattern, and a new search() is called with the new Range.
//
// If there is a match, however, it returns its index. If there isn't one, it will return
// -1
func search(t int, v []int, r *Range) int {
	if r.Diff() <= 2 {
		if t == v[r.start] {
			return r.start
		} else if t == v[r.end] {
			return r.end
		} else if t == v[r.Mid()] {
			return r.Mid()
		}
		return -1
	}

	if t == v[r.Mid()] {
		return r.Mid()
	} else if t > v[r.Mid()] {
		return search(t, v, r.Larger())
	} else if t < v[r.Mid()] {
		return search(t, v, r.Smaller())
	}

	return -1

}

// Search function will take in a target integer and a list of (ordered) integers,
// to find a match in the slice, and returning its index.
//
// It does so by initializing a Range pointer with a start of 0 and end of the length
// of the slice of values (minus one), and then calling the private function `search()`
// with this Range.
func Search(t int, v []int) int {
	if len(v) == 0 {
		return -1
	}

	r := &Range{
		start: 0,
		end:   len(v) - 1,
	}
	return search(t, v, r)
}
