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
	digest := make([]byte, 32)
	copy(digest, hash)

	// Combine header+nonce into a 64 byte seed
	seed := make([]byte, 40)
	copy(seed, hash)
	binary.LittleEndian.PutUint64(seed[32:], nonce)

	return digest, crypto.Keccak256(append(seed, digest...))
}

func (littleEndian) PutUint64(b []byte, v uint64) {
	_ = b[7] // early bounds check to guarantee safety of writes below
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
	b[4] = byte(v >> 32)
	b[5] = byte(v >> 40)
	b[6] = byte(v >> 48)
	b[7] = byte(v >> 56)
}
```

Now digest is just `header.HashNoNonce().Bytes()`.

# VerifySeal

In `./consensus/ethash/consensus.go`,

`VerifySeal` implements consensus.Engine, checking whether the given block satisfies the PoW difficulty requirements.

Use `naivePow` function instead of `HashimotoLight`.

```go
digest, result := naivePow(header.HashNoNonce().Bytes(), header.Nonce.Uint64())

if !bytes.Equal(header.MixDigest[:], digest) {
	return errInvalidMixDigest
}
target := new(big.Int).Div(maxUint256, header.Difficulty)
if new(big.Int).SetBytes(result).Cmp(target) > 0 {
	return errInvalidPoW
}
return nil
```

# References

Refer Bitcoin's PoW. See https://github.com/bitcoin/bitcoin/blob/master/src/pow.cpp#L74

```cpp
bool CheckProofOfWork(uint256 hash, unsigned int nBits, const Consensus::Params& params)
{
    bool fNegative;
    bool fOverflow;
    arith_uint256 bnTarget;

    bnTarget.SetCompact(nBits, &fNegative, &fOverflow);

    // Check range
    if (fNegative || bnTarget == 0 || fOverflow || bnTarget > UintToArith256(params.powLimit))
        return false;

    // Check proof of work matches claimed amount
    if (UintToArith256(hash) > bnTarget)
        return false;

    return true;
}
```
