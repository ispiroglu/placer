package genetic

import (
	"github.com/ispiroglu/placer/area"
)

type Chromosome struct {
	Fitness int
	area    area.Area
}

/*
   !Chromosome should represent area.
*/

// ? Do I need geneSize? len(area.Rectangles) should be enough?
func initChromosome(geneSize int, a area.Area) *Chromosome {
	a.BottomLeftFill()

	return &Chromosome{
		Fitness: 0,
		area:    a,
	}
}
