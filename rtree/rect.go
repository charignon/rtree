package rtree

import "math"

/* Represents a 2D Rectangle */
type Rect struct {
	Xmin float64
	Xmax float64
	Ymin float64
	Ymax float64
}

func (r Rect) area() float64 {
	return (r.Xmax - r.Xmin) * (r.Ymax - r.Ymin)
}

func (r1 Rect) IntersectP(r2 Rect) bool {
	return r1.Xmin < r2.Xmax && r1.Xmax > r2.Xmin && r1.Ymin < r2.Ymax && r1.Ymax > r2.Ymin
}

func (r1 Rect) intersect(r2 Rect) bool {
	return r1.Xmin < r2.Xmax && r1.Xmax > r2.Xmin && r1.Ymin < r2.Ymax && r1.Ymax > r2.Ymin
}
func (r1 Rect) union(r2 Rect) Rect {
	return Rect{
		math.Min(r1.Xmin, r2.Xmin),
		math.Max(r1.Xmax, r2.Xmax),
		math.Min(r1.Ymin, r2.Ymin),
		math.Max(r1.Ymax, r2.Ymax),
	}
}
