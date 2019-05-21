# Research

## MixDigest

Mixhash is actually calculated from nonce as intermediate value when validating PoW with Hashimoto algorithm. But this calculation is still pretty heavy and a node might be DDoSed by blocks with incorrect nonces. mixhash is included into block to perform lightweight PoW 'pre-validation' to avoid such attack, as generating a correct mixhash still requires at least some work for attacker.

## Calculate Hash
```go
// Hash returns the block hash of the header, which is simply the keccak256 hash of its
// RLP encoding.
func (h *Header) Hash() common.Hash {
	return rlpHash(h)
}

func rlpHash(x interface{}) (h common.Hash) {
	hw := sha3.NewKeccak256()
	rlp.Encode(hw, x)
	hw.Sum(h[:0])
	return h
}
```
Use all block header's info with a keccak256 hash function.

# References

[1] https://ethereum.stackexchange.com/questions/5833/why-do-we-need-both-nonce-and-mixhash-values-in-a-block   
