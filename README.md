# Vending machine

A vending machine implementation in go. 

## Domain

1. Machine -- the vending machine, it contains all the rest of the parts.
2. Tray -- this is the "input tray". It is a staging area for coins that have been inserted but not yet used for a product purchase. After purchase, the tray coins are released
   into the hopper. 
3. Hopper -- this contains the change that is inside the machine. The hopper is also responsible for the purchase, and change logic.
4. Denominations -- the currency / coins that are valid within the machine. 
5. Inventory -- holds the products, and stock available inside the machine.
6. Products -- the product, holds the price and the location (coordinates) in the machine.

## Assumptions

1. This vending machine assumes that customers will want to purchase a product even if exact change cannot be given. In this scenario the machine will give the best possible change
   from the tray and hopper. 
2. This vending machine assumes that there will only be one customer at once. Some of the operations are not thread safe. 
3. This vending machine doesn't currently provide an interface (beyond the API in code). With more time; I would build an HTTP or CLI interface to interact with the machine. 

## Main algorithm

The main purchase and change algorithm is within the coinage/hopper.go file. The `hopper.Deposit` method decides which coins to use for the purchase and change. 

The algorithm works by selecting coins, and performing a depth-first-search on a search tree to calculate if using that coin can meet the price of the item. 
If the coin causes the price to be exceeded, the same depth-first-search over the same tree is used to find the change. 

The algorithm descends the tree looking for the best "solution". The best solution is one that nets payment and change to zero. If this cannot be found, the next best is the
solution that minimises the overpayment by the customer. 

## Machine interactions

1. To restock the machine, or top up the hopper's change, you must first `machine.Unlock()` the machine. This gives access to an `UnlockedMachine` with methods to reload the stock
   and change. 
