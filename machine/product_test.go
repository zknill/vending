package machine_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/zknill/vending/domain"
	"github.com/zknill/vending/machine"
)

var _ = Describe("Product", func() {

	Describe("Inventory", func() {
		var inventory *machine.Inventory

		BeforeEach(func() {
			var err error
			inventory, err = machine.NewInventory(
				domain.Product{
					Coordinate: "A1",
					Price:      250,
					Name:       "Stroop Waffle",
				},

				domain.Product{
					Coordinate: "A2",
					Price:      3,
					Name:       "3 Penny Sweet",
				},
				domain.Product{
					Coordinate: "A3",
					Price:      1000,
					Name:       "Multipack Mars",
				},
			)
			Expect(err).To(BeNil())
		})

		It("adds stock", func() {
			Expect(inventory.ModifyStock("A1", 1)).To(BeNil())

			level, found := inventory.StockLevel("A1")
			Expect(found).To(BeTrue())
			Expect(level).To(Equal(1))
		})

		It("handles unknown stock", func() {
			Expect(inventory.ModifyStock("B1", 1)).NotTo(BeNil())

			level, found := inventory.StockLevel("B1")
			Expect(found).To(BeFalse())
			Expect(level).To(BeZero())
		})

		It("handles out of stock", func() {
			Expect(inventory.ModifyStock("A3", 1)).To(BeNil())

			Expect(inventory.ModifyStock("A3", -2)).NotTo(BeNil())

			level, found := inventory.StockLevel("A3")
			Expect(found).To(BeTrue())
			Expect(level).To(Equal(1))
		})

		It("checks unique coordinates for product catalog", func() {
			_, err := machine.NewInventory(
				domain.Product{
					Coordinate: "A1",
					Price:      250,
					Name:       "Stroop Waffle",
				},

				domain.Product{
					Coordinate: "A1",
					Price:      3,
					Name:       "3 Penny Sweet",
				},
			)
			Expect(err).NotTo(BeNil())
		})
	})
})
