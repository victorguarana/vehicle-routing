package brkga

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Chromossome", func() {
	Context("when chromossomes have same value", func() {
		It("should not be identical", func() {
			chromossomeValue1 := Chromossome(1.5)
			chromossome1 := &chromossomeValue1
			chromossomeValue2 := Chromossome(1.5)
			chromossome2 := &(chromossomeValue2)

			Expect(chromossome1).NotTo(BeIdenticalTo(chromossome2))
		})
	})
})

var _ = Describe("Individual", func() {
	Describe("newMutantIndividual", func() {
		It("should create mutant with random cromossomes", func() {
			mutant := newMutantIndividual(3)
			Expect(mutant.Chromosomes).To(HaveLen(3))
			Expect(mutant.Chromosomes).NotTo(ContainElement(nil))
			Expect(mutant.Chromosomes).NotTo(ContainElement(0.0))
			Expect(mutant.Score).To(BeZero())
		})
	})
})
