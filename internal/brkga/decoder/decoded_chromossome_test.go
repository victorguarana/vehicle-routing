package decoder

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/victorguarana/vehicle-routing/internal/brkga"
)

var _ = Describe("DecodedChromossome", func() {

	Describe("orderDecodedChromossomesByChromossome", func() {
		var decodedChromossome1 *decodedChromossome
		var decodedChromossome2 *decodedChromossome
		var decodedChromossome3 *decodedChromossome
		var decodedChromossome4 *decodedChromossome

		BeforeEach(func() {
			c1 := brkga.Chromossome(0.1)
			c2 := brkga.Chromossome(0.2)
			c3 := brkga.Chromossome(0.3)
			c4 := brkga.Chromossome(0.4)

			decodedChromossome1 = &decodedChromossome{
				chromossome: &c1,
			}
			decodedChromossome2 = &decodedChromossome{
				chromossome: &c2,
			}
			decodedChromossome3 = &decodedChromossome{
				chromossome: &c3,
			}
			decodedChromossome4 = &decodedChromossome{
				chromossome: &c4,
			}
		})

		It("should return ordered decoded chromossome list", func() {
			decodedChromossomeList := []*decodedChromossome{
				decodedChromossome2, decodedChromossome4, decodedChromossome3, decodedChromossome1,
			}

			expectedDecodedChromossomeList := []*decodedChromossome{
				decodedChromossome1, decodedChromossome2, decodedChromossome3, decodedChromossome4,
			}

			receivedDecodedChromossomeList := orderDecodedChromossomesByChromossome(decodedChromossomeList)

			Expect(receivedDecodedChromossomeList).To(HaveExactElements(expectedDecodedChromossomeList))
		})
	})
})
