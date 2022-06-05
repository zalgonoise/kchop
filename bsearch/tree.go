package bsearch

// Node struct defines a basic binary search tree object, which includes both the
// value and its index (in the original slice).
//
// The Node struct will contain a pointer to a small and a large Node, which will
// allow it to branch evenly as new values are added, also allowing an efficient search
// algorithm
type Node struct {
	idx   int
	val   int
	small *Node
	large *Node
}

// Ordered function will create a new node from an ordered slice of integers. It will
// take the value in the middle of the slice as its root, and insert the remaining values
// (as well as their indexes) as child nodes.
func Ordered(v ...int) *Node {
	if len(v) == 0 {
		return nil
	}

	small := v[:len(v)/2]
	large := v[len(v)/2:]

	node := &Node{
		idx: len(small),
		val: large[0],
	}

	for idx := 1; idx < len(large); idx++ {
		node.Insert(large[idx], idx+len(small))
	}

	for idx := len(small) - 1; idx >= 0; idx-- {
		node.Insert(small[idx], idx)
	}

	return node
}

// Insert method will take a value and its index (in the original slice), and store it
// in Node n.
//
// If the value is greater than the node's value, it will insert it in the large child-node
// (creating it if nil). If the value is smaller than the node's value, it will do the same
// but in the small child-node.
func (n *Node) Insert(v, idx int) {
	if v > n.val {
		if n.large == nil {
			n.large = &Node{val: v, idx: idx}
			return
		}
		n.large.Insert(v, idx)

	} else if v < n.val {
		if n.small == nil {
			n.small = &Node{val: v, idx: idx}
			return
		}
		n.small.Insert(v, idx)

	}
}

// Search method will take a target integer and navigate through the tree to find a match
// if there is one, it will return the node's stored index value; otherwise, returns -1
func (n *Node) Search(target int) (idx int) {
	if target == n.val {
		return n.idx
	}

	if target > n.val && n.large != nil {
		return n.large.Search(target)
	} else if target < n.val && n.small != nil {
		return n.small.Search(target)
	}

	return -1
}
