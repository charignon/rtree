// Package rtree provides RTree 2D Indexing Data Structures
//
// Reference: https://en.wikipedia.org/wiki/R-tree
// The R*Tree is an improved version of the RTree, details
// at https://en.wikipedia.org/wiki/R*_tree
//
// Example:
// Building a RTree with some values
//  r := rtree.NewRegularRTree().WithCapacity(3)
//  r.Insert(rtree.Rect{1, 2, 3, 4}, "A")
//  r.Insert(rtree.Rect{1, 5, 12, 23}, "B")
//  r.Insert(rtree.Rect{5, 10, 5, 10}, "C")
//  r.Insert(rtree.Rect{500, 560, 23, 24}, "D")
//  r.Insert(rtree.Rect{-200, -100, 50, 10}, "E")
// Finding all the entries in the rect [0,20], [0,20]:
//  fmt.Println(rtree.Search({0,20,0,20}))
package rtree

type RTree interface {
	Insert(r Rect, v interface{})
	Search(r Rect) []interface{}
}
