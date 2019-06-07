# Seal

In `./consensus/ethash/sealer.go`,

`Seal` implements consensus.Engine, attempting to find a nonce that satisfies the block's difficulty requirements.

```go
if threads == 0 {
  threads = runtime.NumCPU()
}
if threads < 0 {
  threads = 0 // Allows disabling local mining without extra logic around local/remote
}
var pend sync.WaitGroup
for i := 0; i < threads; i++ {
  pend.Add(1)
  go func(id int, nonce uint64) {
    defer pend.Done()
    ethash.mine(block, id, nonce, abort, found)
  }(i, uint64(ethash.rand.Int63()))
}
```

* Call `mine` function.

* `uint64(ethash.rand.Int63())` is properly seeded random source for nonces.

## mine

`mine` is the actual proof-of-work miner that searches for a nonce starting from seed that results in correct final block difficulty.

```go
func (ethash *Ethash) mine(block *types.Block, id int, seed uint64, abort chan struct{}, found chan *types.Block) {
  ...
}
```

```go
// Extract some data from the header
var (
  header  = block.Header()
  hash    = header.HashNoNonce().Bytes()
  target  = new(big.Int).Div(maxUint256, header.Difficulty)
  // number  = header.Number.Uint64()
  // dataset = ethash.dataset(number)
)
// Start generating random nonces until we abort or find a good one
var (
  attempts = int64(0)
  nonce    = seed
)
```

## Compute the PoW value of this nonce

```go
digest, result := hashimotoFull(dataset.dataset, hash, nonce)
```

will be changed to

```go
func naivePow(hash []byte, nonce uint64) ([]byte, []byte) {
	// rows := uint32(size / mixBytes)

	// Combine header+nonce into a 64 byte seed
	seed := make([]byte, 40)
	copy(seed, hash)
	binary.LittleEndian.PutUint64(seed[32:], nonce)

	seed = crypto.Keccak512(seed)
	// seedHead := binary.LittleEndian.Uint32(seed)

	// Start the mix with replicated seed
	mix := make([]uint32, mixBytes/4)
	for i := 0; i < len(mix); i++ {
		mix[i] = binary.LittleEndian.Uint32(seed[i%16*4:])
	}
	// Mix in random dataset nodes
	// temp := make([]uint32, len(mix))

	// for i := 0; i < loopAccesses; i++ {
	// 	parent := fnv(uint32(i)^seedHead, mix[i%len(mix)]) % rows
	// 	for j := uint32(0); j < mixBytes/hashBytes; j++ {
	// 		copy(temp[j*hashWords:], lookup(2*parent+j))
	// 	}
	// 	fnvHash(mix, temp)
	// }
	// Compress mix
	for i := 0; i < len(mix); i += 4 {
		mix[i/4] = fnv(fnv(fnv(mix[i], mix[i+1]), mix[i+2]), mix[i+3])
	}
	mix = mix[:len(mix)/4]

	digest := make([]byte, common.HashLength)
	for i, val := range mix {
		binary.LittleEndian.PutUint32(digest[i*4:], val)
	}
	return digest, crypto.Keccak256(append(seed, digest...))
}
```

Refer `hashimoto` function.

## hashimoto

In `./go-ethereum/consensus/ethash/algorithm.go`,

`hashimoto` aggregates data from the full dataset in order to produce our final value for a particular header hash and nonce.

# VerifySeal

In `./consensus/ethash/consensus.go`,

`VerifySeal` implements consensus.Engine, checking whether the given block satisfies the PoW difficulty requirements.

Use `naivePow` function instead of `HashimotoLight`.

```go
digest, result := naivePow(header.HashNoNonce().Bytes(), header.Nonce.Uint64())
```
