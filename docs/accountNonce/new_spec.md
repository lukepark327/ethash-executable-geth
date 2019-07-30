* Using a function `newObject` in `core/state/state_object.go`.

* Modifying a function `createObject` called by `CreateAccount` in `core/state/statedb.go`.

```go
// createObject creates a new state object. If there is an existing account with
// the given address, it is overwritten and returned as the second return value.
func (self *StateDB) createObject(addr common.Address) (newobj, prev *stateObject) {
	prev = self.getStateObject(addr)
	newobj = newObject(self, addr, Account{})
	newobj.setNonce(0) // sets the object to dirty
	if prev == nil {
		self.journal.append(createObjectChange{account: &addr})
	} else {
		self.journal.append(resetObjectChange{prev: prev})
	}
	self.setStateObject(newobj)
	return newobj, prev
}
```

## Remark

```go
newobj.setNonce(0) // sets the object to dirty
```


```go
var (
	MaxTransactionLimit = uint64(64)
)

func (self *StateDB) createObject_eth4nos(addr common.Address, base *big.Int) (newobj, prev *stateObject) {
	prev = self.getStateObject(addr)
	newobj = newObject(self, addr, Account{})

	log.Info("NONCE:", "nonce", base.Uint64()*MaxTransactionLimit)

	newobj.setNonce(base.Uint64() * MaxTransactionLimit)

	if prev == nil {
		self.journal.append(createObjectChange{account: &addr})
	} else {
		self.journal.append(resetObjectChange{prev: prev})
	}
	self.setStateObject(newobj)
	return newobj, prev
}
```

CreateAccount

* ApplyDAOHardFork in `consensus/misc/dao.go`

	* mineBlock in `cmd/geth/retesteth.go`
	
	* GenerateChain in `core/chain_makers.go`
	
	* Process in `core/state_processor.go`
	
	* commitNewWork in `miner/worker.go`


* GetOrNewStateObject

	* AddBalance
	
		* Transfer
	
	
	* SubBalance
	
	* SetBalance
	
	* SetNonce
	
	* SetCode
	
	* SetState
	

