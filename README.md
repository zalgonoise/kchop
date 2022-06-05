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
func chop(t int, v []int, c int) int
```

This signature will be close to the original signature (for the `Chop()` function), adding one more parameter for the carrying index value. The idea is to break-down the slice with a binary search pattern until it only holds one value, which is compared to the target.

Implementing this algorithm required more fine-tuning by the end of the first draft, as this one started to raise _off-by-one_ errors. This was easily corrected by ensuring that, if the target value ends up in the second, bigger slice, [the carrying index must be incremented by the length of the first, smaller slice](https://github.com/zalgonoise/kchop/blob/master/recursive/chop.go#L57).

All in all, this approach seems to bring more readability, and also a bit more compact than the Iterative solution. Once it's working it feels more like magic and less like a repetitive action.

3. [Functional](https://github.com/zalgonoise/kchop/blob/master/functional/chop.go#L7)

Similarly to the previous implementation, the original document also suggests a functional approach to this algorithm. As such, we want to keep in mind the statlessness / immutability in functional programming, taking functions as first-class citizens -- also while keeping it simple and not excessively abstract or complex.

For this task, three functions are introduced: `reduce()`, `reduceR()` and `match()`. Here are their signatures: 

```go
func reduce(t int, v []int) (s []int, c int)
func reduceR(t, c int, v []int) (s []int, i int)
func match(t, c int, v []int) int
```

The first two functions are to approach the same task -- the actual binary search, returning a new, smaller slice, recursively (a hint from the previous implementation). It also returns an index carry value to serve as reference in case there is a match. This process is repeated until the slice only contains zero or one elements, when it exits.

Lastly, these values are fed into the `match()` function, which takes in the target value, the index carry value, and the (new, single or no-value) list to match its content to the target. If there are no values in the list, it returns `-1`. If there is an element in the list and it doesn't match the target, returns `-1`. If it matches, however, it will return the carry value.

There weren't any major hiccups when writing this implementation as the process in itself is very straight-forward. A great thing about this pattern is readability -- surely it will be one of the most easily readable and testable solutions. Its downsides will surely be performance with the number of operations and allocations involved.

4. [Binary Search Tree (OOP)](https://github.com/zalgonoise/kchop/blob/master/bsearch/chop.go#L12)

This implementation will focus on the actual data structure that is allowing this type of query to be efficient: binary search trees.

For this, a struct with two methods is introduced, as well as an initializer function:

```go
type Node struct {
	idx   int
	val   int
	small *Node
	large *Node
}

func (n *Node) Insert(v, idx int)
func (n *Node) Search(target int) (idx int)

func Ordered(v ...int) *Node
```

The downside of this particular approach is that the `Chop()` function will first load all the values into memory (as a pointer to a `Node` object), and then perform a query on it. The query in itself will be super efficient and reliable -- so this type of implementation would be more useful if the interactions happened with the `Node` object directly, instead of loading the data into memory on each query.

In terms of accuracy, I didn't find any _off-by-one_ issues when writing this implementation. What I struggled the most was to find a decent `init()` function for it, considering the input was an ordered list -- if it was unordered, the `Node` object could simply be initialized with the first element of the slice as its root, and then calling the `Insert()` method on each remaining value in the slice.

Being an ordered list of integers, the approach above would make the whole implementation useless as it the binary search tree implementation would not benefit at all in comparison to an purely iterative approach:

```
[1]
 +- small: <nil>
 +- large: [3]
            +- small: <nil>
            +- large: [5]
                       +- small: <nil>
                       +- large: [7]
```

Since the same would occur if inserting while iterating from the end of the list to the beginning, the root is initialized with:
- splitting the input slice in two: smaller and larger slices.
- taking the larger slice's first item as the root
- iterating upwards through the remaining large items, inserting them into the `Node`
- iterating downwards through the small slice of items, inserting them into the `Node`

5. [Index-Based](https://github.com/zalgonoise/kchop/blob/master/inspector/chop.go#L7)

This last approach took some time to think of, to provide an original approach that also worked, and wasn't a refactored copy (or just a mix) of the other approaches above.

In this approach I figured that in all approaches the indexes are being more or less disregarded -- or in other words, the index is the cumbersome part in the logic where it needs to either be carried, or is the always mutable piece of data (at a first glance).

As such, why not simply base the whole logic off of the index of the slice? That is where this _inspector_ approach comes in.

It introduces a `Range` struct to define start and end points for the look-up, and four methods: `Diff()` to return the difference of the start and end points (how many items in the range), `Mid()` to return the index of the middle position in the range, and two similar methods: `Smaller()` and `Larger()`. These last two will apply the binary search pattern by reducing the range according to the value comparison:


```go
type Range struct {
	start int
	end   int
}

func (r *Range) Diff() int
func (r *Range) Mid() int
func (r *Range) Smaller() *Range
func (r *Range) Larger() *Range
```

There are also two other functions, one public (`Search()`) and a similarly named private one. The public function is simplified for external use (checking the input slice's length, initializing the "default" range) while the private function will recursively search through the slice, applying a binary search pattern to reduce the range.

it will recursively break the range down until its difference is 2 or lower, where it finally compares it the _final_ values, returning the appropriate index on a match or `-1` if non-existent. 

```go
func search(t int, v []int, r *Range) int
func Search(t int, v []int) int
```