# Install and Run Geth

## 수정된 geth 다운로드

* git을 이용해 다음 경로에서 다운로드 `https://github.com/twodude/geth-breakdown.git`

* Build
```bash
$ cd geth-breakdown/go-ethereum
$ make geth
```

## Genesis Block 설정

* 다음 파일을 geth-breakdown/go-ethereum/mydata/genesis/default.json 경로에 작성

```json
{
    "config": {
        "chainId": 950327,
        "homesteadBlock": 0,
        "eip155Block": 0,
        "eip158Block": 0,
        "eip160Block": 0,
        "byzantiumBlock": 0,
        "ethash": {}
    },
    "difficulty": "20000000",
    "gasLimit": "7000000",
    "alloc": {
        "6282ad5f86c03726722ec397844d2f87ced3af89": {
            "balance": "60000000000000000000"
        }
    }
}
```

* Set Genesis Block

```bash
$ touch mydata/genesis/default.json
$ build/bin/geth --datadir ./mydata/ init ./mydata/genesis/default.json
```

## 새 account 생성
```bash
$ build/bin/geth --datadir ./mydata/ account new
```

* 다음 명령으로 account가 잘 생성됐는지 확인
```bash
$ build/bin/geth --datadir ./mydata/ account list
```

## 수정된 geth 실행

```bash
$ build/bin/geth console --datadir ./mydata/ --networkid 950327 --port 32222 --rpc --rpcport "8002" --rpcaddr "0.0.0.0" --rpccorsdomain "*" --rpcapi db,eth,net,web3,personal --nodiscover
```

* 이제 `http://localhost:8002` or `http://"IP주소":8002`로 접근 가능

# Geth Console

## 채굴 시작하기

```bash
> miner.start()
```

## 채굴 중단하기

```bash
> miner.stop()
```

# Run Relayer

## Relay 다운로드

* git을 이용해 다음 경로에서 다운로드 `https://github.com/twodude/eth-proof-sol.git`

## Private chain -> Testnet(Ethereum) chain

* Compile Smart Contract

```bash
$ sh naivePoW/deps.sh
```

* Set config.json
    * 다음 파일을 geth-breakdown/go-ethereum/mydata/genesis/default.json 경로에 작성
    
```json
{
  "owner": {
    "address": "사용할 ethereum 계좌 주소",
    "password": "geth에서 설정한 비밀번호"
  }
}
```

* `genesis.js` 실행
```bash
$ node naivePoW/genesis.js
```

* `contractDeploy.js` 실행
```bash
$ node naivePoW/contractDeploy.js
```

### Relay 구동

```bash
$ node naivePoW/relay.js 
```

### 트랜잭션 검증하기

```bash
$ node naivePow/txProof.js [0x8b68b49ea234880ea061803aae2322ba3ac57a2ae8a5feac4e13a4b3f67622f1]
```

`txProof.js` 구동 시 인자로 검증하고자 하는 트랜잭션 ID를 명시

* 올바른 트랜잭션일 경우:
```bash
$ node naivePoW/txProof.js 0x8b68b49ea234880ea061803aae2322ba3ac57a2ae8a5feac4e13a4b3f67622f1
> Unlocking an account
> checkTxProof: true
```

* 올바르지 않은 트랜잭션일 경우:
```bash
node naivePoW/txProof.js 0x8b68b49ea234880ea061803aae2322ba3ac57a2ae8a5feac4e13a4b3f67622f2
> Unlocking an account
node:26813) UnhandledPromiseRejectionWarning: transaction not found
(node:26813) UnhandledPromiseRejectionWarning: Unhandled promise rejection. This error originated either by throwing inside of an async function without a catch block, or by rejecting a promise which was not handled with .catch(). (rejection id: 1)
(node:26813) [DEP0018] DeprecationWarning: Unhandled promise rejections are deprecated. In the future, promise rejections that are not handled will terminate the Node.js process with a non-zero exit code.
```
