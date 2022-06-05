package bsearch

// Chop function will take a target integer and a list of ordered values,
// and return the index of the target in the slice if it exists.
//
// Otherwise, it returns -1
//
// It does so loading the entire slice of values into memory as a binary search tree,
// making search actions more efficient. Ideally, this implementation would continue
// to follow an OOP design where the data isn't being constantly loaded again and again,
// but the user works with the Node struct directly to perform search queries.
func Chop(target int, values []int) int {
	if n := Ordered(values...); n != nil {
		return n.Search(target)
	}
	return -1
}
