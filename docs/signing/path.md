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

### signTransaction()

`signed, err := s.signTransaction(ctx, &args, passwd)`

`signTransction` calls `setDefaults` and `toTranscation`.

### setDefaults()

`setDefaults` is a helper function that fills in default values for unspecified tx fields.

### toTransaction()

`tx := args.toTransaction()`

Setting receipt address with `args.To`.

`signed, err := wallet.SignTx(account, tx, s.b.ChainConfig().ChainID)`

### signed

`signed`'s type is `Transaction`.

```go
type Transaction struct {
	data txdata
	// caches
	hash atomic.Value
	size atomic.Value
	from atomic.Value
}
```

*Set `from` here?*

## SubmitTransaction()

Finally, `SendTransaction` calls `SubmitTransaction`. `SubmitTransaction` is a helper function that submits tx to txPool and logs a message.

`SubmitTransaction` uses several variables and functions;

### SendTx()

`eth/api_backend.go`

`SendTx` calls `AddLocal`.

### AddLocal()

`core/tx_pool.go`

`AddLocal` calls `AddLocals`.

### AddLocals()

`addTxs` calls `addTxs`.

### addTxs()

`addTxs` attempts to queue a batch of transactions if they are valid.

`addTxs` calls `addTxsLocked` and `requestPromoteExecutables`.

*Remark*

```go
// Cache senders in transactions before obtaining lock (pool.signer is immutable)
for _, tx := range txs {
	types.Sender(pool.signer, tx)
}
```

Now `tx`'s `from` field contains some values...

```go
INFO [08-03|15:46:43.471] [eth4nos] addTxs                         i=0 tx="&{data:{AccountNonce:0 Price:0xc46b77e860 GasLimit:21000 Recipient:0xc46b6bddc0 Amount:0xc46b77e840 Payload:[] V:0xc46b77ea00 R:0xc46b77e9a0 S:0xc46b77e9c0 Hash:<nil>} hash:{noCopy:{} v:<nil>} size:{noCopy:{} v:<nil>} from:{noCopy:{} v:{signer:{chainId:0xc420556660 chainIdMul:0xc4203de7e0} from:[196 66 45 28 24 233 234 216 169 187 152 235 13 139 185 219 223 40 17 215]}}}"
```

### addTxsLocked()

`addTxsLocked` calls `add` and `addTx`.

### add()

### addTx()

### requestPromoteExecutables()

`requestPromoteExecutables` requests transaction promotion checks for the given addresses. The returned channel is closed when the promotion checks have occurred.
