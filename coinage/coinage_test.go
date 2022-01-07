package coinage_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/zknill/vending/coinage"
)

var _ = Describe("Coinage", func() {
	d := coinage.NewDenominations(1, 5, 10, 25, 50, 100)

	Describe("Denominations", func() {
		DescribeTable("Valid coins for demoinations", func(d coinage.Denominations, coin int, valid bool) {
			Expect(d.Valid(uint(coin))).To(Equal(valid))
		},
			Entry("When coin is 0", d, 0, false),
			Entry("When coin is 1", d, 1, true),
			Entry("When coin is 2", d, 2, false),
			Entry("When coin is 5", d, 5, true),
			Entry("When coin is 10", d, 10, true),
			Entry("When coin is 25", d, 25, true),
			Entry("When coin is 50", d, 50, true),
			Entry("When coin is 100", d, 100, true),
			Entry("When coin is 101", d, 101, false),
		)

		It("succeeds on mulitple coins all valid", func() {
			Expect(d.Valid(1, 5, 10, 25, 50, 100)).To(BeTrue())
		})

		It("fails on mulitple coins one invalid", func() {
			Expect(d.Valid(1, 5, 10, 25, 21, 50, 100)).To(BeFalse())
		})

		It("fails on mulitple coins all invalid", func() {
			Expect(d.Valid(11, 15, 110, 125, 150, 1100)).To(BeFalse())
		})
	})

	Describe("Tray", func() {
		t := coinage.NewTray(d)

		BeforeEach(func() {
			t.Reset()
		})

		It("succeeds paying with exact money", func() {
			Expect(t.Insert(1, 5, 50)).To(BeTrue())
			Expect(t.MeetsPrice(56)).To(BeTrue())
		})

		It("succeeds paying with more money", func() {
			Expect(t.Insert(1, 5, 50)).To(BeTrue())
			Expect(t.MeetsPrice(56)).To(BeTrue())
		})

		It("fails paying with too little money", func() {
			t.Insert(5, 100)
			Expect(t.MeetsPrice(56)).To(BeTrue())
		})

		It("ejects money", func() {
			Expect(t.Insert(1, 5, 50)).To(BeTrue())
			Expect(t.MeetsPrice(56)).To(BeTrue())

			coins := t.Reset()
			Expect(t.MeetsPrice(1)).To(BeFalse())
			Expect(t.MeetsPrice(0)).To(BeTrue())
			Expect(coins).To(BeEquivalentTo([]uint{1, 5, 50}))
		})

		It("rejects invalid coins", func() {
			Expect(t.Insert(21)).To(BeFalse())
		})
	})
})
