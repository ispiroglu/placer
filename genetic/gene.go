package genetic

import "github.com/ispiroglu/placer/area"

type Gene interface {
	Point() *area.Point
}

// Can gene be interface?
/*func NewGene(rect *area.Rectangle) *Gene {
	return &Gene{
		Rectangle: rect,
	}
}*/
