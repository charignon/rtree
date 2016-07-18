package rtree

var capacity = 3

func NewRegularRTree() *RegularRTree {
	return &RegularRTree{&node{}, 3}
}

func (r *RegularRTree) WithCapacity(capacity int) *RegularRTree {
	r.capacity = capacity
	return r
}

// Insert a new entry in the RTree, the entry has a surface (Rect) and
// a value
func (r *RegularRTree) Insert(re Rect, v interface{}) {
	n := &node{re, nil, nil, v}
	leaf := pickLeaf(r, n)
	var split_result *node
	if len(leaf.children) == r.capacity {
		split_result = leaf.split(n)
	} else {
		leaf.addChild(n)
	}
	Root_splitted := r.adjust(leaf, split_result)
	if Root_splitted != nil {
		oldRoot := r.Root
		r.Root = &node{}
		r.Root.addChild(oldRoot)
		r.Root.addChild(Root_splitted)
	}
}

type searchres struct {
	results []interface{}
}

func (n *node) searchEntries(re Rect, results *searchres) {
	if n.children == nil && n.r.intersect(re) {
		results.results = append(results.results, n.value)
		return
	}
	for _, c := range n.children {
		if c.r.intersect(re) {
			c.searchEntries(re, results)
		}
	}

}

// Search for an entry in the RTree
func (r *RegularRTree) Search(re Rect) []interface{} {
	results := make([]interface{}, 0, 0)
	res := searchres{results}
	r.Root.searchEntries(re, &res)
	return res.results
}

func isLeaf(n *node) bool {
	return len(n.children) == 0 || n.children[0].children == nil
}

func pickClosestChild(parent *node, tosearch *node) *node {
	// Find the child that will lead to the minimum enlargment
	// Initialize with the worst case
	min_enlargment := parent.r.union(tosearch.r).area()
	var chosen *node
	for _, child := range parent.children {
		enlargment := child.r.union(tosearch.r).area()
		if enlargment <= min_enlargment {
			chosen = child
			min_enlargment = enlargment
		}
	}
	return chosen
}

func pickLeaf(r *RegularRTree, n *node) *node {
	current := r.Root
	for !isLeaf(current) {
		current = pickClosestChild(current, n)
	}
	return current
}

func (r *RegularRTree) adjust(leaf *node, split_result *node) *node {
	if leaf == r.Root {
		return split_result
	}
	leaf.parent.recomputeRect()
	if split_result == nil {
		return r.adjust(leaf.parent, nil)
	}
	if len(leaf.children) == r.capacity {
		to_split := leaf.parent
		splitted := to_split.split(split_result)
		return r.adjust(to_split, splitted)
	} else {
		leaf.parent.addChild(split_result)
		return r.adjust(leaf.parent, nil)
	}
}

type node struct {
	r        Rect
	parent   *node
	children []*node     // nil for entities
	value    interface{} // nil for internal node
}

func (n *node) addChild(child *node) {
	n.children = append(n.children, child)
	child.parent = n
	n.r = n.r.union(child.r)
}

func (n *node) recomputeRect() {
	if len(n.children) == 0 {
		return
	}
	n.r = n.children[0].r
	for _, c := range n.children[1:] {
		n.r = n.r.union(c.r)
	}
}

func (n *node) removeChild(childtoremove *node) {
	newchildren := make([]*node, 0, 0)
	for _, child := range n.children {
		if child != childtoremove {
			newchildren = append(newchildren, child)
		}
	}
	n.recomputeRect()
	n.children = newchildren
}

func (n *node) split(toadd *node) *node {
	child := pickClosestChild(n, toadd)
	n.removeChild(child)
	new_child := &node{}
	new_child.addChild(child)
	new_child.addChild(toadd)
	return new_child
}

// RegularRTree is the most basic kind of RTrees
type RegularRTree struct {
	Root     *node
	capacity int
}
