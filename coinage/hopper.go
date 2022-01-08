package coinage

import (
	"sort"
	"strconv"
	"strings"
)

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
// has not been made, e.g. not enough money.
// Deposit is designed to maximise the purchase opportunities.
// If the exact change cannot be given, the purchase will succeed,
// and we return the most possible money to the customer (smallest
// overcharge possible).
func (h *Hopper) Deposit(t *Tray, price int) ([]uint, bool) {
	var (
		usedTrayCoins   = make([]bool, len(t.coins))
		usedHopperCoins = make([]bool, len(h.coins))
	)

	solution := resolveCoins(make(map[memoKey]solution), int(price), t.coins, usedTrayCoins, h.coins, usedHopperCoins)

	if solution.solvedPrice > 0 {
		// not enough money for the purchase
		return nil, false
	}

	var hopper []uint
	var change []uint

	// collect the coins used, and the change to be given.

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

// We use memoisation to ensure we don't solve the same sub-problems
// more than once. We memoise based on the set of coins used, and
// we store the best solution for that set of coins. We are using a
// pair of strings for the memoisation map key, because slices don't
// have equality defined. I don't really love this, and could be improved.
type memoKey struct {
	usedTrayCoins   string
	usedHopperCoins string
}

func newKey(usedTray, usedHopper []bool, tray, hopper []uint) memoKey {
	return memoKey{
		usedTrayCoins:   stringify(usedTray, tray),
		usedHopperCoins: stringify(usedHopper, hopper),
	}
}

func stringify(used []bool, coins []uint) string {
	var c []uint

	for i := range used {
		if used[i] {
			c = append(c, coins[i])
		}
	}

	sort.Slice(c, func(i, j int) bool { return c[i] < c[j] })

	sb := &strings.Builder{}
	for _, cc := range c {
		sb.WriteString(strconv.FormatUint(uint64(cc), 10))
	}

	return sb.String()
}

// This method works out the "best" set of coins to use from the hopper and tray.
// Best is defined by the 'solution' struct.
// We want to pay for the item and return the change. We have a price to resolve.
// At the start the price is the price of the product we want to buy. We then
// start to select coins, the coins we select affect the remaining price to solve.
// A positive price means that we have not yet met the total cost of the product.
// A negative price means that we are overpaying for the product.
// A price of zero means that we have met the price of the product exactly, and
// possibly have the correct change to give.
func resolveCoins(
	memomap map[memoKey]solution,
	price int,
	trayCoins []uint,
	usedTrayCoins []bool,
	hopperCoins []uint,
	usedHopperCoins []bool,
) solution {

	// create the best solution we know so far.
	bestSolution := solution{
		usedTrayCoins:   usedTrayCoins,
		usedHopperCoins: usedHopperCoins,
		solvedPrice:     price,
	}

	k := newKey(usedTrayCoins, usedHopperCoins, trayCoins, hopperCoins)

	if price == 0 {
		// price is met exactly, break
		memomap[k] = bestSolution
		return bestSolution
	}

	// check if this branch of the search tree
	// has already been solved. If yes, return
	// the cached solution.
	if cached, ok := memomap[k]; ok {
		return cached
	}

	if price > 0 {
		// price of the item is not yet met
		// keep using tray coins to meet the price

		// pick each coin in turn, and see if we can
		// solve the price using that coin.
		for i := range usedTrayCoins {
			if usedTrayCoins[i] {
				continue
			}

			coin := trayCoins[i]
			remainingPrice := price - int(coin)

			updatedUsed := markIndex(usedTrayCoins, i)

			solution := resolveCoins(
				memomap,
				remainingPrice,
				trayCoins,
				updatedUsed,
				hopperCoins,
				usedHopperCoins,
			)

			if solution.solvedPrice == 0 {
				memomap[newKey(usedTrayCoins, usedHopperCoins, trayCoins, hopperCoins)] = solution
				return solution
			}

			bestSolution = bestSolution.isBetter(solution)
		}
	}

	if price < 0 {
		// currently the customer is paying too much
		// try and give change from the hopper.

		// pick each coin in turn, and see if we can
		// solve the price using that coin.
		for i := range usedHopperCoins {
			if usedHopperCoins[i] {
				continue
			}

			coin := hopperCoins[i]
			remainingPrice := price + int(coin)

			updatedUsed := markIndex(usedHopperCoins, i)

			solution := resolveCoins(
				memomap,
				remainingPrice,
				trayCoins,
				usedTrayCoins,
				hopperCoins,
				updatedUsed,
			)

			if solution.solvedPrice == 0 {
				memomap[newKey(usedTrayCoins, usedHopperCoins, trayCoins, hopperCoins)] = solution
				return solution
			}

			if solution.solvedPrice > 0 {
				continue
			}

			bestSolution = bestSolution.isBetter(solution)
		}
	}

	memomap[newKey(usedTrayCoins, usedHopperCoins, trayCoins, hopperCoins)] = bestSolution
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
