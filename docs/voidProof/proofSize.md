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

### Test Code: Trie

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

### Proof

Find `a77d337`:

```go
node:
{006137: <835569daf61b98f48e7bc41f6a559862763fcf1c25c8942719aa45a861d2b948> } 

enc(rlped):
[229 131 0 97 55 160 131 85 105 218 246 27 152 244 142 123 196 31 106 85 152 98 118 63 207 28 37 200 148 39 25 170 69 168 97 210 185 72]

hash:
<d7821864f131e69ab04177317473be69bb364ed3b8f5cefe3c5df86101cce89f> 
```

```go
node:
[
  0:  1:  2:  3: <f156b03f5a6cf7c392f1a187c7b035a47f8669ceb5b65434c1de1bf07c09d361> 4:  5:  6: {3639333635: 312e31 } 7:  8:  9:  a:  b:  c:  d:  e:  f:  [17]:  
] 

enc(rlped):
[248 59 128 128 128 160 241 86 176 63 90 108 247 195 146 241 161 135 199 176 53 164 127 134 105 206 181 182 84 52 193 222 27 240 124 9 211 97 128 128 202 133 54 57 51 54 53 131 49 46 49 128 128 128 128 128 128 128 128 128 128]

hash:
<835569daf61b98f48e7bc41f6a559862763fcf1c25c8942719aa45a861d2b948> 
```

```go
node:
[
  0:  1: {2031333535: 34352e30 } 2:  3:  4:  5:  6:  7: <a9097c68a91120ddc39fd75edf5ad0685f5dd4b0a680169853099142a5532edb> 8:  9:  a:  b:  c:  d:  e:  f:  [17]:  
] 

enc(rlped):
[248 60 128 203 133 32 49 51 53 53 132 52 53 46 48 128 128 128 128 128 160 169 9 124 104 169 17 32 221 195 159 215 94 223 90 208 104 95 93 212 176 166 128 22 152 83 9 145 66 165 83 46 219 128 128 128 128 128 128 128 128 128]

hash:
<f156b03f5a6cf7c392f1a187c7b035a47f8669ceb5b65434c1de1bf07c09d361> 
```

```go
node:
{164333: <7b4016f5fc15c94692e203070c448e808c4698d0e2d15dfd28cc47d2e0823162> } 

enc(rlped):
[229 131 22 67 51 160 123 64 22 245 252 21 201 70 146 226 3 7 12 68 142 128 140 70 152 208 226 209 93 253 40 204 71 210 224 130 49 98]

hash:
<a9097c68a91120ddc39fd75edf5ad0685f5dd4b0a680169853099142a5532edb> 
```

```go
node:
[
  0:  1:  2:  3: {2037: 312e3030 } 4:  5:  6:  7:  8:  9: {2037: 302e3132 } a:  b:  c:  d:  e:  f:  [17]:  
] 

enc(rlped):
[225 128 128 128 200 130 32 55 132 49 46 48 48 128 128 128 128 128 200 130 32 55 132 48 46 49 50 128 128 128 128 128 128 128]

hash:
<7b4016f5fc15c94692e203070c448e808c4698d0e2d15dfd28cc47d2e0823162> 
```

```go
(*ethdb.MemDatabase)(0xc4200f2b00)({
    db: (map[string][]uint8) (len=5) {
        (string) (len=32) "\xf1V\xb0?Zl\xf7Ò\xf1\xa1\x87ǰ5\xa4\u007f\x86iε\xb6T4\xc1\xde\x1b\xf0|\t\xd3a": ([]uint8) (len=62 cap=62) {
            00000000  f8 3c 80 cb 85 20 31 33  35 35 84 34 35 2e 30 80  |.<... 1355.45.0.|
            00000010  80 80 80 80 a0 a9 09 7c  68 a9 11 20 dd c3 9f d7  |.......|h.. ....|
            00000020  5e df 5a d0 68 5f 5d d4  b0 a6 80 16 98 53 09 91  |^.Z.h_]......S..|
            00000030  42 a5 53 2e db 80 80 80  80 80 80 80 80 80        |B.S...........|
        },
        (string) (len=32) "\xa9\t|h\xa9\x11 \xddß\xd7^\xdfZ\xd0h_]\u0530\xa6\x80\x16\x98S\t\x91B\xa5S.\xdb": ([]uint8) (len=38 cap=38) {
            00000000  e5 83 16 43 33 a0 7b 40  16 f5 fc 15 c9 46 92 e2  |...C3.{@.....F..|
            00000010  03 07 0c 44 8e 80 8c 46  98 d0 e2 d1 5d fd 28 cc  |...D...F....].(.|
            00000020  47 d2 e0 82 31 62                                 |G...1b|
        },
        (string) (len=32) "{@\x16\xf5\xfc\x15\xc9F\x92\xe2\x03\a\fD\x8e\x80\x8cF\x98\xd0\xe2\xd1]\xfd(\xccG\xd2\xe0\x821b": ([]uint8) (len=34 cap=34) {
            00000000  e1 80 80 80 c8 82 20 37  84 31 2e 30 30 80 80 80  |...... 7.1.00...|
            00000010  80 80 c8 82 20 37 84 30  2e 31 32 80 80 80 80 80  |.... 7.0.12.....|
            00000020  80 80                                             |..|
        },
        (string) (len=32) "ׂ\x18d\xf11暰Aw1ts\xbei\xbb6NӸ\xf5\xce\xfe<]\xf8a\x01\xcc\xe8\x9f": ([]uint8) (len=38 cap=38) {
            00000000  e5 83 00 61 37 a0 83 55  69 da f6 1b 98 f4 8e 7b  |...a7..Ui......{|
            00000010  c4 1f 6a 55 98 62 76 3f  cf 1c 25 c8 94 27 19 aa  |..jU.bv?..%..'..|
            00000020  45 a8 61 d2 b9 48                                 |E.a..H|
        },
        (string) (len=32) "\x83Ui\xda\xf6\x1b\x98\xf4\x8e{\xc4\x1fjU\x98bv?\xcf\x1c%Ȕ'\x19\xaaE\xa8aҹH": ([]uint8) (len=61 cap=61) {
            00000000  f8 3b 80 80 80 a0 f1 56  b0 3f 5a 6c f7 c3 92 f1  |.;.....V.?Zl....|
            00000010  a1 87 c7 b0 35 a4 7f 86  69 ce b5 b6 54 34 c1 de  |....5...i...T4..|
            00000020  1b f0 7c 09 d3 61 80 80  ca 85 36 39 33 36 35 83  |..|..a....69365.|
            00000030  31 2e 31 80 80 80 80 80  80 80 80 80 80           |1.1..........|
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
})
PASS
ok      _/Users/luke/Desktop/go-wired-blockchain/trie   0.029s
```





# References

* https://github.com/ethereum/wiki/wiki/Patricia-Tree

* https://medium.com/coinmonks/ethereum-account-212feb9c4154
