package inspector

// Chop function will take a target integer and a list of ordered values,
// and return the index of the target in the slice if it exists.
//
// Otherwise, it returns -1
func Chop(target int, values []int) (index int) {
	return Search(target, values)
}
