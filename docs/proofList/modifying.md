In `core/state/statedb.go`,

`
type ProofList [][]byte
`

is defined with two methods;

* `Put` and `Delete`

## Adding Methods

We add `Has` and `Get` methods to use proofs;

### Has

```go
func (n *ProofList) Has(key []byte) (bool, error) {
	panic("not supported")
}
```

Actually we do not need `Has`...

### Get

```go
func (n *ProofList) Get(key []byte) ([]byte, error) {
	x := (*n)[0]
	*n = (*n)[1:]
	return x, nil
}
```

Now we can use `proof`s.
