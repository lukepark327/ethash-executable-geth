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

## 새 계좌 생성
```bash
$ build/bin/geth --datadir ./mydata/ account new
```

### 계좌 생성 확인
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
