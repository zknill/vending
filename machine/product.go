package machine

import "fmt"

type Product struct {
	Coordinate string
	Price      uint
	Name       string
}

type Inventory struct {
	catalog   map[string]Product
	inventory map[string]int
}

func NewInventory(products ...Product) (*Inventory, error) {
	i := &Inventory{
		catalog:   map[string]Product{},
		inventory: map[string]int{},
	}

	for _, p := range products {
		if _, exists := i.catalog[p.Coordinate]; exists {
			return nil, fmt.Errorf("only one product in each coordinate")
		}

		i.catalog[p.Coordinate] = p
	}

	return i, nil
}

func (i *Inventory) ModifyStock(coordinate string, change int) error {
	if _, exists := i.catalog[coordinate]; !exists {
		return fmt.Errorf("unknown coordinate: %s", coordinate)
	}

	if i.inventory[coordinate]+change < 0 {
		return fmt.Errorf("not enough stock: %s", coordinate)
	}

	i.inventory[coordinate] += change

	return nil
}

func (i *Inventory) StockLevel(coordinate string) (int, bool) {
	stock, found := i.inventory[coordinate]
	return stock, found
}
