# Overview
TBA

# First: Create My Own Consensus

## Set a Genesis Block

```bash
$ mkdir mydata/genesis
$ vi mydata/genesis/rawpow.json
```

rawpow.json:
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
    "gasLimit": "7000000",
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
version
Geth
Version: 1.8.12-unstable
Architecture: amd64
Protocol Versions: [63 62]
Network Id: 1
Go Version: go1.9.7
Operating System: darwin
GOPATH=
GOROOT=/usr/local/Cellar/go@1.9/1.9.7/libexec
```
Latest Geth Version is 'Punisher (v1.8.27)' (Apr. 26th, 2019)

# Second: Modify EVM for Adding New OPCODEs
TBA
