package config

import "strconv"

type GeneticConfig struct {
	PopulationSize   int
	GeneSize         int
	GenerationCount  int
	MutationRate     float32
	FitnessThreshold int
}

const defaultPopulationSize = 100
const defaultGeneSize = 10
const defaultGenerationCount = 100
const defaultMutationRate = 0.1
const defaultFitnessThreshold = 3

func ParseConfig(args []string) *GeneticConfig {

	populationSize, err := strconv.Atoi(args[1])
	if err != nil {
		populationSize = defaultPopulationSize
	}

	geneSize, err := strconv.Atoi(args[2])
	if err != nil {
		geneSize = defaultGeneSize
	}

	generationCount, err := strconv.Atoi(args[3])
	if err != nil {
		generationCount = defaultGenerationCount
	}

	mutationRate, err := strconv.ParseFloat(args[4], 32)
	if err != nil {
		mutationRate = defaultMutationRate
	}

	fitnessThreshold, err := strconv.Atoi(args[5])
	if err != nil {
		fitnessThreshold = defaultFitnessThreshold
	}

	return &GeneticConfig{
		PopulationSize:   populationSize,
		GeneSize:         geneSize,
		GenerationCount:  generationCount,
		MutationRate:     float32(mutationRate),
		FitnessThreshold: fitnessThreshold,
	}
}
