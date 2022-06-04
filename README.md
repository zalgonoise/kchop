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

1. [Iterative](https://github.com/zalgonoise/kchop/blob/master/iterative/chop.go#L7)

This implementation, as suggested in the original document, will perform a binary search using a classic iterative approach, within a function.

This approach seems very verbose and long, with a lot of complexity in the loops / conditionals it contains. Although this is diminished a bit with added documentation, it still makes it hard to review and follow.

As a first approach, it seems straight-forward and simple to write, however -- so this took the least amount of time to write and get working; with no major issues or blockers.

2. [Recursive](https://github.com/zalgonoise/kchop/blob/master/recursive/chop.go#L11)

Also suggested in the original document, this implementation will focus on a recursive pattern for the same algorithm. The way this is done is by adding a new (private) function with a slightly different signature:

```go
chop(t int, v []int, c int) int
```

This signature will be close to the original signature (for the `Chop()` function), adding one more parameter for the carrying index value. The idea is to break-down the slice with a binary search pattern until it only holds one value, which is compared to the target.

Implementing this algorithm required more fine-tuning by the end of the first draft, as this one started to raise _off-by-one_ errors. This was easily corrected by ensuring that, if the target value ends up in the second, bigger slice, [the carrying index must be incremented by the length of the first, smaller slice](https://github.com/zalgonoise/kchop/blob/master/recursive/chop.go#L57).

All in all, this approach seems to bring more readability, and also a bit more compact than the Iterative solution. Once it's working it feels more like magic and less like a repetitive action.

3. [Functional](https://github.com/zalgonoise/kchop/blob/master/functional/chop.go#L7)

Similarly to the previous implementation, the original document also suggests a functional approach to this algorithm. As such, we want to keep in mind the statlessness / immutability in functional programming, taking functions as first-class citizens -- also while keeping it simple and not excessively abstract or complex.

For this task, three functions are introduced: `reduce()`, `reduceR()` and `match()`. Here are their signatures: 

```go
reduce(t int, v []int) (s []int, c int)
reduceR(t, c int, v []int) (s []int, i int)
match(t, c int, v []int) int
```

The first two functions are to approach the same task -- the actual binary search, returning a new, smaller slice, recursively (a hint from the previous implementation). It also returns an index carry value to serve as reference in case there is a match. This process is repeated until the slice only contains zero or one elements, when it exits.

Lastly, these values are fed into the `match()` function, which takes in the target value, the index carry value, and the (new, single or no-value) list to match its content to the target. If there are no values in the list, it returns `-1`. If there is an element in the list and it doesn't match the target, returns `-1`. If it matches, however, it will return the carry value.

There weren't any major hiccups when writing this implementation as the process in itself is very straight-forward. A great thing about this pattern is readability -- surely it will be one of the most easily readable and testable solutions. Its downsides will surely be performance with the number of operations and allocations involved.