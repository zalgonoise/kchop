# kchop
Several implementations of a prosaic binary search, as part of [Kata02](http://codekata.com/kata/kata02-karate-chop/)


___________

### Targets

- A function that takes as input a target integer value, and a slice of ordered integers
- This function will look up if the target value is present in the slice
- It implements a basic prosaic binary search, where the slice is split in half and the target value is matched against the highest and lowest values of each smaller and bigger half, accordingly. This process repeats itself until either a match is found, or it's assured that there isn't a match.
- This function will return the index of the target value in the slice, or `-1` if it is not present in it.


____________

### Implementations

1. Iterative

This implementation, as suggested in the original document, will perform a binary search using a classic iterative approach, within a function.

