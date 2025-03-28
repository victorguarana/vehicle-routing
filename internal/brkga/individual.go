package brkga

import "math/rand"

type Chromossome float64

type Individual struct {
	Chromosomes []*Chromossome
	Score       float64
}

func newIndividual(chromosomesLen int) *Individual {
	return &Individual{
		Chromosomes: make([]*Chromossome, chromosomesLen),
	}
}

func newMutantIndividual(chromosomesLen int) *Individual {
	individual := &Individual{
		Chromosomes: make([]*Chromossome, chromosomesLen),
	}

	for i := range individual.Chromosomes {
		chromossome := Chromossome(rand.Float64())
		individual.Chromosomes[i] = &chromossome
	}
	return individual
}

func (c Chromossome) Gene() float64 {
	return float64(c)
}
