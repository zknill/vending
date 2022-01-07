package machine

import (
	"errors"

	"github.com/zknill/vending/coinage"
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
		return errors.New("unrecognised coin")
	}

	return nil
}

func (m Machine) EjectCoins() []uint {
	return m.tray.Reset()
}

func (m Machine) Purchase(coordinate string) (Product, []uint, error) {
	stock, found := m.inventory.StockLevel(coordinate)
	if !found {
		return Product{}, nil, errors.New("unknown product coordinate")
	}

	if stock < 1 {
		return Product{}, nil, errors.New("out of stock")
	}

	product := m.inventory.catalog[coordinate]

	if !m.tray.MeetsPrice(product.Price) {
		return Product{}, nil, errors.New("not enough money")
	}

	change, success := m.hopper.Deposit(m.tray, product.Price)

	if !success {
		return Product{}, nil, errors.New("use exact change")
	}

	return product, change, nil
}

func (m Machine) Unlock(key *Key) UnlockedMachine {
	return UnlockedMachine{m}
}

type Key struct{}

type UnlockedMachine struct {
	Machine
}

func (um UnlockedMachine) Restock(inventory map[string]int) {

}

func (um UnlockedMachine) RefloatChange(coins []uint) {

}
