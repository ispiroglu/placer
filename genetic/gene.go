package genetic

import "github.com/ispiroglu/placer/file"

type Gene struct {
	*file.Rectangle
}

func NewGene(rect *file.Rectangle) *Gene {
	return &Gene{
		Rectangle: rect,
	}
}
