package machine

import (
	"fmt"

	"github.com/zknill/vending/domain"
)

type errUnknownCoin struct {
	coin uint
}

func (e errUnknownCoin) UnknownCoin() uint { return e.coin }
func (e errUnknownCoin) Error() string     { return fmt.Sprintf("unknown coin: %d", e.coin) }

var _ domain.ErrUnknownCoin = (*errUnknownCoin)(nil)

type errUnknownProduct struct {
	coordinate string
}

func (e errUnknownProduct) UnknownProduct() string { return e.coordinate }
func (e errUnknownProduct) Error() string          { return fmt.Sprintf("unknown product: %q", e.coordinate) }

var _ domain.ErrUnknownProduct = (*errUnknownProduct)(nil)

type errNotEnoughMoney struct {
	required int
}

func (e errNotEnoughMoney) Error() string          { return fmt.Sprintf("requires more money: %d", e.required) }
func (e errNotEnoughMoney) RemainingRequired() int { return e.required }

var _ domain.ErrNotEnoughMoney = (*errNotEnoughMoney)(nil)

type errExactChange struct{}

func (e errExactChange) Error() string        { return "use exact change" }
func (e errExactChange) UseExactChange() bool { return true }

var _ domain.ErrExactChange = (*errExactChange)(nil)

type errOutOfStock struct {
	coordinate string
}

func (e errOutOfStock) Error() string             { return fmt.Sprintf("out of stock product: %q", e.coordinate) }
func (e errOutOfStock) OutOfStockProduct() string { return e.coordinate }
