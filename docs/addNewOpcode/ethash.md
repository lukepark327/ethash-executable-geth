## Hardcoding Factors
In `vendor/github.com/ethereum/ethash/src/libethash/data_sizes.h`
* dag_sizes
  ```go
  static const uint64_t dag_sizes[2048] = {
    1073739904U, 1082130304U, 1090514816U, 1098906752U, 1107293056U,
    1115684224U, 1124070016U, 1132461952U, 1140849536U, 1149232768U,
    ...,
    18228444544U, 18236833408U, 18245220736U
  };
  ```
* cache_sizes
  ```go
  const uint64_t cache_sizes[2048] = {
    16776896U, 16907456U, 17039296U, 17170112U, 17301056U, 17432512U, 17563072U,
    17693888U, 17824192U, 17955904U, 18087488U, 18218176U, 18349504U, 18481088U,
    284950208U, 285081536U
  };
  ```
We regard those two factors are fixed.

## VerifySeal
VerifySeal implements `consensus/ethash/consensus.go`, checking whether the given block satisfies the PoW difficulty requirements.
```go
func (ethash *Ethash) VerifySeal(chain consensus.ChainReader, header *types.Header) error {
	...
	// Ensure that we have a valid difficulty for the block
	if header.Difficulty.Sign() <= 0 {
		return errInvalidDifficulty
	}
	// Recompute the digest and PoW value and verify against the header
	number := header.Number.Uint64()

	cache := ethash.cache(number)
	size := datasetSize(number)
	if ethash.config.PowMode == ModeTest {
		size = 32 * 1024
	}
	digest, result := hashimotoLight(size, cache.cache, header.HashNoNonce().Bytes(), header.Nonce.Uint64())
	// Caches are unmapped in a finalizer. Ensure that the cache stays live
	// until after the call to hashimotoLight so it's not unmapped while being used.
	runtime.KeepAlive(cache)

	if !bytes.Equal(header.MixDigest[:], digest) {
		return errInvalidMixDigest
	}
	target := new(big.Int).Div(maxUint256, header.Difficulty)
	if new(big.Int).SetBytes(result).Cmp(target) > 0 {
		return errInvalidPoW
	}
	return nil
}
```

## Use Engine for Calculating and Verifying Ethash
```go
otherEngine := ethash.New(ethash.Config{
	CacheDir:       ~,
	CachesInMem:    ~,
	CachesOnDisk:   ~,
	DatasetDir:     ~,
	DatasetsInMem:  ~,
	DatasetsOnDisk: ~,
})
```
Config's structure is:
```go
type Config struct {
	CacheDir       string
	CachesInMem    int
	CachesOnDisk   int
	DatasetDir     string
	DatasetsInMem  int
	DatasetsOnDisk int
	PowMode        Mode
}
```
and its default value is
```go
Ethash: ethash.Config{
	CacheDir:       "ethash",
	CachesInMem:    2,
	CachesOnDisk:   3,
	DatasetsInMem:  1,
	DatasetsOnDisk: 2,
},
```
also
```go
	PoWMode: 	ModeNormal
```
But default `DatasetDir` needs some calculation;
```go
func init() {
	home := os.Getenv("HOME")
	if home == "" {
		if user, err := user.Current(); err == nil {
			home = user.HomeDir
		}
	}
	if runtime.GOOS == "darwin" {
		DefaultConfig.Ethash.DatasetDir = filepath.Join(home, "Library", "Ethash")
	} else if runtime.GOOS == "windows" {
		localappdata := os.Getenv("LOCALAPPDATA")
		if localappdata != "" {
			DefaultConfig.Ethash.DatasetDir = filepath.Join(localappdata, "Ethash")
		} else {
			DefaultConfig.Ethash.DatasetDir = filepath.Join(home, "AppData", "Local", "Ethash")
		}
	} else {
		DefaultConfig.Ethash.DatasetDir = filepath.Join(home, ".ethash")
	}
}
```
Refer to `eth/config.go` for details.

## References
* https://github.com/ethereum/wiki/wiki/Ethash#defining-the-seed-hash
