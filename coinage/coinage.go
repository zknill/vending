package coinage

// Denominations holds the valid coins
// that can be accepted.
type Denominations struct {
	// 1, 5, 10, 25, 50, 100
	coins map[uint]bool
}

func NewDenominations(coins ...uint) Denominations {
	allowed := make(map[uint]bool)

	for _, coin := range coins {
		allowed[coin] = true
	}

	return Denominations{coins: allowed}
}

// Valid returns true if all coins are valid.
// False if any single coin is invalid.
func (d Denominations) Valid(coins ...uint) bool {
	for _, c := range coins {
		if !d.coins[c] {
			return false
		}
	}

	return true
}

// Tray is the coin input tray. Coins are held
// in the tray until a purchase is made. Coins
// should be deposited in the hopper afterwards.
type Tray struct {
	demoinations Denominations
	coins        []uint
}

func NewTray(d Denominations) *Tray {
	return &Tray{demoinations: d}
}

// Insert a coin into the tray
func (t *Tray) Insert(coin ...uint) bool {
	if !t.demoinations.Valid(coin...) {
		return false
	}

	t.coins = append(t.coins, coin...)

	return true
}

// Reset ejects the coins from the tray
func (t *Tray) Reset() []uint {
	c := t.coins
	t.coins = make([]uint, 0)

	return c
}

// Check if there are enough coins to meet
// the price of the purchase.
func (t Tray) MeetsPrice(price int) bool {
	return sumCoins(t.coins) >= uint(price)
}

func sumCoins(coins []uint) uint {
	var total uint

	for _, c := range coins {
		total += c
	}

	return total
}
