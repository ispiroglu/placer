package genetic

type Chromosome struct {
	genes   []*gene
	fitness int
}

func initChromosome(geneSize int) *Chromosome {
	genes := make([]*gene, geneSize)
	for range genes {
		g := newGene()
		genes = append(genes, g)
	}

	return &Chromosome{
		genes:   genes,
		fitness: 0,
	}
}
