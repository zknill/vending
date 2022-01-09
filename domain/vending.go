package domain

type Product struct {
	Coordinate string
	Price      int
	Name       string
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate
//counterfeiter:generate . Machine
type Machine interface {
	InsertCoin(coin uint) error
	EjectCoins() []uint

	Products() []Product

	Purchase(coordinate string) (Product, []uint, error)
}
