# How to Use JSON RPC

```bash
$ curl -X POST --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' https://mainnet.infura.io:443

{"jsonrpc":"2.0","id":1,"result":"0x7cb26c"}%
```

* Use `infura`

# web3.js in go-ethereum

* In `internal/jsre/deps/web3.js`.

* There are some js file's binary (bytecode) datas in `internal/jsre/deps/bindata.go`.

* Running binary with `c.jsre.Run('[codes]')` like:

```go
// Print some generic Geth metadata
if res, err := c.jsre.Run(`
  var message = "instance: " + web3.version.node + "\n";
  try {
    message += "coinbase: " + eth.coinbase + "\n";
  } catch (err) {}
  message += "at block: " + eth.blockNumber + " (" + new Date(1000 * eth.getBlock(eth.blockNumber).timestamp) + ")\n";
  try {
    message += " datadir: " + admin.datadir + "\n";
  } catch (err) {}
  message
`); err == nil {
  message += res.String()
}
```

in `console/console.go`.

# Conclusion

Creating a new js code instead of modifying a current web3.js.

# References

* https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_blocknumber
