# `trie/proof.go`

## `func (t *Trie) Prove`

* Prove constructs a merkle proof for key.

* The result contains all encoded nodes on the path to the value at key.

  * The value itself is also included in the last node and can be retrieved by verifying the proof.

* If the trie does not contain a value for key,

  * the returned proof contains all nodes of the longest existing prefix of the key (at least the root node),
  * ending with the node that **proves the absence of the key.**




