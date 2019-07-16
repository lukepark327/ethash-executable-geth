# Trie

## Remark

* Call `fmt.Printf()` or `spew.Dump()` after `Commit` the trie.

## Test Code

### Single Element Void Proof

at `trie/proof_test.go`,

```go
func TestMyOwnTestCode(t *testing.T) {
	trie := new(Trie)
	updateString(trie, "key", "value")

	proofs, _ := ethdb.NewMemDatabase()
  
	// Void Proof
	trie.Prove([]byte("k"), 0, proofs)
	
	// Print
	spew.Dump(proofs.Len())
	spew.Dump(proofs.Keys())
	spew.Dump(proofs)

	// Verifying Void Proof
	val, err, _ := VerifyProof(trie.Hash(), []byte("k"), proofs)
	if err != nil {
		t.Fatalf("VerifyProof error: %v\nproof hashes: %v", err, proofs.Keys())
	}
	
	// Print
	spew.Dump(val)
}
```

### Running Test Code

```bash
$ cd trie
$ go test
```

# `trie.Prove()`

# `VerifyProof()`
