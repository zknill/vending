package coinage_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/zknill/vending/coinage"
)

var _ = Describe("Hopper", func() {
	d := coinage.NewDenominations(1, 5, 10, 25, 50, 100)

	Context("Depositing correct money", func() {
		tray := coinage.NewTray(d)
		var hopper *coinage.Hopper

		BeforeEach(func() {
			tray.Reset()

			hopper = coinage.NewHopper(nil)
		})

		It("accepts setup change float", func() {
			hopper = coinage.NewHopper([]uint{1, 5, 10, 25, 50, 100})

			Expect(hopper.Value()).To(Equal(uint(191)))

		})

		It("succeeds with exact money", func() {
			Expect(hopper.Value()).To(BeZero())

			tray.Insert(25, 1, 5)

			change, ok := hopper.Deposit(tray, 31)

			Expect(ok).To(BeTrue())
			Expect(change).To(BeEmpty())

			Expect(tray.MeetsPrice(1)).To(BeFalse())
			Expect(tray.MeetsPrice(0)).To(BeTrue())

			Expect(hopper.Value()).To(Equal(uint(31)))
		})
	})
})
