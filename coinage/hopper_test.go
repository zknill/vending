package coinage_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/zknill/vending/coinage"
)

var _ = Describe("Hopper", func() {
	d := coinage.NewDenominations(1, 5, 10, 25, 50, 100)

	Context("Depositing correct money", func() {
		var (
			hopper *coinage.Hopper
			tray   *coinage.Tray
		)

		BeforeEach(func() {
			tray = coinage.NewTray(d)
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

		It("succeeds with extra money", func() {
			Expect(hopper.Value()).To(BeZero())

			tray.Insert(1, 1, 25, 5)

			change, ok := hopper.Deposit(tray, 31)

			Expect(ok).To(BeTrue())
			Expect(change).To(BeEquivalentTo([]uint{1}))

			Expect(tray.MeetsPrice(1)).To(BeFalse())
			Expect(tray.MeetsPrice(0)).To(BeTrue())

			Expect(hopper.Value()).To(Equal(uint(31)))
		})

		It("succeeds with change from hopper", func() {
			float := make([]uint, 99)
			expectedChange := make([]uint, 99)
			for i := range float {
				float[i] = 1
				expectedChange[i] = 1
			}
			hopper = coinage.NewHopper(float)

			tray.Insert(100)

			change, ok := hopper.Deposit(tray, 1)

			Expect(ok).To(BeTrue())
			Expect(change).To(BeEquivalentTo(expectedChange))

			Expect(tray.MeetsPrice(1)).To(BeFalse())
			Expect(tray.MeetsPrice(0)).To(BeTrue())

			Expect(hopper.Value()).To(Equal(uint(100)))
		})
	})

	Context("Overpaying", func() {
		var (
			hopper *coinage.Hopper
			tray   *coinage.Tray
		)

		BeforeEach(func() {
			tray = coinage.NewTray(d)
			hopper = coinage.NewHopper(nil)
		})

		It("succeeds when change cannot be given", func() {
			tray.Insert(25)

			change, ok := hopper.Deposit(tray, 24)
			Expect(ok).To(BeTrue())
			Expect(change).To(BeEmpty())
		})

		It("overcharges by smallest amount when change cannot be given", func() {
			tray.Insert(100)
			tray.Insert(25)

			change, ok := hopper.Deposit(tray, 24)
			Expect(ok).To(BeTrue())
			Expect(change).To(BeEquivalentTo([]uint{100}))
		})
	})
})
