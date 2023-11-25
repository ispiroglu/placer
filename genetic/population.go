package genetic

import (
	"math/rand"
	"sort"

	"github.com/ispiroglu/placer/area"
)

type population struct {
	chromosomes []*Chromosome
}

func initPopulation(populationSize, geneSize int, a *area.Area) *population {
	chromosomes := make([]*Chromosome, populationSize)

	// TODO: this should be done in parallel
	for range chromosomes {
		c := initChromosome(geneSize, a)
		chromosomes = append(chromosomes, c)
	}

	return &population{
		chromosomes: chromosomes,
	}
}

func (p *population) calculateFitness() {
	for _, c := range p.chromosomes {
		c.calculateFitness()
	}
}

func (p *population) evaluateSurvivors() []*Chromosome {
	survivorSize := int(0.1 * float32(len(p.chromosomes)))
	survivors := make([]*Chromosome, survivorSize)
	copy(survivors, p.chromosomes[:survivorSize])

	return survivors
}

// tournament selection
func (p *population) selectParents() (*Chromosome, *Chromosome) {
	tournamentSize := int(0.3 * float32(len(p.chromosomes))) // TODO: how to determine tournament size?
	var p1, p2 *Chromosome

	for i := 0; i < tournamentSize; i++ {
		candidate := p.chromosomes[rand.Intn(len(p.chromosomes))]
		if p1 == nil || candidate.Fitness > p1.Fitness {
			p1 = candidate
		}
	}

	for i := 0; i < tournamentSize; i++ {
		candidate := p.chromosomes[rand.Intn(len(p.chromosomes))]
		if candidate != p1 && (p2 == nil || candidate.Fitness > p2.Fitness) {
			p2 = candidate
		}
	}

	return p1, p2
}

func (p *population) sort() {
	sort.SliceStable(p, p.compareFn)
}

func (p *population) compareFn(i, j int) bool {
	return p.chromosomes[i].Fitness < p.chromosomes[j].Fitness
}
