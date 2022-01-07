package interaction_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/zknill/vending/coinage"
	"github.com/zknill/vending/machine"
)

var _ = Describe("Interaction", func() {
	var (
		vendingMachine machine.Machine
	)

	BeforeEach(func() {
		inventory, err := machine.NewInventory(
			machine.Product{
				Coordinate: "A1",
				Price:      49,
				Name:       "Stroop Waffle",
			},

			machine.Product{
				Coordinate: "A2",
				Price:      3,
				Name:       "3 Penny Sweet",
			},
			machine.Product{
				Coordinate: "A3",
				Price:      82,
				Name:       "Multipack Mars",
			},
		)
		Expect(err).To(BeNil())

		inventory.ModifyStock("A1", 1)
		inventory.ModifyStock("A2", 1)
		inventory.ModifyStock("A3", 1)

		d := coinage.NewDenominations(1, 5, 10, 25, 50, 100)

		vendingMachine = machine.New(inventory, coinage.NewHopper(nil), coinage.NewTray(d))
	})

	It("Vends with correct money", func() {
		vendingMachine.InsertCoin(1)
		vendingMachine.InsertCoin(1)
		vendingMachine.InsertCoin(1)

		product, change, err := vendingMachine.Purchase("A2")
		Expect(err).To(BeNil())

		Expect(product.Name).To(Equal("3 Penny Sweet"))
		Expect(change).To(BeEmpty())
	})
})
