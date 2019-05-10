# Overview
* geth tutorial: https://github.com/twodude/ghost-relay/wiki

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

## Run geth
```bash
$ build/bin/geth --datadir ./mydata/ --networkid 950327 --port 32222 --rpc --rpcport 8222 --nodiscover console
```

## Deploy Contract
in geth console,
```bash
> loadScript("/Users/luke/Desktop/solidity/solexam/testEthash.js")
true
> testOutput
{
  contracts: {
    solexam/testEthash.sol:testEthash: {
      abi: "[{\"constant\":true,\"inputs\":[],\"name\":\"getEthash\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]",
      bin: "6080604052348015600f57600080fd5b5060067f30783132000000000000000000000000000000000000000000000000000000006003466000806101000a81548160ff021916908315150217905550608f8061005c6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c806303c6f5b314602d575b600080fd5b6033604d565b604051808215151515815260200191505060405180910390f35b60008060009054906101000a900460ff1690509056fea165627a7a72305820843ff4b78ea01dc4651a5bbe25a9c1c382fc50689719c0edcc25913df3bbc7390029"
    }
  },
  version: "0.5.9-develop.2019.5.7+commit.0fcb3e85.mod.Darwin.appleclang"
}
```
You can check contract bytecodes;
```bash
> testOutput.contracts
{
  solexam/testEthash.sol:testEthash: {
    abi: "[{\"constant\":true,\"inputs\":[],\"name\":\"getEthash\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]",
    bin: "6080604052348015600f57600080fd5b5060067f30783132000000000000000000000000000000000000000000000000000000006003466000806101000a81548160ff021916908315150217905550608f8061005c6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c806303c6f5b314602d575b600080fd5b6033604d565b604051808215151515815260200191505060405180910390f35b60008060009054906101000a900460ff1690509056fea165627a7a72305820843ff4b78ea01dc4651a5bbe25a9c1c382fc50689719c0edcc25913df3bbc7390029"
  }
}
```

Deploy by user whose index number is zero.

```bash
> var testContract = web3.eth.contract(JSON.parse(testOutput.contracts["solexam/testEthash.sol:testEthash"].abi));
undefined
> personal.unlockAccount(eth.accounts[0], "12341234");
true
> var test = testContract.new({ from: eth.accounts[0], data: "0x" + testOutput.contracts["solexam/testEthash.sol:testEthash"].bin, gas: 2000000},
  function (e, contract) {
      console.log(e, contract);
      if (typeof contract.address !== 'undefined') {
           console.log('Contract mined! address: ' + contract.address + ' transactionHash: ' + contract.transactionHash);
      }
    }
  );
INFO [05-07|17:15:14.598917] Submitted contract creation              fullhash=0x338c4ae7acf8006e522e7c23bebf82a88fd0aa275b2b5664803b035121007bde contract=0x37b601a8d2367CB5962DD3D67d6Dd9c36F0d8040
null [object Object]
undefined
```
where "12341234" is a password.

## Start mining
```bash
> miner.start()
INFO [05-07|17:15:33.802993] Transaction pool price threshold updated price=18000000000
INFO [05-07|17:15:33.803083] Starting mining operation
null
> INFO [05-07|17:15:33.805907] will Finalize the block
INFO [05-07|17:15:33.806014] Commit new mining work                   number=3 txs=1 uncles=0 elapsed=2.877ms
INFO [05-07|17:15:33.806132] will Seal the block
hash is : 0x2e52b73bd6918b29b96c932a4e3f4722194f5c0b077e4e06ef29a275d9edb10aINFO [05-07|17:15:48.807011] Successfully sealed new block            number=3 hash=33258e…1bd400
INFO [05-07|17:15:48.810284] 🔨 mined potential block                  number=3 hash=33258e…1bd400
INFO [05-07|17:15:48.811704] will Finalize the block
INFO [05-07|17:15:48.811775] Commit new mining work                   number=4 txs=0 uncles=0 elapsed=1.429ms
INFO [05-07|17:15:48.811839] will Seal the block
null [object Object]
Contract mined! address: 0x37b601a8d2367cb5962dd3d67d6dd9c36f0d8040 transactionHash: 0x338c4ae7acf8006e522e7c23bebf82a88fd0aa275b2b5664803b035121007bde
> miner.stop()
true
```

