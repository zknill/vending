// Code generated by counterfeiter. DO NOT EDIT.
package domainfakes

import (
	"sync"

	"github.com/zknill/vending/domain"
)

type FakeMachine struct {
	EjectCoinsStub        func() []uint
	ejectCoinsMutex       sync.RWMutex
	ejectCoinsArgsForCall []struct {
	}
	ejectCoinsReturns struct {
		result1 []uint
	}
	ejectCoinsReturnsOnCall map[int]struct {
		result1 []uint
	}
	InsertCoinStub        func(uint) error
	insertCoinMutex       sync.RWMutex
	insertCoinArgsForCall []struct {
		arg1 uint
	}
	insertCoinReturns struct {
		result1 error
	}
	insertCoinReturnsOnCall map[int]struct {
		result1 error
	}
	ProductsStub        func() []domain.Product
	productsMutex       sync.RWMutex
	productsArgsForCall []struct {
	}
	productsReturns struct {
		result1 []domain.Product
	}
	productsReturnsOnCall map[int]struct {
		result1 []domain.Product
	}
	PurchaseStub        func(string) (domain.Product, []uint, error)
	purchaseMutex       sync.RWMutex
	purchaseArgsForCall []struct {
		arg1 string
	}
	purchaseReturns struct {
		result1 domain.Product
		result2 []uint
		result3 error
	}
	purchaseReturnsOnCall map[int]struct {
		result1 domain.Product
		result2 []uint
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMachine) EjectCoins() []uint {
	fake.ejectCoinsMutex.Lock()
	ret, specificReturn := fake.ejectCoinsReturnsOnCall[len(fake.ejectCoinsArgsForCall)]
	fake.ejectCoinsArgsForCall = append(fake.ejectCoinsArgsForCall, struct {
	}{})
	stub := fake.EjectCoinsStub
	fakeReturns := fake.ejectCoinsReturns
	fake.recordInvocation("EjectCoins", []interface{}{})
	fake.ejectCoinsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeMachine) EjectCoinsCallCount() int {
	fake.ejectCoinsMutex.RLock()
	defer fake.ejectCoinsMutex.RUnlock()
	return len(fake.ejectCoinsArgsForCall)
}

func (fake *FakeMachine) EjectCoinsCalls(stub func() []uint) {
	fake.ejectCoinsMutex.Lock()
	defer fake.ejectCoinsMutex.Unlock()
	fake.EjectCoinsStub = stub
}

func (fake *FakeMachine) EjectCoinsReturns(result1 []uint) {
	fake.ejectCoinsMutex.Lock()
	defer fake.ejectCoinsMutex.Unlock()
	fake.EjectCoinsStub = nil
	fake.ejectCoinsReturns = struct {
		result1 []uint
	}{result1}
}

func (fake *FakeMachine) EjectCoinsReturnsOnCall(i int, result1 []uint) {
	fake.ejectCoinsMutex.Lock()
	defer fake.ejectCoinsMutex.Unlock()
	fake.EjectCoinsStub = nil
	if fake.ejectCoinsReturnsOnCall == nil {
		fake.ejectCoinsReturnsOnCall = make(map[int]struct {
			result1 []uint
		})
	}
	fake.ejectCoinsReturnsOnCall[i] = struct {
		result1 []uint
	}{result1}
}

func (fake *FakeMachine) InsertCoin(arg1 uint) error {
	fake.insertCoinMutex.Lock()
	ret, specificReturn := fake.insertCoinReturnsOnCall[len(fake.insertCoinArgsForCall)]
	fake.insertCoinArgsForCall = append(fake.insertCoinArgsForCall, struct {
		arg1 uint
	}{arg1})
	stub := fake.InsertCoinStub
	fakeReturns := fake.insertCoinReturns
	fake.recordInvocation("InsertCoin", []interface{}{arg1})
	fake.insertCoinMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeMachine) InsertCoinCallCount() int {
	fake.insertCoinMutex.RLock()
	defer fake.insertCoinMutex.RUnlock()
	return len(fake.insertCoinArgsForCall)
}

func (fake *FakeMachine) InsertCoinCalls(stub func(uint) error) {
	fake.insertCoinMutex.Lock()
	defer fake.insertCoinMutex.Unlock()
	fake.InsertCoinStub = stub
}

func (fake *FakeMachine) InsertCoinArgsForCall(i int) uint {
	fake.insertCoinMutex.RLock()
	defer fake.insertCoinMutex.RUnlock()
	argsForCall := fake.insertCoinArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeMachine) InsertCoinReturns(result1 error) {
	fake.insertCoinMutex.Lock()
	defer fake.insertCoinMutex.Unlock()
	fake.InsertCoinStub = nil
	fake.insertCoinReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMachine) InsertCoinReturnsOnCall(i int, result1 error) {
	fake.insertCoinMutex.Lock()
	defer fake.insertCoinMutex.Unlock()
	fake.InsertCoinStub = nil
	if fake.insertCoinReturnsOnCall == nil {
		fake.insertCoinReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.insertCoinReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeMachine) Products() []domain.Product {
	fake.productsMutex.Lock()
	ret, specificReturn := fake.productsReturnsOnCall[len(fake.productsArgsForCall)]
	fake.productsArgsForCall = append(fake.productsArgsForCall, struct {
	}{})
	stub := fake.ProductsStub
	fakeReturns := fake.productsReturns
	fake.recordInvocation("Products", []interface{}{})
	fake.productsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeMachine) ProductsCallCount() int {
	fake.productsMutex.RLock()
	defer fake.productsMutex.RUnlock()
	return len(fake.productsArgsForCall)
}

func (fake *FakeMachine) ProductsCalls(stub func() []domain.Product) {
	fake.productsMutex.Lock()
	defer fake.productsMutex.Unlock()
	fake.ProductsStub = stub
}

func (fake *FakeMachine) ProductsReturns(result1 []domain.Product) {
	fake.productsMutex.Lock()
	defer fake.productsMutex.Unlock()
	fake.ProductsStub = nil
	fake.productsReturns = struct {
		result1 []domain.Product
	}{result1}
}

func (fake *FakeMachine) ProductsReturnsOnCall(i int, result1 []domain.Product) {
	fake.productsMutex.Lock()
	defer fake.productsMutex.Unlock()
	fake.ProductsStub = nil
	if fake.productsReturnsOnCall == nil {
		fake.productsReturnsOnCall = make(map[int]struct {
			result1 []domain.Product
		})
	}
	fake.productsReturnsOnCall[i] = struct {
		result1 []domain.Product
	}{result1}
}

func (fake *FakeMachine) Purchase(arg1 string) (domain.Product, []uint, error) {
	fake.purchaseMutex.Lock()
	ret, specificReturn := fake.purchaseReturnsOnCall[len(fake.purchaseArgsForCall)]
	fake.purchaseArgsForCall = append(fake.purchaseArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.PurchaseStub
	fakeReturns := fake.purchaseReturns
	fake.recordInvocation("Purchase", []interface{}{arg1})
	fake.purchaseMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeMachine) PurchaseCallCount() int {
	fake.purchaseMutex.RLock()
	defer fake.purchaseMutex.RUnlock()
	return len(fake.purchaseArgsForCall)
}

func (fake *FakeMachine) PurchaseCalls(stub func(string) (domain.Product, []uint, error)) {
	fake.purchaseMutex.Lock()
	defer fake.purchaseMutex.Unlock()
	fake.PurchaseStub = stub
}

func (fake *FakeMachine) PurchaseArgsForCall(i int) string {
	fake.purchaseMutex.RLock()
	defer fake.purchaseMutex.RUnlock()
	argsForCall := fake.purchaseArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeMachine) PurchaseReturns(result1 domain.Product, result2 []uint, result3 error) {
	fake.purchaseMutex.Lock()
	defer fake.purchaseMutex.Unlock()
	fake.PurchaseStub = nil
	fake.purchaseReturns = struct {
		result1 domain.Product
		result2 []uint
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeMachine) PurchaseReturnsOnCall(i int, result1 domain.Product, result2 []uint, result3 error) {
	fake.purchaseMutex.Lock()
	defer fake.purchaseMutex.Unlock()
	fake.PurchaseStub = nil
	if fake.purchaseReturnsOnCall == nil {
		fake.purchaseReturnsOnCall = make(map[int]struct {
			result1 domain.Product
			result2 []uint
			result3 error
		})
	}
	fake.purchaseReturnsOnCall[i] = struct {
		result1 domain.Product
		result2 []uint
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeMachine) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.ejectCoinsMutex.RLock()
	defer fake.ejectCoinsMutex.RUnlock()
	fake.insertCoinMutex.RLock()
	defer fake.insertCoinMutex.RUnlock()
	fake.productsMutex.RLock()
	defer fake.productsMutex.RUnlock()
	fake.purchaseMutex.RLock()
	defer fake.purchaseMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeMachine) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ domain.Machine = new(FakeMachine)
