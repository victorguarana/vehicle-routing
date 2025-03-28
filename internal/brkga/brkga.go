package brkga

import (
	"math/rand"
	"sort"
)

//go:generate mockgen -source=brkga.go -destination=brkgamock_test.go -package=brkga
type IDecoder[T any] interface {
	Decode(*Individual) T
}

type IMeasurer[T any] interface {
	Measure(T) float64
}

type BRKGAParams[T any] struct {
	MaxPop              int
	TopPercentage       float64
	CrossoverPercentage float64
	BiasPercentage      float64
	ChromosomeLen       int
	GenerationLimit     int
	Decoder             IDecoder[T]
	Measurer            IMeasurer[T]
}

type BRKGA[T any] struct {
	chromosomeLen   int
	maxPop          int
	topQnt          int
	mutantQnt       int
	biasPercentage  float64
	generationLimit int

	decoder  func(*Individual) T
	measurer func(T) float64
}

func NewBRKGA[T any](params BRKGAParams[T]) BRKGA[T] {
	topQnt := calculateQuantity(params.MaxPop, params.TopPercentage)
	crossoverQnt := calculateQuantity(params.MaxPop, params.CrossoverPercentage)
	mutantQnt := params.MaxPop - (topQnt + crossoverQnt)
	return BRKGA[T]{
		chromosomeLen:   params.ChromosomeLen,
		maxPop:          params.MaxPop,
		topQnt:          topQnt,
		mutantQnt:       mutantQnt,
		generationLimit: params.GenerationLimit,
		biasPercentage:  params.BiasPercentage,

		measurer: params.Measurer.Measure,
		decoder:  params.Decoder.Decode,
	}
}

func (b BRKGA[T]) Execute() T {
	var prevGeneration []*Individual
	currentGeneration := b.createInitialGeneration()
	b.evaluateGeneration(currentGeneration)
	b.orderGeneration(currentGeneration)
	generationCounter := 0

	for {
		generationCounter++
		prevGeneration = currentGeneration
		currentGeneration = b.newGeneration(prevGeneration)
		b.evaluateGeneration(currentGeneration)
		b.orderGeneration(currentGeneration)
		if generationCounter >= b.generationLimit {
			return b.decoder(currentGeneration[0])
		}
	}
}

func (b BRKGA[T]) createInitialGeneration() []*Individual {
	initialGeneration := make([]*Individual, b.maxPop)
	for i := range initialGeneration {
		initialGeneration[i] = newMutantIndividual(b.chromosomeLen)
	}
	return initialGeneration
}

func (b BRKGA[T]) newGeneration(orderedGeneration []*Individual) []*Individual {
	newGeneration := make([]*Individual, b.maxPop)

	for i := 0; i < b.maxPop; i++ {
		if i < b.topQnt {
			newGeneration[i] = orderedGeneration[i]
			continue
		}

		if i < b.maxPop-b.mutantQnt {
			individualTop := b.randomIndividualTop(orderedGeneration)
			individualBottom := b.randomIndividualNotTop(orderedGeneration)
			newGeneration[i] = b.crossover(individualTop, individualBottom)
			continue
		}

		newGeneration[i] = newMutantIndividual(b.chromosomeLen)
	}

	return newGeneration
}

func (b BRKGA[T]) randomIndividualTop(orderedGeneration []*Individual) *Individual {
	return orderedGeneration[rand.Intn(b.topQnt)]
}

func (b BRKGA[T]) randomIndividualNotTop(orderedGeneration []*Individual) *Individual {
	return orderedGeneration[rand.Intn(b.maxPop-b.topQnt)+b.topQnt]
}

func (b BRKGA[T]) crossover(individualTop *Individual, individualBottom *Individual) *Individual {
	newIndividual := newIndividual(b.chromosomeLen)
	for i := 0; i < b.chromosomeLen; i++ {
		newIndividual.Chromosomes[i] = b.chooseChromossome(individualTop.Chromosomes[i], individualBottom.Chromosomes[i])
	}

	return newIndividual
}

func (b BRKGA[T]) chooseChromossome(chromosomeTop, chromosomeBottom *Chromossome) *Chromossome {
	if rand.Float64() < b.biasPercentage {
		return chromosomeTop
	}
	return chromosomeBottom
}

func (b BRKGA[T]) evaluateGeneration(generation []*Individual) {
	for _, individual := range generation {
		if individual.Score == 0 {
			individual.Score = b.evaluateIndividual(individual)
		}
	}
}

func (b BRKGA[T]) evaluateIndividual(individual *Individual) float64 {
	decodedIndividual := b.decoder(individual)
	return b.measurer(decodedIndividual)
}

func (BRKGA[T]) orderGeneration(generation []*Individual) {
	sort.Slice(generation, func(i, j int) bool {
		return generation[i].Score > generation[j].Score
	})
}

func calculateQuantity(totalQnt int, percentage float64) int {
	return int(float64(totalQnt) * percentage)
}
