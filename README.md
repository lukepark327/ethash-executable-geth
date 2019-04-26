# geth_breakdown

## Set a Genesis Block

```bash
$ mkdir mydata/genesis
$ vi mydata/genesis/rawpow.json
```

first.json:
```JSON
{
    "config": {
        "chainId": 950327,
        "homesteadBlock": 0,
        "eip155Block": 0,
        "eip158Block": 0,
        "rawpow": {}
    },
    "difficulty": "20",
    "gasLimit": "2100000",
    "alloc": {
        "29da13e995efac76e10f36e40f4736c64ebf0048": { "balance": "60000000000000000000" }
    }
}
```

## Geth Command

```bash
miner.start()
rawpow.echoNumber(5)
```
