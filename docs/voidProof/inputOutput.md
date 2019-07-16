# Remark

* Call `fmt.Printf()` or `spew.Dump()` after `Commit` the trie.

* Use `spew.Dump()`.

## Running Test Code

```bash
$ cd trie
$ go test -run <NameOfTest>
```

<!--
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

```go
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

```go
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
-->

<!--
# 🤔 Multiple Elements Proof

at `trie/proof_test.go`,

```go
func TestMyOwnTestCode(t *testing.T) {
	trie := new(Trie)
	updateString(trie, "key", "value")
	updateString(trie, "keccak", "sha3")
	updateString(trie, "k1", "v1")
	updateString(trie, "k2", "v2")
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

```go
(int) 4
([][]uint8) (len=4 cap=4) {
    ([]uint8) (len=32 cap=32) {
        00000000  52 8d 52 cb 60 bf 19 05  af f5 a0 5a b8 bb 95 67  |R.R.`......Z...g|
        00000010  4d 15 44 10 5c be 85 96  7b da eb 39 ee 22 2e bd  |M.D.\...{..9."..|
    },
    ([]uint8) (len=32 cap=32) {
        00000000  6d d0 2b 05 08 64 b7 72  a7 bf e4 66 95 ff aa 60  |m.+..d.r...f...`|
        00000010  ff 9b b7 3f 99 66 26 c3  84 01 27 16 3d bb 9b 36  |...?.f&...'.=..6|
    },
    ([]uint8) (len=32 cap=32) {
        00000000  01 70 26 65 ba 59 2f f3  ae 04 f6 52 c8 31 27 55  |.p&e.Y/....R.1'U|
        00000010  42 9f 23 e1 66 0d 86 80  e4 ee 03 a5 af f7 5e 7d  |B.#.f.........^}|
    },
    ([]uint8) (len=32 cap=32) {
        00000000  49 c4 a7 7f b3 d9 4a 36  ba 6e 33 77 45 77 01 82  |I.....J6.n3wEw..|
        00000010  22 95 40 83 11 32 cc 4e  46 e2 5b a5 c4 ea 3e 29  |".@..2.NF.[...>)|
    }
}
(*ethdb.MemDatabase)(0xc42031e440)({
    db: (map[string][]uint8) (len=4) {
        (string) (len=32) "Iħ\u007f\xb3\xd9J6\xban3wEw\x01\x82\"\x95@\x83\x112\xccNF\xe2[\xa5\xc4\xea>)": ([]uint8) (len=35 cap=35) {
            00000000  e2 1b a0 52 8d 52 cb 60  bf 19 05 af f5 a0 5a b8  |...R.R.`......Z.|
            00000010  bb 95 67 4d 15 44 10 5c  be 85 96 7b da eb 39 ee  |..gM.D.\...{..9.|
            00000020  22 2e bd                                          |"..|
        },
        (string) (len=32) "R\x8dR\xcb`\xbf\x19\x05\xaf\xf5\xa0Z\xb8\xbb\x95gM\x15D\x10\\\xbe\x85\x96{\xda\xeb9\xee\".\xbd": ([]uint8) (len=83 cap=83) {
            00000000  f8 51 80 80 80 a0 6d d0  2b 05 08 64 b7 72 a7 bf  |.Q....m.+..d.r..|
            00000010  e4 66 95 ff aa 60 ff 9b  b7 3f 99 66 26 c3 84 01  |.f...`...?.f&...|
            00000020  27 16 3d bb 9b 36 80 80  a0 39 08 c1 b6 37 10 38  |'.=..6...9...7.8|
            00000030  9a 63 e1 e5 c4 a8 61 8f  ce ea 4c 1a 15 ae fd c7  |.c....a...L.....|
            00000040  ef b4 35 40 41 e8 96 6a  3e 80 80 80 80 80 80 80  |..5@A..j>.......|
            00000050  80 80 80                                          |...|
        },
        (string) (len=32) "m\xd0+\x05\bd\xb7r\xa7\xbf\xe4f\x95\xff\xaa`\xff\x9b\xb7?\x99f&Ä\x01'\x16=\xbb\x9b6": ([]uint8) (len=38 cap=38) {
            00000000  e5 80 c4 20 82 76 31 c4  20 82 76 32 c4 20 82 76  |... .v1. .v2. .v|
            00000010  33 c4 20 82 76 34 c4 20  82 76 35 80 80 80 80 80  |3. .v4. .v5.....|
            00000020  80 80 80 80 80 80                                 |......|
        },
        (string) (len=32) "\x01p&e\xbaY/\xf3\xae\x04\xf6R\xc81'UB\x9f#\xe1f\r\x86\x80\xe4\xee\x03\xa5\xaf\xf7^}": ([]uint8) (len=64 cap=64) {
            00000000  f8 3e 80 80 80 80 80 80  a0 49 c4 a7 7f b3 d9 4a  |.>.......I.....J|
            00000010  36 ba 6e 33 77 45 77 01  82 22 95 40 83 11 32 cc  |6.n3wEw..".@..2.|
            00000020  4e 46 e2 5b a5 c4 ea 3e  29 cd 83 34 6d 70 88 62  |NF.[...>)..4mp.b|
            00000030  6c 61 68 62 6c 61 68 80  80 80 80 80 80 80 80 80  |lahblah.........|
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
([]uint8) (len=2 cap=8) {
    00000000  76 32                                             |v2|
}
PASS
ok      _/Users/luke/Desktop/go-wired-blockchain/trie   0.686s
```

# 🤔 Multiple Elements Void Proof

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

```go
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
-->

# 😲 Some Trivial Example

* refer to https://github.com/ethereum/wiki/wiki/Patricia-Tree#example-trie
	* Suppose we want a trie containing four path/value pairs
		* ('do', 'verb')
		* ('dog', 'puppy')
		* ('doge', 'coin')
		* ('horse', 'stallion')

at `trie/proof_test.go`,

```go
func TestMyOwnTestCode(t *testing.T) {
	trie := new(Trie)
	updateString(trie, "do", "verb")
	updateString(trie, "dog", "puppy")
	updateString(trie, "doge", "coin")
	updateString(trie, "horse", "stallion")

	proofs, _ := ethdb.NewMemDatabase()
	trie.Prove([]byte("doge"), 0, proofs)

	spew.Dump(proofs.Len())
	spew.Dump(proofs.Keys())
	spew.Dump(proofs)

	val, err, _ := VerifyProof(trie.Hash(), []byte("doge"), proofs)
	if err != nil {
		t.Fatalf("VerifyProof error: %v\nproof hashes: %v", err, proofs.Keys())
	}

	spew.Dump(val)
}
```

## Outputs

```go
(int) 4
([][]uint8) (len=4 cap=4) {
    ([]uint8) (len=32 cap=32) {
        00000000  59 91 bb 8c 65 14 14 8a  29 db 67 6a 14 ac 50 6c  |Y...e...).gj..Pl|
        00000010  d2 cd 57 75 ac e6 3c 30  a4 fe 45 77 15 e9 ac 84  |..Wu..<0..Ew....|
    },
    ([]uint8) (len=32 cap=32) {
        00000000  bd 3e e5 07 e6 c6 7c fe  fc a9 8f 84 be 47 c1 bb  |.>....|......G..|
        00000010  c0 09 31 5f ab c4 40 5d  b4 ba 32 19 03 74 57 2a  |..1_..@]..2..tW*|
    },
    ([]uint8) (len=32 cap=32) {
        00000000  94 a9 f9 5b d8 96 98 e4  da 18 12 e0 51 80 53 81  |...[........Q.S.|
        00000010  3b 4d 5b 87 ca af 6b 3c  6f a5 7e 9e 50 c0 ff 68  |;M[...k<o.~.P..h|
    },
    ([]uint8) (len=32 cap=32) {
        00000000  d4 3b 87 fd cd 42 17 01  3c cc 92 d0 46 62 e1 2d  |.;...B..<...Fb.-|
        00000010  36 e4 cc 25 dc 69 00 77  cd 82 1a 19 56 fc 3e 36  |6..%.i.w....V.>6|
    }
}
(*ethdb.MemDatabase)(0xc4203a0a80)({
    db: (map[string][]uint8) (len=4) {
        (string) (len=32) "\xd4;\x87\xfd\xcdB\x17\x01<̒\xd0Fb\xe1-6\xe4\xcc%\xdci\x00w͂\x1a\x19V\xfc>6": ([]uint8) (len=52 cap=52) {
            00000000  f3 80 80 80 80 80 80 de  17 dc 80 80 80 80 80 80  |................|
            00000010  c6 35 84 63 6f 69 6e 80  80 80 80 80 80 80 80 80  |.5.coin.........|
            00000020  85 70 75 70 70 79 80 80  80 80 80 80 80 80 80 84  |.puppy..........|
            00000030  76 65 72 62                                       |verb|
        },
        (string) (len=32) "Y\x91\xbb\x8ce\x14\x14\x8a)\xdbgj\x14\xacPl\xd2\xcdWu\xac\xe6<0\xa4\xfeEw\x15鬄": ([]uint8) (len=35 cap=35) {
            00000000  e2 16 a0 bd 3e e5 07 e6  c6 7c fe fc a9 8f 84 be  |....>....|......|
            00000010  47 c1 bb c0 09 31 5f ab  c4 40 5d b4 ba 32 19 03  |G....1_..@]..2..|
            00000020  74 57 2a                                          |tW*|
        },
        (string) (len=32) "\xbd>\xe5\a\xe6\xc6|\xfe\xfc\xa9\x8f\x84\xbeG\xc1\xbb\xc0\t1_\xab\xc4@]\xb4\xba2\x19\x03tW*": ([]uint8) (len=66 cap=66) {
            00000000  f8 40 80 80 80 80 a0 94  a9 f9 5b d8 96 98 e4 da  |.@........[.....|
            00000010  18 12 e0 51 80 53 81 3b  4d 5b 87 ca af 6b 3c 6f  |...Q.S.;M[...k<o|
            00000020  a5 7e 9e 50 c0 ff 68 80  80 80 cf 85 20 6f 72 73  |.~.P..h..... ors|
            00000030  65 88 73 74 61 6c 6c 69  6f 6e 80 80 80 80 80 80  |e.stallion......|
            00000040  80 80                                             |..|
        },
        (string) (len=32) "\x94\xa9\xf9[ؖ\x98\xe4\xda\x18\x12\xe0Q\x80S\x81;M[\x87ʯk<o\xa5~\x9eP\xc0\xffh": ([]uint8) (len=37 cap=37) {
            00000000  e4 82 00 6f a0 d4 3b 87  fd cd 42 17 01 3c cc 92  |...o..;...B..<..|
            00000010  d0 46 62 e1 2d 36 e4 cc  25 dc 69 00 77 cd 82 1a  |.Fb.-6..%.i.w...|
            00000020  19 56 fc 3e 36                                    |.V.>6|
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
([]uint8) (len=4 cap=8) {
    00000000  63 6f 69 6e                                       |coin|
}
PASS
ok      _/Users/luke/Desktop/go-wired-blockchain/trie   0.697s
```

## Analysis

Convert both paths and values to bytes.

Below, actual byte representations for paths are denoted by <>,
although values are still shown as strings, denoted by '', for easier comprehension (they, too, would actually be bytes):

```
<64 6f> : 'verb'
<64 6f 67> : 'puppy'
<64 6f 67 65> : 'coin'
<68 6f 72 73 65> : 'stallion'
```

Now, we build such a trie with the following key/value pairs in the underlying DB:

```
rootHash: [ <16>, hashA ]
hashA:    [ <>, <>, <>, <>, hashB, <>, <>, <>, hashC, <>, <>, <>, <>, <>, <>, <>, <> ]
hashC:    [ <20 6f 72 73 65>, 'stallion' ]
hashB:    [ <00 6f>, hashD ]
hashD:    [ <>, <>, <>, <>, <>, <>, hashE, <>, <>, <>, <>, <>, <>, <>, <>, <>, 'verb' ]
hashE:    [ <17>, hashF ]
hashF:    [ <>, <>, <>, <>, <>, <>, hashG, <>, <>, <>, <>, <>, <>, <>, <>, <>, 'puppy' ]
hashG:    [ <35>, 'coin' ]
```

### Build Trie in Geth

The test code at `trie/trie_test.go`

```go
func TestMyOwnTrie(t *testing.T) {
	trie := newEmpty()
	updateString(trie, "do", "verb")
	updateString(trie, "dog", "puppy")
	updateString(trie, "doge", "coin")
	updateString(trie, "horse", "stallion")

	root, err := trie.Commit()
	if err != nil {
		t.Fatalf("commit error: %v", err)
	}

	spew.Dump(root)
	spew.Dump(trie)
}
```

prints

```go
(common.Hash) (len=32 cap=32) 0x5991bb8c6514148a29db676a14ac506cd2cd5775ace63c30a4fe457715e9ac84
(*trie.Trie)(0xc4207f9680)({
    root: (*trie.shortNode)(0xc4207f99f0)({06: [
    0: <nil> 1: <nil> 2: <nil> 3: <nil> 4: {060f: [
        0: <nil> 1: <nil> 2: <nil> 3: <nil> 4: <nil> 5: <nil> 6: {07: [
            0: <nil> 1: <nil> 2: <nil> 3: <nil> 4: <nil> 5: <nil> 6: {0510: 636f696e } 7: <nil> 8: <nil> 9: <nil> a: <nil> b: <nil> c: <nil> d: <nil> e: <nil> f: <nil> [17]: 7075707079 
          ] } 7: <nil> 8: <nil> 9: <nil> a: <nil> b: <nil> c: <nil> d: <nil> e: <nil> f: <nil> [17]: 76657262 
      ] } 5: <nil> 6: <nil> 7: <nil> 8: {060f07020703060510: 7374616c6c696f6e } 9: <nil> a: <nil> b: <nil> c: <nil> d: <nil> e: <nil> f: <nil> [17]: <nil> 
  ] } ),
    db: (*ethdb.MemDatabase)(0xc4201c3800)({
        db: (map[string][]uint8) (len=4) {
            (string) (len=32) "\x94\xa9\xf9[ؖ\x98\xe4\xda\x18\x12\xe0Q\x80S\x81;M[\x87ʯk<o\xa5~\x9eP\xc0\xffh": ([]uint8) (len=37 cap=37) {
                00000000  e4 82 00 6f a0 d4 3b 87  fd cd 42 17 01 3c cc 92  |...o..;...B..<..|
                00000010  d0 46 62 e1 2d 36 e4 cc  25 dc 69 00 77 cd 82 1a  |.Fb.-6..%.i.w...|
                00000020  19 56 fc 3e 36                                    |.V.>6|
            },
            (string) (len=32) "\xbd>\xe5\a\xe6\xc6|\xfe\xfc\xa9\x8f\x84\xbeG\xc1\xbb\xc0\t1_\xab\xc4@]\xb4\xba2\x19\x03tW*": ([]uint8) (len=66 cap=66) {
                00000000  f8 40 80 80 80 80 a0 94  a9 f9 5b d8 96 98 e4 da  |.@........[.....|
                00000010  18 12 e0 51 80 53 81 3b  4d 5b 87 ca af 6b 3c 6f  |...Q.S.;M[...k<o|
                00000020  a5 7e 9e 50 c0 ff 68 80  80 80 cf 85 20 6f 72 73  |.~.P..h..... ors|
                00000030  65 88 73 74 61 6c 6c 69  6f 6e 80 80 80 80 80 80  |e.stallion......|
                00000040  80 80                                             |..|
            },
            (string) (len=32) "Y\x91\xbb\x8ce\x14\x14\x8a)\xdbgj\x14\xacPl\xd2\xcdWu\xac\xe6<0\xa4\xfeEw\x15鬄": ([]uint8) (len=35 cap=35) {
                00000000  e2 16 a0 bd 3e e5 07 e6  c6 7c fe fc a9 8f 84 be  |....>....|......|
                00000010  47 c1 bb c0 09 31 5f ab  c4 40 5d b4 ba 32 19 03  |G....1_..@]..2..|
                00000020  74 57 2a                                          |tW*|
            },
            (string) (len=32) "\xd4;\x87\xfd\xcdB\x17\x01<̒\xd0Fb\xe1-6\xe4\xcc%\xdci\x00w͂\x1a\x19V\xfc>6": ([]uint8) (len=52 cap=52) {
                00000000  f3 80 80 80 80 80 80 de  17 dc 80 80 80 80 80 80  |................|
                00000010  c6 35 84 63 6f 69 6e 80  80 80 80 80 80 80 80 80  |.5.coin.........|
                00000020  85 70 75 70 70 79 80 80  80 80 80 80 80 80 80 84  |.puppy..........|
                00000030  76 65 72 62                                       |verb|
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
ok      _/Users/luke/Desktop/go-wired-blockchain/trie   0.732s
```

### `proofs.Keys()`

```go
([][]uint8) (len=4 cap=4) {
    ([]uint8) (len=32 cap=32) {
        00000000  59 91 bb 8c 65 14 14 8a  29 db 67 6a 14 ac 50 6c  |Y...e...).gj..Pl|
        00000010  d2 cd 57 75 ac e6 3c 30  a4 fe 45 77 15 e9 ac 84  |..Wu..<0..Ew....|
    },
    ([]uint8) (len=32 cap=32) {
        00000000  bd 3e e5 07 e6 c6 7c fe  fc a9 8f 84 be 47 c1 bb  |.>....|......G..|
        00000010  c0 09 31 5f ab c4 40 5d  b4 ba 32 19 03 74 57 2a  |..1_..@]..2..tW*|
    },
    ([]uint8) (len=32 cap=32) {
        00000000  94 a9 f9 5b d8 96 98 e4  da 18 12 e0 51 80 53 81  |...[........Q.S.|
        00000010  3b 4d 5b 87 ca af 6b 3c  6f a5 7e 9e 50 c0 ff 68  |;M[...k<o.~.P..h|
    },
    ([]uint8) (len=32 cap=32) {
        00000000  d4 3b 87 fd cd 42 17 01  3c cc 92 d0 46 62 e1 2d  |.;...B..<...Fb.-|
        00000010  36 e4 cc 25 dc 69 00 77  cd 82 1a 19 56 fc 3e 36  |6..%.i.w....V.>6|
    }
}
```

1. rootHash = `<59 91 bb 8c 65 14 14 8a 29 db 67 6a 14 ac 50 6c d2 cd 57 75 ac e6 3c 30 a4 fe 45 77 15 e9 ac 84>`
2. 
3.
4.

### `proofs.Len()`

```go
(int) 4
```

Trivial.

### `proofs`

TBA

## `val`

TBA
