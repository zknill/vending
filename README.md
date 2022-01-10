# Vending machine

A vending machine implementation in go.

To run the tests;

```bash
$ make test
# or
$ go test ./...
```

Example;
```go
// Initialise a new machine with an inventory of
// products, a hopper for the change, and a tray
// for the inserted coins.
var machine = machine.New(inventory, hopper, tray)

// View the products
products = machine.Products()

// Insert coins one-by-one, like a customer
// would do at a real vending machine.
machine.InsertCoin(50)
machine.InsertCoin(5)
machine.InsertCoin(1)

// Purchase using the product's coordinates. Row A, Column 1.
// The machine will return the product and the change.
// There are error cases for out of stock, not enough money, etc.
product, change, err := machine.Purchase("A1")

// Unlock the machine to access it's "internals".
unlockedMachine := machine.Unlock()

// Restock the machine
unlockedMachine.Restock(map[string]int{"A1": 3})

// Add more change to the machine
unlockedMachine.RefloatChange([]uint{1, 1, 1, 5, 5, 25, 50, 100})
```

## Domain

1. **Machine** -- the vending machine, it contains all the rest of the parts.
2. **Tray** -- this is the "input tray". It is a staging area for coins that have been inserted but not yet used for a product purchase. After purchase, the tray coins are released
   into the hopper.
3. **Hopper** -- this contains the change that is inside the machine. The hopper is also responsible for the purchase, and change logic.
4. **Denominations** -- the currency / coins that are valid within the machine.
5. **Inventory** -- holds the products, and stock available inside the machine.
6. **Products** -- the product, holds the price and the location (coordinates) in the machine.

## Assumptions

1. This vending machine assumes that customers will want to purchase a product even if exact change cannot be given. In this scenario the machine will give the best possible change
   from the tray and hopper.
2. This vending machine assumes that there will only be one customer at once. Some of the operations are not thread safe.
3. This vending machine doesn't currently provide a proper interface (beyond the API in code). With more time; I would build a threadsafe HTTP or CLI interface to interact with the machine.
4. There are no interfaces for dependency injection, I could replace the `Hopper`, `Tray`, and `Inventory` with interfaces that would allow for dependency injection. i.e. an
   inventory that stored its stock data in a database instead of in memory. The only interface is for the Machine, `domain.Machine`. This has mocks for use in the http server
   tests.

## Main algorithm

The main purchase and change algorithm is within the coinage/hopper.go file. The `hopper.Deposit` method decides which coins to use for the purchase and change.

The algorithm works by selecting coins, and performing a depth-first-search on a search tree to calculate if using that coin can meet the price of the item.
If the coin causes the price to be exceeded, the same depth-first-search over the same tree is used to find the change.

The algorithm descends the tree looking for the best "solution". The best solution is one that nets payment and change to zero. If this cannot be found, the next best is the
solution that minimises the overpayment by the customer.

## Toy server

There's a toy server implemented in the http package, its not thread safe as it builds on the assumptions above^.
The server does show some APIs, error handling, and request / response bodies. There are tests for the server.
