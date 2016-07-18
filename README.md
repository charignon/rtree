RTrees in golang

Example use:
```
package main

import (
	"fmt"

	"github.com/charignon/rtree/rtree"
)

func main() {
	r := rtree.NewRegularRTree().WithCapacity(3)
	r.Insert(rtree.Rect{1, 2, 3, 4}, "A")
	r.Insert(rtree.Rect{1, 5, 12, 23}, "B")
	r.Insert(rtree.Rect{5, 10, 5, 10}, "C")
	r.Insert(rtree.Rect{500, 560, 23, 24}, "D")
	r.Insert(rtree.Rect{-200, -100, 50, 10}, "E")
	fmt.Println(r.Search(rtree.Rect{0, 20, 0, 20}))
}
```
