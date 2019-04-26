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
        "0x5066f597D21F2aD5976F7bcA1B89cdf74052a4Cf": { "balance": "60000000000000000000" }
    }
}
```

## Geth Command

```bash
miner.start()
rawpow.echoNumber(5)
```

## Get Geth Version
```bash
build/bin/geth version
```
