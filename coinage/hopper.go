package coinage

import "fmt"

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
// Deposit will return true if the purchase has been made and
// empty the tray. Deposit will return false if the purchase
// has not been made, e.g. not enough money, or change cannot
// be given for the purchase.
func (h *Hopper) Deposit(t *Tray, price uint) ([]uint, bool) {
	usedTrayCoins := make([]bool, len(t.coins))
	usedHopperCoins := make([]bool, len(h.coins))

	tSolution, hSolution, solved := resolveCoins(int(price), t.coins, usedTrayCoins, h.coins, usedHopperCoins)
	if !solved {
		return nil, false
	}

	var hopper []uint
	var change []uint

	for i, used := range tSolution {
		if used {
			hopper = append(hopper, t.coins[i])
		} else {
			change = append(change, t.coins[i])
		}
	}

	for i, used := range hSolution {
		if used {
			change = append(change, h.coins[i])
		} else {
			hopper = append(hopper, h.coins[i])
		}
	}

	fmt.Printf("hopper: %v", hopper)
	fmt.Printf("change: %v", change)

	t.coins = make([]uint, 0)
	h.coins = hopper

	return change, true
}

// We have coins in the tray
// we have coins in the hopper
// we want to pay for the item, and return change.
// ideally this nets-out to zero.
// +ve number is overpay (currently, not acceptable)
// -ve number is over-change (underpay, not acceptable)
// 0 number is net-out

// search tree problem
// what set of coins on left (tray) and right (hopper)
// can be used to net out to zero

// find the coins that meet or exceed the price
//  - if meets exactly --> DONE
//  - if overpay, find the coins from the hopper that we can use to net out.
// best solution is one that minimises overpayment
// (config for if overpayment is allowed)

// returns []usedTrayCoins, []usedHopperCoinsForChange, matched
func resolveCoins(
	price int,
	trayCoins []uint,
	usedTrayCoins []bool,
	hopperCoins []uint,
	usedHopperCoins []bool,
) ([]bool, []bool, bool) {

	fmt.Printf("---\nprice: %d\ntray: %v\nhopper: %v\n", price, usedTrayCoins, usedHopperCoins)

	if price == 0 {
		return usedTrayCoins, usedHopperCoins, true
	}

	if price > 0 {
		// not met
		// try and make price from tray

		for i := range usedTrayCoins {
			if usedTrayCoins[i] {
				continue
			}

			coin := trayCoins[i]
			remainingPrice := price - int(coin)

			updatedUsed := markIndex(usedTrayCoins, i)

			traySolution, hopperSolution, solved := resolveCoins(
				remainingPrice,
				trayCoins,
				updatedUsed,
				hopperCoins,
				usedHopperCoins,
			)

			if solved {
				return traySolution, hopperSolution, true
			}
		}
	}

	if price < 0 {
		// currently overpaying
		// try and give change from hopper

		for i := range usedHopperCoins {
			if usedHopperCoins[i] {
				continue
			}

			coin := hopperCoins[i]
			remainingPrice := price + int(coin)

			updatedUsed := markIndex(usedHopperCoins, i)

			traySolution, hopperSolution, solved := resolveCoins(
				remainingPrice,
				trayCoins,
				usedTrayCoins,
				hopperCoins,
				updatedUsed,
			)

			if solved {
				return traySolution, hopperSolution, true
			}
		}
	}

	return nil, nil, false
}

func markIndex(in []bool, idx int) []bool {
	out := make([]bool, len(in))
	copy(out, in)

	out[idx] = true

	return out
}

func (h *Hopper) Value() uint {
	return sumCoins(h.coins)
}
