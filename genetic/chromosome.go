package genetic

type Chromosome struct {
	Genes   []*Gene
	Fitness int
}

func initChromosome(geneSize int) *Chromosome {
	genes := make([]*Gene, geneSize)
	for range genes {
		g := NewGene()
		genes = append(genes, g)
	}

	return &Chromosome{
		Genes:   genes,
		Fitness: 0,
	}
}
