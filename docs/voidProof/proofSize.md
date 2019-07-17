# Tries in Ethereum

## State Trie

* There is one global state trie, and it updates over time.

* In it, a path is always: `sha3(ethereumAddress)` and a value is always: `rlp(ethereumAccount)`.

  * More specifically an ethereum account is a 4 item array of [nonce,balance,storageRoot,codeHash].
  
    * At this point it's worth noting that this storageRoot is the root of another patricia trie:

# References

* https://github.com/ethereum/wiki/wiki/Patricia-Tree
