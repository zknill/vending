package domain

type ErrUnknownCoin interface {
	error
	UnknownCoin() uint
}

type ErrUnknownProduct interface {
	error
	UnknownProduct() string
}

type ErrNotEnoughMoney interface {
	error
	RemainingRequired() int
}

type ErrExactChange interface {
	error
	UseExactChange() bool
}

type ErrOutOfStock interface {
	error
	OutOfStockProduct() string
}
