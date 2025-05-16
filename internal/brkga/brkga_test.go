package brkga

import (
	"errors"
	"math"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	gomock "go.uber.org/mock/gomock"
)

var _ = Describe("BRKGA", func() {
	var mockedCtrl *gomock.Controller
	var mockedDecoder *MockIDecoder[int]
	var mockedMeasurer *MockIMeasurer[int]
	var params BRKGAParams[int]
	var sut BRKGA[int]

	BeforeEach(func() {
		mockedCtrl = gomock.NewController(GinkgoT())
		mockedDecoder = NewMockIDecoder[int](mockedCtrl)
		mockedMeasurer = NewMockIMeasurer[int](mockedCtrl)

		params = BRKGAParams[int]{
			MaxPop:              10,
			TopPercentage:       0.2,
			CrossoverPercentage: 0.6,
			BiasPercentage:      0.7,
			ChromosomeLen:       5,
			GenerationLimit:     5,
			Decoder:             mockedDecoder,
			Measurer:            mockedMeasurer,
		}

		sut = NewBRKGA(params)
	})

	Describe("NewBRKGA", func() {
		It("should create brkga instance with correct atributes", func() {
			receivedSut := NewBRKGA(BRKGAParams[int]{
				MaxPop:              10,
				TopPercentage:       0.3,
				CrossoverPercentage: 0.5,
				BiasPercentage:      0.7,
				ChromosomeLen:       5,
				GenerationLimit:     20,
				Decoder:             mockedDecoder,
				Measurer:            mockedMeasurer,
			})

			expectedSut := BRKGA[int]{
				chromosomeLen:   5,
				maxPop:          10,
				topQnt:          3,
				mutantQnt:       2,
				biasPercentage:  0.7,
				generationLimit: 20,
				decoder:         mockedDecoder,
				measurer:        mockedMeasurer,
			}

			Expect(receivedSut.chromosomeLen).To(Equal(expectedSut.chromosomeLen))
			Expect(receivedSut.maxPop).To(Equal(expectedSut.maxPop))
			Expect(receivedSut.topQnt).To(Equal(expectedSut.topQnt))
			Expect(receivedSut.mutantQnt).To(Equal(expectedSut.mutantQnt))
			Expect(receivedSut.biasPercentage).To(Equal(expectedSut.biasPercentage))
			Expect(receivedSut.generationLimit).To(Equal(expectedSut.generationLimit))
			Expect(receivedSut.measurer).To(BeAssignableToTypeOf(expectedSut.measurer))
			Expect(receivedSut.decoder).To(BeAssignableToTypeOf(expectedSut.decoder))
		})
	})

	Describe("createInitialGeneration", func() {
		It("should create new generation with correct values", func() {
			sut.chromosomeLen = 10
			receivedGeneration := sut.createInitialGeneration()
			Expect(receivedGeneration).To(HaveLen(10))
			Expect(receivedGeneration).NotTo(ContainElements(nil))
			Expect(receivedGeneration).NotTo(ContainElements(&Individual{}))
		})

	})

	Describe("newGeneration", func() {
		var individual1 = newMutantIndividual(1)
		var individual2 = newMutantIndividual(1)
		var individual3 = newMutantIndividual(1)
		var individual4 = newMutantIndividual(1)
		var individual5 = newMutantIndividual(1)
		var generation = []*Individual{individual1, individual2, individual3, individual4, individual5}

		BeforeEach(func() {
			sut.topQnt = 2
			sut.maxPop = 5
			sut.mutantQnt = 1
			sut.chromosomeLen = 1
			sut.biasPercentage = 0.5
		})

		It("should select only top individuals to next generation", func() {
			receivedGeneration := sut.newGeneration(generation)

			Expect(receivedGeneration[0]).To(BeIdenticalTo(individual1))
			Expect(receivedGeneration[1]).To(BeIdenticalTo(individual2))
		})

		It("should eventually crossover from top and non top individuals", func() {
			// Received individuals should be equal (same single chromossome) but not identical
			newIndividual2Func := func() *Individual { return sut.newGeneration(generation)[2] }
			Expect(newIndividual2Func()).ToNot(
				SatisfyAny(BeIdenticalTo(individual1), BeIdenticalTo(individual2), BeIdenticalTo(individual3), BeIdenticalTo(individual4), BeIdenticalTo(individual5)))
			Eventually(newIndividual2Func).Should(Equal(individual1))
			Eventually(newIndividual2Func).Should(Equal(individual2))
			Eventually(newIndividual2Func).Should(Equal(individual3))
			Eventually(newIndividual2Func).Should(Equal(individual4))
			Eventually(newIndividual2Func).Should(Equal(individual5))

			newIndividual3Func := func() *Individual { return sut.newGeneration(generation)[3] }
			Expect(newIndividual3Func()).ToNot(
				SatisfyAny(BeIdenticalTo(individual1), BeIdenticalTo(individual2), BeIdenticalTo(individual3), BeIdenticalTo(individual4), BeIdenticalTo(individual5)))
			Eventually(newIndividual3Func).Should(Equal(individual1))
			Eventually(newIndividual3Func).Should(Equal(individual2))
			Eventually(newIndividual3Func).Should(Equal(individual3))
			Eventually(newIndividual3Func).Should(Equal(individual4))
			Eventually(newIndividual3Func).Should(Equal(individual5))
		})

		It("should create mutant to next generation", func() {
			receivedGeneration := sut.newGeneration(generation)

			Expect(receivedGeneration[4]).NotTo(
				BeElementOf(individual1, individual2, individual3, individual4, individual5))
		})
	})

	Describe("randomIndividualTop", func() {
		var individual1 = newMutantIndividual(1)
		var individual2 = newMutantIndividual(1)
		var individual3 = newMutantIndividual(1)
		var individual4 = newMutantIndividual(1)
		var individual5 = newMutantIndividual(1)
		var generation = []*Individual{individual1, individual2, individual3, individual4, individual5}

		It("returns individual from top", func() {
			sut.topQnt = 2
			receivedIndividual := sut.randomIndividualTop(generation)
			Eventually(receivedIndividual).MustPassRepeatedly(10).Should(BeElementOf(individual1, individual2))
		})
	})

	Describe("randomIndividualNotTop", func() {
		var individual1 = newMutantIndividual(1)
		var individual2 = newMutantIndividual(1)
		var individual3 = newMutantIndividual(1)
		var individual4 = newMutantIndividual(1)
		var individual5 = newMutantIndividual(1)
		var generation = []*Individual{individual1, individual2, individual3, individual4, individual5}

		It("returns individual from bottom", func() {
			sut.topQnt = 2
			sut.maxPop = 5
			receivedIndividual := sut.randomIndividualNotTop(generation)
			Eventually(receivedIndividual).MustPassRepeatedly(10).Should(BeElementOf(individual3, individual4, individual5))
		})
	})

	Describe("crossover", func() {
		var chromossome1 = Chromossome(0.1)
		var chromossome2 = Chromossome(0.2)
		var chromossome3 = Chromossome(0.3)
		var chromossome4 = Chromossome(0.4)
		var chromossome5 = Chromossome(0.5)
		var chromossome6 = Chromossome(0.6)

		var individualTop = &Individual{
			Chromosomes: []*Chromossome{&chromossome1, &chromossome3, &chromossome5},
			Score:       10,
		}
		var individualBottom = &Individual{
			Chromosomes: []*Chromossome{&chromossome2, &chromossome4, &chromossome6},
			Score:       5,
		}

		Context("when bias percentage is 100%", func() {
			It("should only choose top chromossomes", func() {
				sut.biasPercentage = 1
				sut.chromosomeLen = 3
				receivedIndividual := sut.crossover(individualTop, individualBottom)
				expectedIndividual := &Individual{
					Chromosomes: []*Chromossome{&chromossome1, &chromossome3, &chromossome5},
					Score:       0,
				}
				Expect(receivedIndividual).To(Equal(expectedIndividual))

			})
		})

		Context("when bias percentage is 0%", func() {
			It("should only choose bottom chromossomes", func() {
				sut.biasPercentage = 0
				sut.chromosomeLen = 3
				receivedIndividual := sut.crossover(individualTop, individualBottom)
				expectedIndividual := &Individual{
					Chromosomes: []*Chromossome{&chromossome2, &chromossome4, &chromossome6},
					Score:       0,
				}
				Expect(receivedIndividual).To(Equal(expectedIndividual))

			})
		})

		It("should choose correct chromossome based on bias percentage", func() {
			chromossomeTop := Chromossome(0.1)
			chromossomeBottom := Chromossome(0.2)

			sut.biasPercentage = 0
			Expect(sut.chooseChromossome(&chromossomeTop, &chromossomeBottom)).To(Equal(&chromossomeBottom))

			sut.biasPercentage = 1
			Expect(sut.chooseChromossome(&chromossomeTop, &chromossomeBottom)).To(Equal(&chromossomeTop))
		})
	})

	Describe("evaluateGeneration", func() {
		var individual1 *Individual
		var individual2 *Individual
		var individual3 *Individual
		var initialGeneration []*Individual

		BeforeEach(func() {
			individual1 = newMutantIndividual(1)
			individual2 = newMutantIndividual(1)
			individual3 = newMutantIndividual(1)
			individual2.Score = 20.0
			initialGeneration = []*Individual{individual1, individual2, individual3}
		})

		Context("when optimization goal is set to maximize", func() {
			It("should evaluate and fill score from not evaluated individuals", func() {
				sut.optimizationGoal = Maximize
				mockedDecoder.EXPECT().Decode(individual1).Return(1, nil)
				mockedDecoder.EXPECT().Decode(individual3).Return(0, errors.New("mocked error"))
				mockedMeasurer.EXPECT().Measure(1).Return(10.0)

				sut.evaluateGeneration(initialGeneration)

				Expect(initialGeneration[0].Score).To(Equal(10.0))
				Expect(initialGeneration[1].Score).To(Equal(20.0))
				Expect(initialGeneration[2].Score).To(Equal(math.Inf(-1)))
			})
		})

		Context("when optimization goal is set to minimize", func() {
			It("should evaluate and fill score from not evaluated individuals", func() {
				sut.optimizationGoal = Minimize
				mockedDecoder.EXPECT().Decode(individual1).Return(1, nil)
				mockedDecoder.EXPECT().Decode(individual3).Return(0, errors.New("mocked error"))
				mockedMeasurer.EXPECT().Measure(1).Return(10.0)

				sut.evaluateGeneration(initialGeneration)

				Expect(initialGeneration[0].Score).To(Equal(10.0))
				Expect(initialGeneration[1].Score).To(Equal(20.0))
				Expect(initialGeneration[2].Score).To(Equal(math.Inf(1)))
			})
		})
	})

	Describe("orderGeneration", func() {
		var individualWithScore1 = &Individual{Score: 1}
		var individualWithScore2 = &Individual{Score: 2}
		var individualWithScore3 = &Individual{Score: 3}
		var individualWithScore4 = &Individual{Score: 4}
		var individualWithScore1_1 = &Individual{Score: 1}

		Context("when generation already ordered", func() {
			It("should do nothing", func() {
				sut.optimizationGoal = Maximize
				expectedGeneration := []*Individual{
					individualWithScore4, individualWithScore3, individualWithScore2, individualWithScore1, individualWithScore1_1,
				}

				generation := []*Individual{
					individualWithScore4, individualWithScore3, individualWithScore2, individualWithScore1, individualWithScore1_1,
				}

				sut.orderGeneration(generation)
				Expect(generation).To(Equal(expectedGeneration))
			})
		})

		Context("when generation is not ordered", func() {
			Context("when optimizer is maximize", func() {
				It("should order generation desc by score", func() {
					sut.optimizationGoal = Maximize
					expectedGeneration := []*Individual{
						individualWithScore4, individualWithScore3, individualWithScore2, individualWithScore1, individualWithScore1_1,
					}

					generation := []*Individual{
						individualWithScore1_1, individualWithScore4, individualWithScore2, individualWithScore1, individualWithScore3,
					}

					sut.orderGeneration(generation)
					Expect(generation).To(Equal(expectedGeneration))
				})
			})

			Context("when optimizer is minimize", func() {
				It("should order generation asc by score", func() {
					sut.optimizationGoal = Minimize
					expectedGeneration := []*Individual{
						individualWithScore1, individualWithScore1_1, individualWithScore2, individualWithScore3, individualWithScore4,
					}

					generation := []*Individual{
						individualWithScore1_1, individualWithScore4, individualWithScore2, individualWithScore1, individualWithScore3,
					}

					sut.orderGeneration(generation)
					Expect(generation).To(Equal(expectedGeneration))
				})
			})
		})
	})

	Describe("calculateQuantity", func() {
		It("should calculate quantities", func() {
			Expect(calculateQuantity(20, 0.5)).To(Equal(10))
			Expect(calculateQuantity(20, 0.3)).To(Equal(6))
		})

	})

	Describe("defineBestScore", func() {
		var sut BRKGA[int]

		Context("when optimization goal is Maximize", func() {
			BeforeEach(func() {
				sut.optimizationGoal = Maximize
			})

			It("should update best score if current score is higher", func() {
				bestScore := 10.0
				currentScore := 15.0
				generationCounter := 1

				receivedBestScore := sut.defineBestScore(bestScore, currentScore, generationCounter)
				Expect(receivedBestScore).To(Equal(currentScore))
			})

			It("should not update best score if current score is lower", func() {
				bestScore := 10.0
				currentScore := 5.0
				generationCounter := 1

				receivedBestScore := sut.defineBestScore(bestScore, currentScore, generationCounter)
				Expect(receivedBestScore).To(Equal(bestScore))
			})
		})

		Context("when optimization goal is Minimize", func() {
			BeforeEach(func() {
				sut.optimizationGoal = Minimize
			})

			It("should update best score if current score is lower", func() {
				bestScore := 10.0
				currentScore := 5.0
				generationCounter := 1

				receivedBestScore := sut.defineBestScore(bestScore, currentScore, generationCounter)
				Expect(receivedBestScore).To(Equal(currentScore))
			})

			It("should not update best score if current score is higher", func() {
				bestScore := 10.0
				currentScore := 15.0
				generationCounter := 1

				receivedBestScore := sut.defineBestScore(bestScore, currentScore, generationCounter)
				Expect(receivedBestScore).To(Equal(bestScore))
			})
		})
	})
})
