# Construct Merkle Proof

## `func (t *Trie) Prove`

at `trie/proof.go`

* Prove constructs a merkle proof for key.

* The result contains all encoded nodes on the path to the value at key.

  * The value itself is also included in the last node and can be retrieved by verifying the proof.

* If the trie does not contain a value for key,

  * the returned proof contains all nodes of the longest existing prefix of the key (at least the root node),
  * ending with the node that **proves the absence of the key.**

## `func TestMissingKeyProof(t *testing.T)`

at `trie/proof_test.go`

* Tests that missing keys can also be proven.

** The test explicitly uses a single entry trie and checks for missing keys both before and after the single entry.

## Verify Merkle Proof

## `func VerifyProof(rootHash common.Hash, key []byte, proofDb ethdb.KeyValueReader) (value []byte, nodes int, err error)`

at `trie/proof.go`

* VerifyProof checks merkle proofs.

* The given proof must contain the value for key in a trie with the given root hash.

* VerifyProof returns an error if the proof contains invalid trie nodes or the wrong value.

### Usage

For example,

```go
proof := memorydb.New()

trie.Prove([]byte(key), 0, proof)
val, _, err := VerifyProof(trie.Hash(), []byte(key), proof)
```

### Code Review

```go
func VerifyProof(rootHash common.Hash, key []byte, proofDb ethdb.KeyValueReader) (value []byte, nodes int, err error) {
	key = keybytesToHex(key)
	wantHash := rootHash
	for i := 0; ; i++ {
		buf, _ := proofDb.Get(wantHash[:])
		if buf == nil {
			return nil, i, fmt.Errorf("proof node %d (hash %064x) missing", i, wantHash)
		}
		n, err := decodeNode(wantHash[:], buf)
		if err != nil {
			return nil, i, fmt.Errorf("bad proof node %d: %v", i, err)
		}
		keyrest, cld := get(n, key)
		switch cld := cld.(type) {
		case nil:
			// The trie doesn't contain the key.
			return nil, i, nil
		case hashNode:
			key = keyrest
			copy(wantHash[:], cld)
		case valueNode:
			return cld, i + 1, nil
		}
	}
}
```
