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
// Deposit will return true if the purchase has been made and
// empty the tray. Deposit will return false if the purchase
// has not been made, e.g. not enough money, or change cannot
// be given for the purchase.
func (h *Hopper) Deposit(t *Tray, price int) ([]uint, bool) {
	var (
		usedTrayCoins   = make([]bool, len(t.coins))
		usedHopperCoins = make([]bool, len(h.coins))
	)

	solution := resolveCoins(int(price), t.coins, usedTrayCoins, h.coins, usedHopperCoins)

	if solution.solvedPrice > 0 {
		return nil, false
	}

	var hopper []uint
	var change []uint

	for i, used := range solution.usedTrayCoins {
		if used {
			hopper = append(hopper, t.coins[i])
		} else {
			change = append(change, t.coins[i])
		}
	}

	for i, used := range solution.usedHopperCoins {
		if used {
			change = append(change, h.coins[i])
		} else {
			hopper = append(hopper, h.coins[i])
		}
	}

	t.coins = make([]uint, 0)
	h.coins = hopper

	return change, true
}

type solution struct {
	usedTrayCoins   []bool
	usedHopperCoins []bool
	solvedPrice     int
}

// This method decides which solution is better.
// The optimal solution has a solvedPrice of zero,
// which means the exact money and change could be
// found. The next best solution is an overpayment
// but minimising the amount overpaid.
// This method could be extended with biz logic like
// "best solution has smallest number of coins in change", etc.
func (current solution) isBetter(challenger solution) solution {
	if challenger.solvedPrice <= 0 {
		if current.solvedPrice > 0 {
			return challenger
		}

		if challenger.solvedPrice > current.solvedPrice {
			return challenger
		}
	}

	return current
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
) solution {

	bestSolution := solution{
		usedTrayCoins:   usedTrayCoins,
		usedHopperCoins: usedHopperCoins,
		solvedPrice:     price,
	}

	if price == 0 {
		return bestSolution
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

			solution := resolveCoins(
				remainingPrice,
				trayCoins,
				updatedUsed,
				hopperCoins,
				usedHopperCoins,
			)

			if solution.solvedPrice == 0 {
				return solution
			}

			bestSolution = bestSolution.isBetter(solution)
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

			solution := resolveCoins(
				remainingPrice,
				trayCoins,
				usedTrayCoins,
				hopperCoins,
				updatedUsed,
			)

			if solution.solvedPrice == 0 {
				return solution
			}

			if solution.solvedPrice > 0 {
				continue
			}

			bestSolution = bestSolution.isBetter(solution)
		}
	}

	return bestSolution
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
