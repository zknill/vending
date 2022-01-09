package machine

import (
	"github.com/zknill/vending/coinage"
	"github.com/zknill/vending/domain"
)

type Machine struct {
	inventory *Inventory
	hopper    *coinage.Hopper
	tray      *coinage.Tray
}

func New(inventory *Inventory, hopper *coinage.Hopper, tray *coinage.Tray) Machine {
	return Machine{
		inventory: inventory,
		hopper:    hopper,
		tray:      tray,
	}
}

func (m Machine) InsertCoin(coin uint) error {
	accepted := m.tray.Insert(coin)

	if !accepted {
		return errUnknownCoin{coin: coin}
	}

	return nil
}

func (m Machine) EjectCoins() []uint {
	return m.tray.Reset()
}

func (m Machine) Purchase(coordinate string) (domain.Product, []uint, error) {
	stock, found := m.inventory.StockLevel(coordinate)
	if !found {
		return domain.Product{}, nil, errUnknownProduct{coordinate: coordinate}
	}

	if stock < 1 {
		return domain.Product{}, nil, errOutOfStock{coordinate: coordinate}
	}

	product := m.inventory.catalog[coordinate]

	if !m.tray.MeetsPrice(product.Price) {
		return domain.Product{}, nil, errNotEnoughMoney{required: product.Price - int(m.tray.Value())}
	}

	change, success := m.hopper.Deposit(m.tray, product.Price)

	if !success {
		return domain.Product{}, nil, errExactChange{}
	}

	return product, change, nil
}

func (m Machine) Unlock() UnlockedMachine {
	return UnlockedMachine{m}
}

type UnlockedMachine struct {
	Machine
}

func (um UnlockedMachine) Restock(inventory map[string]int) {
	for coord, stock := range inventory {
		if _, found := um.inventory.inventory[coord]; !found {
			continue
		}

		um.inventory.inventory[coord] += stock
	}
}

func (um UnlockedMachine) RefloatChange(coins []uint) error {
	return um.hopper.ReFloat(coins)
}
