package genetic

import (
	"github.com/ispiroglu/placer/config"
	"github.com/ispiroglu/placer/file"
)


func Calculate(cfg *config.GeneticConfig, env *file.Environment) *Chromosome {
	p := initPopulation(cfg.PopulationSize, cfg.GeneSize)

	for i := 0; i < cfg.GenerationCount; i++ {
		p.calculateFitness()
		p.sort()

		survivors := p.evaluateSurvivors()
		var newGeneration []*Chromosome

		for len(newGeneration) < len(p.chromosomes)-len(survivors) {
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
