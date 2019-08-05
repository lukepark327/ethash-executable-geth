`internal/ethapi/api.go`

`SendTxArgs` represents the arguments to sumbit a new transaction into the transaction pool.

```go
type SendTxArgs struct {
	From     common.Address  `json:"from"`
	To       *common.Address `json:"to"`
	Gas      *hexutil.Uint64 `json:"gas"`
	GasPrice *hexutil.Big    `json:"gasPrice"`
	Value    *hexutil.Big    `json:"value"`
	Nonce    *hexutil.Uint64 `json:"nonce"`
	// We accept "data" and "input" for backwards-compatibility reasons. "input" is the
	// newer name and should be preferred by clients.
	Data  *hexutil.Bytes `json:"data"`
	Input *hexutil.Bytes `json:"input"`
}
```

# DelegatedFrom

```go
type SendTxArgs struct {
	From          common.Address  `json:"from"`
	DelegatedFrom *common.Address `json:"delegatedFrom"`
	To            *common.Address `json:"to"`
	Gas           *hexutil.Uint64 `json:"gas"`
	GasPrice      *hexutil.Big    `json:"gasPrice"`
	Value         *hexutil.Big    `json:"value"`
	Nonce         *hexutil.Uint64 `json:"nonce"`
	// We accept "data" and "input" for backwards-compatibility reasons. "input" is the
	// newer name and should be preferred by clients.
	Data  *hexutil.Bytes `json:"data"`
	Input *hexutil.Bytes `json:"input"`
}
```

## setDefaults

setDefaults is a helper function that fills in default values for unspecified tx fields.

```go
func (args *SendTxArgs) setDefaults(ctx context.Context, b Backend) error {
	if args.DelegatedFrom == nil {
		args.DelegatedFrom = &args.From
	}

  // ...

	return nil
}
```

# Using Input Field

`addr = common.BytesToAddress(tx.data.Payload)`
