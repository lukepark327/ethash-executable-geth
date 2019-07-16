# Remark

* Call `fmt.Printf()` or `spew.Dump()` after `Commit` the trie.

* Use `spew.Dump()`.

## Running Test Code

```bash
$ cd trie
$ go test
```

# Single Element Proof

at `trie/proof_test.go`,

```go
func TestOneElementProof(t *testing.T) {
	trie := new(Trie)
	updateString(trie, "k", "v")

	proofs, _ := ethdb.NewMemDatabase()
	trie.Prove([]byte("k"), 0, proofs)
	if len(proofs.Keys()) != 1 {
		t.Error("proof should have one element")
	}

	val, err, _ := VerifyProof(trie.Hash(), []byte("k"), proofs)
	if err != nil {
		t.Fatalf("VerifyProof error: %v\nproof hashes: %v", err, proofs.Keys())
	}
	if !bytes.Equal(val, []byte("v")) {
		t.Fatalf("VerifyProof returned wrong value: got %x, want 'k'", val)
	}
}
```

## Outputs

```bash
(int) 1
([][]uint8) (len=1 cap=1) {
    ([]uint8) (len=32 cap=32) {
        00000000  66 75 ca 08 7d 4e 43 44  aa 13 48 e5 4d 5b 39 e1  |fu..}NCD..H.M[9.|
        00000010  65 7b 57 28 7e b2 07 10  7a 04 ff ae 79 e8 82 15  |e{W(~...z...y...|
    }
}
(*ethdb.MemDatabase)(0xc4203b9940)({
    db: (map[string][]uint8) (len=1) {
        (string) (len=32) "fu\xca\b}NCD\xaa\x13H\xe5M[9\xe1e{W(~\xb2\a\x10z\x04\xff\xaey\xe8\x82\x15": ([]uint8) (len=5 cap=5) {
            00000000  c4 82 20 6b 76                                    |.. kv|
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
([]uint8) (len=1 cap=8) {
    00000000  76                                                |v|
}
```

# Single Element Void Proof

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

## Outputs

```bash
(int) 1
([][]uint8) (len=1 cap=1) {
    ([]uint8) (len=32 cap=32) {
        00000000  98 02 1e ec 76 a3 52 d4  21 4e e9 d2 2f 26 70 f3  |....v.R.!N../&p.|
        00000010  ab e0 1d 58 05 44 12 49  f4 b7 0d da 75 a0 e0 7a  |...X.D.I....u..z|
    }
}
(*ethdb.MemDatabase)(0xc42033af00)({
    db: (map[string][]uint8) (len=1) {
        (string) (len=32) "\x98\x02\x1e\xecv\xa3R\xd4!N\xe9\xd2/&p\xf3\xab\xe0\x1dX\x05D\x12I\xf4\xb7\r\xdau\xa0\xe0z": ([]uint8) (
len=12 cap=12) {
            00000000  cb 84 20 6b 65 79 85 76  61 6c 75 65              |.. key.value|
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
([]uint8) <nil>
PASS
ok      _/Users/luke/Desktop/go-wired-blockchain/trie   0.701s
```

# Analysis

## `proofs.Len()`

TBA

<!--
```
(int) 1
```
-->

## `proofs.Keys()`

```
([][]uint8) (len=1 cap=1) {
    ([]uint8) (len=32 cap=32) {
        00000000  98 02 1e ec 76 a3 52 d4  21 4e e9 d2 2f 26 70 f3  |....v.R.!N../&p.|
        00000010  ab e0 1d 58 05 44 12 49  f4 b7 0d da 75 a0 e0 7a  |...X.D.I....u..z|
    }
}
```

## `proofs`

```
(*ethdb.MemDatabase)(0xc42033af00)({
    db: (map[string][]uint8) (len=1) {
        (string) (len=32) "\x98\x02\x1e\xecv\xa3R\xd4!N\xe9\xd2/&p\xf3\xab\xe0\x1dX\x05D\x12I\xf4\xb7\r\xdau\xa0\xe0z": ([]uint8) (
len=12 cap=12) {
            00000000  cb 84 20 6b 65 79 85 76  61 6c 75 65              |.. key.value|
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
```

## `val`

```
([]uint8) <nil>
```
