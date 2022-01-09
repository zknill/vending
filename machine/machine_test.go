package machine_test

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/zknill/vending/coinage"
	"github.com/zknill/vending/domain"
	"github.com/zknill/vending/machine"
)

var _ = Describe("Machine", func() {
	var (
		vendingMachine machine.Machine
		t              *coinage.Tray
	)

	BeforeEach(func() {
		inventory, err := machine.NewInventory(
			domain.Product{
				Coordinate: "A1",
				Price:      10,
				Name:       "test product",
			},
		)
		Expect(err).To(BeNil())

		inventory.ModifyStock("A1", 1)

		d := coinage.NewDenominations(1, 5, 10, 25, 50, 100)
		t = coinage.NewTray(d)
		h, err := coinage.NewHopper(d, nil)
		Expect(err).NotTo(HaveOccurred())

		vendingMachine = machine.New(inventory, h, t)
	})

	Describe("InsertCoin", func() {
		It("succeeds inserting a coin", func() {
			Expect(vendingMachine.InsertCoin(5)).To(Succeed())
			Expect(t.MeetsPrice(5)).To(BeTrue())
		})

		It("validates coins", func() {
			err := vendingMachine.InsertCoin(13)
			Expect(err).To(HaveOccurred())

			var unknownCoin domain.ErrUnknownCoin
			Expect(errors.As(err, &unknownCoin)).To(BeTrue())
			Expect(unknownCoin.UnknownCoin()).To(BeEquivalentTo(13))
		})
	})

	Describe("Purchase", func() {
		It("purchases", func() {
			vendingMachine.InsertCoin(5)
			vendingMachine.InsertCoin(5)

			p, change, err := vendingMachine.Purchase("A1")
			Expect(err).NotTo(HaveOccurred())
			Expect(change).To(BeZero())

			Expect(p.Name).To(Equal("test product"))
		})

		It("fails on unknown product", func() {
			_, _, err := vendingMachine.Purchase("A2")

			Expect(err).To(HaveOccurred())

			var unknownProduct domain.ErrUnknownProduct
			Expect(errors.As(err, &unknownProduct)).To(BeTrue())
			Expect(unknownProduct.UnknownProduct()).To(Equal("A2"))
		})

		It("fails on out of stock product", func() {
			vendingMachine.Unlock().Restock(map[string]int{
				"A1": -1,
			})

			_, _, err := vendingMachine.Purchase("A1")
			Expect(err).To(HaveOccurred())

			var outOfStock domain.ErrOutOfStock
			Expect(errors.As(err, &outOfStock)).To(BeTrue())
			Expect(outOfStock.OutOfStockProduct()).To(Equal("A1"))
		})

		It("fails on too little money", func() {
			vendingMachine.InsertCoin(5)

			_, _, err := vendingMachine.Purchase("A1")
			Expect(err).To(HaveOccurred())

			var errMoney domain.ErrNotEnoughMoney
			Expect(errors.As(err, &errMoney)).To(BeTrue())
			Expect(errMoney.RemainingRequired()).To(BeEquivalentTo(5))
		})
	})
})
