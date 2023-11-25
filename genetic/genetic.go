package genetic

import (
	"github.com/ispiroglu/placer/area"
	"github.com/ispiroglu/placer/config"
)

func Calculate(cfg *config.GeneticConfig, a area.Area) *Chromosome {
	p := initPopulation(cfg.PopulationSize, cfg.GeneSize, a)

	for i := 0; i < cfg.GenerationCount; i++ {
		p.calculateFitness()
		p.sort()

		survivors := p.evaluateSurvivors()
		var newGeneration []*Chromosome
		newGenCount := len(p.chromosomes) - len(survivors)
		for len(newGeneration) < newGenCount {
			c1, c2 := p.selectParents()

			child := crossover(c1, c2)
			child.mutate(cfg.MutationRate)

			newGeneration = append(newGeneration, child)
		}

		p.chromosomes = append(newGeneration, survivors...)
	}

	p.sort()

	return p.chromosomes[0]
}