Now you can find Contract's address and deploying transaction's hash. Then,
```bash
> eth.getTransaction("0x338c4ae7acf8006e522e7c23bebf82a88fd0aa275b2b5664803b035121007bde");
{
  blockHash: "0x33258ef216148d62863cdbc25e03a48722ef473c997ed391068e2b40081bd400",
  blockNumber: 3,
  from: "0x6282ad5f86c03726722ec397844d2f87ced3af89",
  gas: 2000000,
  gasPrice: 18000000000,
  hash: "0x338c4ae7acf8006e522e7c23bebf82a88fd0aa275b2b5664803b035121007bde",
  input: "0x6080604052348015600f57600080fd5b5060067f30783132000000000000000000000000000000000000000000000000000000006003466000806101000a81548160ff021916908315150217905550608f8061005c6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c806303c6f5b314602d575b600080fd5b6033604d565b604051808215151515815260200191505060405180910390f35b60008060009054906101000a900460ff1690509056fea165627a7a72305820843ff4b78ea01dc4651a5bbe25a9c1c382fc50689719c0edcc25913df3bbc7390029",
  nonce: 1,
  r: "0x4961a7b418345a946f35556c86f9405d9b90d4b363cb001dd66eb706e09a825f",
  s: "0x5dd62a6bd1340ce221076eb75f617232556e69cd02836f017b8ac97c99acba5b",
  to: null,
  transactionIndex: 0,
  v: "0x1d0091",
  value: 0
}
> eth.getBlock(3)
{
  difficulty: 131072,
  extraData: "0xd88301080c846765746887676f312e392e378664617277696e",
  gasLimit: 6982443,
  gasUsed: 115330,
  hash: "0x33258ef216148d62863cdbc25e03a48722ef473c997ed391068e2b40081bd400",
  logsBloom: "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
  miner: "0x6282ad5f86c03726722ec397844d2f87ced3af89",
  mixHash: "0x0000000000000000000000000000000000000000000000000000000000000000",
  nonce: "0x000000000000000a",
  number: 3,
  parentHash: "0x2e52b73bd6918b29b96c932a4e3f4722194f5c0b077e4e06ef29a275d9edb10a",
  receiptsRoot: "0x44dc86364dde328180685b0c72761aff2e857a68402401776503cb24472c7184",
  sha3Uncles: "0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347",
  size: 864,
  stateRoot: "0x3962796cac67bce90ce4a0e7b341be4077dcf979d2a0dafd7f34222fc75d42f8",
  timestamp: 1557216933,
  totalDifficulty: 393236,
  transactions: ["0x338c4ae7acf8006e522e7c23bebf82a88fd0aa275b2b5664803b035121007bde"],
  transactionsRoot: "0x3cb67b954d1cee407ee8abc2d9846e59aa188eab0cb757356149990f16efcee7",
  uncles: []
}
```

## Call the Contract Function
```bash
> web3.sha3("getEthash()")
"0x03c6f5b32fdeb6e00f1d2aef938622d89c2c99b6223d8bff960ba907f006e245"
```
So the function selector for the getNum() function is 0x03c6f5b3.

```bash
> eth.sendTransaction({from:eth.accounts[0], to:"0x37b601a8d2367cb5962dd3d67d6dd9c36f0d8040", value:0, data:"0x03c6f5b30000000000000000000000000000000000000000000000000000000000000000"})
INFO [05-07|17:21:54.413529] Submitted transaction                    fullhash=0xf6c8654293b7d7285370d531a404adbc3df0e59fc47a18fa700891a4dc2e1208 recipient=0x37b601a8d2367CB5962DD3D67d6Dd9c36F0d8040
"0xf6c8654293b7d7285370d531a404adbc3df0e59fc47a18fa700891a4dc2e1208"
```
* from: account of the caller. 
* to: contract address 
* value: since the purpose of this call is not to transfer money, this value is 0.
* data: describes the function to call and what parameters to use.
  * The first four bytes are the function selector. This is to call the getNum() function so 0x67e0badb.
  * Since the set function has no parameter,
    * The word length is 32 bytes.
    * The (parameter's) value set here is 0x0000 because there is no ipnut params. Actually, I think that I show you a wrong example... Anyway,
    
```bash
> eth.getStorageAt("0x37b601a8d2367cb5962dd3d67d6dd9c36f0d8040",0);
"0x0000000000000000000000000000000000000000000000000000000000000001"
```
because 'true' is same as '1'.

### Use Method-Arguments Scheme
For example, there are two methods;
```bash
> cat solexam/testEthash.abi | python -m json.tool
[
    {
        "constant": true,
        "inputs": [],
        "name": "getEthash",
        "outputs": [
            {
                "name": "",
                "type": "bool"
            }
        ],
        "payable": false,
        "stateMutability": "view",
        "type": "function"
    },
    {
        "constant": false,
        "inputs": [
            {
                "name": "blockNumber",
                "type": "uint256"
            },
            {
                "name": "mixDigest",
                "type": "bytes32"
            },
            {
                "name": "hashNoNonce",
                "type": "bytes32"
            },
            {
                "name": "difficulty",
                "type": "uint256"
            },
            {
                "name": "nonce",
                "type": "uint256"
            }
        ],
        "name": "setEthash",
        "outputs": [
            {
                "name": "",
                "type": "bool"
            }
        ],
        "payable": false,
        "stateMutability": "nonpayable",
        "type": "function"
    }
]
```
You can call those method by the following instructions;
```bash
> test.setEthash(3, "0x12", "0x34", 6, 1234, {from: eth.accounts[0], gas: 500000})
INFO [05-10|17:45:18.799955] Submitted transaction                    fullhash=0xdfde98cecc4f5f310df2082921e6600eeb126da69935c1cc10a15f4b10ea9cd3 recipient=0x37b601a8d2367CB5962DD3D67d6Dd9c36F0d8040
```
and
```bash
> test.getEthash()
false
```

# Trouble Shootings

## Fail to Deploy
```bash
Error: The contract code couldn't be stored, please check your gas amount. undefined
```
If the above error occurs, you might need to allocate more gas. If the error doesn't disappear with additional gas, make sure that you modify EVM to treat a new OPCODE.

# References

[1] https://blog.csdn.net/weixin_40401264/article/details/78136346   
