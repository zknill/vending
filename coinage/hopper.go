package coinage

// Hopper contains the money inside the machine
// it controls deposits.
type Hopper struct {
	coins []uint
}

func NewHopper(changeFloat []uint) *Hopper {
	h := &Hopper{coins: make([]uint, 0)}

	if len(changeFloat) > 0 {
		h.coins = append(h.coins, changeFloat...)
	}

	return h
}

// Deposit the coins from the input tray into the hopper.
// Coins are checked against the price, and change is returned.
// Only exact coins are currently accepted, and change will
// never be given. Deposit will return true if the purchase
// has been made and empty the tray. Deposit will return false
// if the purchase has not been made, e.g. not enough money, or
// change cannot be given for the purchase.
func (h *Hopper) Deposit(t *Tray, price uint) ([]uint, bool) {
	// Exact coins for
	if price == sumCoins(t.coins) {
		h.coins = append(h.coins, t.coins...)
		t.Reset()

		return nil, true
	}

	return nil, false
}

func (h *Hopper) Value() uint {
	return sumCoins(h.coins)
}
