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

## Trie Structure

![MPT](https://miro.medium.com/max/5040/1*ZbdfL2TWmxj4b1fCuN6BIQ.png)

A node in a Merkle Patricia trie is one of the following:

1. `NULL` (represented as the empty string)
2. `branch` A 17-item node `[ v0 ... v15, vt ]`
3. `leaf` A 2-item node `[ encodedPath, value ]`
4. `extension` A 2-item node `[ encodedPath, key ]`

### Test Code

at `trie/trie_test.go`,

```go
func TestMyOwnTrie(t *testing.T) {
	trie := newEmpty()
	updateString(trie, "a711355", "45.0")
	updateString(trie, "a77d337", "1.00")
	updateString(trie, "a7f9365", "1.1")
	updateString(trie, "a77d397", "0.12")

	// root
	_, err := trie.Commit()
	if err != nil {
		t.Fatalf("commit error: %v", err)
	}

	// spew.Dump(root)
	spew.Dump(trie)
}
```

returns

```go
(*trie.Trie)(0xc42008c3c0)({
    root: (*trie.shortNode)(0xc42008c780)({06010307: [
    0: <nil> 1: <nil> 2: <nil> 3: [
      0: <nil> 1: {030103030305030510: 34352e30 } 2: <nil> 3: <nil> 4: <nil> 5: <nil> 6: <nil> 7: {0604030
303: [
          0: <nil> 1: <nil> 2: <nil> 3: {030710: 312e3030 } 4: <nil> 5: <nil> 6: <nil> 7: <nil> 8: <nil> 9
: {030710: 302e3132 } a: <nil> b: <nil> c: <nil> d: <nil> e: <nil> f: <nil> [17]: <nil> 
        ] } 8: <nil> 9: <nil> a: <nil> b: <nil> c: <nil> d: <nil> e: <nil> f: <nil> [17]: <nil> 
    ] 4: <nil> 5: <nil> 6: {06030903030306030510: 312e31 } 7: <nil> 8: <nil> 9: <nil> a: <nil> b: <nil> c: <nil> d: <nil> e: <nil> f: <nil> [17]: <nil> 
  ] } ),
    db: (*ethdb.MemDatabase)(0xc4200f2aa0)({
        db: (map[string][]uint8) (len=5) {
            (string) (len=32) "\xa9\t|h\xa9\x11 \xddß\xd7^\xdfZ\xd0h_]\u0530\xa6\x80\x16\x98S\t\x91B\xa5S.\xdb": ([]uint8) (len=38 cap=38) {
                00000000  e5 83 16 43 33 a0 7b 40  16 f5 fc 15 c9 46 92 e2  |...C3.{@.....F..|
                00000010  03 07 0c 44 8e 80 8c 46  98 d0 e2 d1 5d fd 28 cc  |...D...F....].(.|
                00000020  47 d2 e0 82 31 62                                 |G...1b|
            },
            (string) (len=32) "\xf1V\xb0?Zl\xf7Ò\xf1\xa1\x87ǰ5\xa4\u007f\x86iε\xb6T4\xc1\xde\x1b\xf0|\t\xd3a": ([]uint8) (len=62 cap=62) {
                00000000  f8 3c 80 cb 85 20 31 33  35 35 84 34 35 2e 30 80  |.<... 1355.45.0.|
                00000010  80 80 80 80 a0 a9 09 7c  68 a9 11 20 dd c3 9f d7  |.......|h.. ....|
                00000020  5e df 5a d0 68 5f 5d d4  b0 a6 80 16 98 53 09 91  |^.Z.h_]......S..|
                00000030  42 a5 53 2e db 80 80 80  80 80 80 80 80 80        |B.S...........|
            },
            (string) (len=32) "\x83Ui\xda\xf6\x1b\x98\xf4\x8e{\xc4\x1fjU\x98bv?\xcf\x1c%Ȕ'\x19\xaaE\xa8aҹH": ([]uint8) (len=61 cap=61) {
                00000000  f8 3b 80 80 80 a0 f1 56  b0 3f 5a 6c f7 c3 92 f1  |.;.....V.?Zl....|
                00000010  a1 87 c7 b0 35 a4 7f 86  69 ce b5 b6 54 34 c1 de  |....5...i...T4..|
                00000020  1b f0 7c 09 d3 61 80 80  ca 85 36 39 33 36 35 83  |..|..a....69365.|
                00000030  31 2e 31 80 80 80 80 80  80 80 80 80 80           |1.1..........|
            },
            (string) (len=32) "ׂ\x18d\xf11暰Aw1ts\xbei\xbb6NӸ\xf5\xce\xfe<]\xf8a\x01\xcc\xe8\x9f": ([]uint8) (len=38 cap=38) {
                00000000  e5 83 00 61 37 a0 83 55  69 da f6 1b 98 f4 8e 7b  |...a7..Ui......{|
                00000010  c4 1f 6a 55 98 62 76 3f  cf 1c 25 c8 94 27 19 aa  |..jU.bv?..%..'..|
                00000020  45 a8 61 d2 b9 48                                 |E.a..H|
            },
            (string) (len=32) "{@\x16\xf5\xfc\x15\xc9F\x92\xe2\x03\a\fD\x8e\x80\x8cF\x98\xd0\xe2\xd1]\xfd(\xccG\xd2\xe0\x821b": ([]uint8) (len=34 cap=34) {
                00000000  e1 80 80 80 c8 82 20 37  84 31 2e 30 30 80 80 80  |...... 7.1.00...|
                00000010  80 80 c8 82 20 37 84 30  2e 31 32 80 80 80 80 80  |.... 7.0.12.....|
                00000020  80 80                                             |..|
            }
        },
        lock: (sync.RWMutex) {
            w: (sync.Mutex) {
                state: (int32) 0,
                sema: (uint32) 0
            },
            writerSem: (uint32) 0,
            readerSem: (uint32) 0,
            readerCount: (int32) 0,
            readerWait: (int32) 0
        }
    }),
    originalRoot: (common.Hash) (len=32 cap=32) 0x0000000000000000000000000000000000000000000000000000000000000000,
    cachegen: (uint16) 1,
    cachelimit: (uint16) 0
})
PASS
ok      _/Users/luke/Desktop/go-wired-blockchain/trie   0.028s
```

# References

* https://github.com/ethereum/wiki/wiki/Patricia-Tree

* https://medium.com/coinmonks/ethereum-account-212feb9c4154
