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

# Multiple Elements Void Proof

at `trie/proof_test.go`,

```go
func TestMyOwnTestCode(t *testing.T) {
	trie := new(Trie)
	updateString(trie, "key", "value")
	updateString(trie, "keccak", "sha3")
	updateString(trie, "k1", "v1")
	// updateString(trie, "k2", "v2")
	updateString(trie, "k3", "v3")
	updateString(trie, "k4", "v4")
	updateString(trie, "k5", "v5")
	updateString(trie, "tmp", "blahblah")

	proofs, _ := ethdb.NewMemDatabase()
	trie.Prove([]byte("k2"), 0, proofs)

	spew.Dump(proofs.Len())
	spew.Dump(proofs.Keys())
	spew.Dump(proofs)

	val, err, _ := VerifyProof(trie.Hash(), []byte("k2"), proofs)
	if err != nil {
		t.Fatalf("VerifyProof error: %v\nproof hashes: %v", err, proofs.Keys())
	}

	spew.Dump(val)
}
```

## Outputs

```bash
(int) 4
([][]uint8) (len=4 cap=4) {
    ([]uint8) (len=32 cap=32) {
        00000000  c4 06 fc 66 80 b7 56 69  b8 8f 7c b7 1b a4 a3 45  |...f..Vi..|....E|
        00000010  bb 7d 7d 09 79 e1 6d 88  b0 5c d7 b4 8f bd 80 2d  |.}}.y.m..\.....-|
    },
    ([]uint8) (len=32 cap=32) {
        00000000  8e 19 e9 c5 0b d9 05 5e  09 ae be 04 2c b3 9b 80  |.......^....,...|
        00000010  a4 d9 fe 1d bc 5b b0 ae  52 52 b4 4b 97 71 8c 54  |.....[..RR.K.q.T|
    },
    ([]uint8) (len=32 cap=32) {
        00000000  d9 24 38 5f b2 77 90 af  57 5b e2 02 df 10 74 96  |.$8_.w..W[....t.|
        00000010  92 e3 34 eb 1a 76 62 bb  76 19 7e ce 67 12 95 9d  |..4..vb.v.~.g...|
    },
    ([]uint8) (len=32 cap=32) {
        00000000  8b e6 67 54 38 ff 40 e5  b1 d7 0a 71 cb 84 14 f4  |..gT8.@....q....|
        00000010  c6 d5 49 2a c2 f3 f8 39  7d d4 4f de 1f c0 9f 0e  |..I*...9}.O.....|
    }
}
(*ethdb.MemDatabase)(0xc420185860)({
    db: (map[string][]uint8) (len=4) {
        (string) (len=32) "\xc4\x06\xfcf\x80\xb7Vi\xb8\x8f|\xb7\x1b\xa4\xa3E\xbb}}\ty\xe1m\x88\xb0\\״\x8f\xbd\x80-": ([]uint8) (len=64 cap=64) {
            00000000  f8 3e 80 80 80 80 80 80  a0 8e 19 e9 c5 0b d9 05  |.>..............|
            00000010  5e 09 ae be 04 2c b3 9b  80 a4 d9 fe 1d bc 5b b0  |^....,........[.|
            00000020  ae 52 52 b4 4b 97 71 8c  54 cd 83 34 6d 70 88 62  |.RR.K.q.T..4mp.b|
            00000030  6c 61 68 62 6c 61 68 80  80 80 80 80 80 80 80 80  |lahblah.........|
        },
        (string) (len=32) "\x8e\x19\xe9\xc5\v\xd9\x05^\t\xae\xbe\x04,\xb3\x9b\x80\xa4\xd9\xfe\x1d\xbc[\xb0\xaeRR\xb4K\x97q\x8cT": ([]uint8) (len=35 cap=35) {
            00000000  e2 1b a0 d9 24 38 5f b2  77 90 af 57 5b e2 02 df  |....$8_.w..W[...|
            00000010  10 74 96 92 e3 34 eb 1a  76 62 bb 76 19 7e ce 67  |.t...4..vb.v.~.g|
            00000020  12 95 9d                                          |...|
        },
        (string) (len=32) "\xd9$8_\xb2w\x90\xafW[\xe2\x02\xdf\x10t\x96\x92\xe34\xeb\x1avb\xbbv\x19~\xceg\x12\x95\x9d": ([]uint8) (len=83 cap=83) {
            00000000  f8 51 80 80 80 a0 8b e6  67 54 38 ff 40 e5 b1 d7  |.Q......gT8.@...|
            00000010  0a 71 cb 84 14 f4 c6 d5  49 2a c2 f3 f8 39 7d d4  |.q......I*...9}.|
            00000020  4f de 1f c0 9f 0e 80 80  a0 39 08 c1 b6 37 10 38  |O........9...7.8|
            00000030  9a 63 e1 e5 c4 a8 61 8f  ce ea 4c 1a 15 ae fd c7  |.c....a...L.....|
            00000040  ef b4 35 40 41 e8 96 6a  3e 80 80 80 80 80 80 80  |..5@A..j>.......|
            00000050  80 80 80                                          |...|
        },
        (string) (len=32) "\x8b\xe6gT8\xff@\xe5\xb1\xd7\nq˄\x14\xf4\xc6\xd5I*\xc2\xf3\xf89}\xd4O\xde\x1f\xc0\x9f\x0e": ([]uint8) (len=34 cap=34) {
            00000000  e1 80 c4 20 82 76 31 80  c4 20 82 76 33 c4 20 82  |... .v1.. .v3. .|
            00000010  76 34 c4 20 82 76 35 80  80 80 80 80 80 80 80 80  |v4. .v5.........|
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
})
([]uint8) <nil>
PASS
ok      _/Users/luke/Desktop/go-wired-blockchain/trie   0.691s
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
