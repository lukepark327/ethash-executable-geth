# Tries in Ethereum

## State Trie

* There is one global state trie, and it updates over time.

* In it, a path is always: `sha3(ethereumAddress)` and a value is always: `rlp(ethereumAccount)`.

  * More specifically an ethereum account is a 4 item array of [nonce,balance,storageRoot,codeHash].
  
    * At this point it's worth noting that this storageRoot is the root of another patricia trie:
    
    * The codeHash is the hash of the EVM (Ethereum Virtual Machine) code of this account. For contract accounts, this is the code that gets hashed and stored as the codeHash. For externally owned accounts, the codeHash field is the hash of the empty string.

### Details

See an `Account` structure at `core/state/state_object.go`:

```go
// Account is the Ethereum consensus representation of accounts.
// These objects are stored in the main account trie.
type Account struct {
	Nonce    uint64
	Balance  *big.Int
	Root     common.Hash // merkle root of the storage trie
	CodeHash []byte
}
```

Actually Ethereum uses no `rlp(ethereumAccount)` but `rlp(stateObject)` to the trie. See `core/state/statedb.go`:

```go
// updateStateObject writes the given object to the trie.
func (self *StateDB) updateStateObject(stateObject *stateObject) {
	addr := stateObject.Address()
	data, err := rlp.EncodeToBytes(stateObject)
	if err != nil {
		panic(fmt.Errorf("can't encode object at %x: %v", addr[:], err))
	}
	self.setError(self.trie.TryUpdate(addr[:], data))
}
```

where `stateObject` is:

```go
// stateObject represents an Ethereum account which is being modified.
//
// The usage pattern is as follows:
// First you need to obtain a state object.
// Account values can be accessed and modified through the object.
// Finally, call CommitTrie to write the modified storage trie into a database.
type stateObject struct {
	address  common.Address
	addrHash common.Hash // hash of ethereum address of the account
	data     Account
	db       *StateDB

	// DB error.
	// State objects are used by the consensus core and VM which are
	// unable to deal with database-level errors. Any error that occurs
	// during a database read is memoized here and will eventually be returned
	// by StateDB.Commit.
	dbErr error

	// Write caches.
	trie Trie // storage trie, which becomes non-nil on first access
	code Code // contract bytecode, which gets set when code is loaded

	cachedStorage Storage // Storage entry cache to avoid duplicate reads
	dirtyStorage  Storage // Storage entries that need to be flushed to disk

	// Cache flags.
	// When an object is marked suicided it will be delete from the trie
	// during the "update" phase of the state transition.
	dirtyCode bool // true if the code was updated
	suicided  bool
	touched   bool
	deleted   bool
	onDirty   func(addr common.Address) // Callback method to mark a state object newly dirty
}
```

### Trie Structure

![MPT](https://miro.medium.com/max/5040/1*ZbdfL2TWmxj4b1fCuN6BIQ.png)

A node in a Merkle Patricia trie is one of the following:

1. `NULL` (represented as the empty string)
2. `branch` A 17-item node `[ v0 ... v15, vt ]`
3. `leaf` A 2-item node `[ encodedPath, value ]`
4. `extension` A 2-item node `[ encodedPath, key ]`

# References

* https://github.com/ethereum/wiki/wiki/Patricia-Tree

* https://medium.com/coinmonks/ethereum-account-212feb9c4154
