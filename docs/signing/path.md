# Path

How transactions are handled? What variables and functions handle the `eth.sendTransaction` request?

# api.go

`internal/ethapi/api.go`

## SendTransaction()

`SendTransaction` will create a transaction from the given arguments and tries to sign it with the key associated with `args.To`.
If the given passwd isn't able to decrypt the key it fails.

`SendTransaction` uses several variables and functions;

### wallet

`wallet` is used to create signed Tx. Using private key associated with `From`.

### setDefaults()

`setDefaults` is a helper function that fills in default values for unspecified tx fields.

### toTransaction()

`tx := args.toTransaction()`

Setting receipt address with `args.To`.

### SubmitTransaction()

Finally, `SendTransaction` calls `SubmitTransaction`.
