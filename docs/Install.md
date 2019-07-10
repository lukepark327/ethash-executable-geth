# 수정된 geth

* EVM(이더리움 가상 머신)에서 새로운 OPCODE인 `ethash`를 처리할 수 있도록 수정한 geth 클라이언트. `ethash`는 ETHASH라는 이더리움의 memory-hard-Proof-of-Work를 통해 형성된 블록 헤더를 검증할 수 있는 명령어이다.

* 사용법 Simple Solidity Example: [validEthashTest.md](https://github.com/twodude/geth-breakdown/blob/master/docs/addNewOpcode/validEthashTest.md) 참조.

## 수정된 geth 설치

### Download
git을 이용해 다음 경로에서 다운로드 `https://github.com/twodude/geth-breakdown.git`

### Build
프로그래밍 언어 `Go`를 우선 설치해야 함.

```bash
$ cd geth-breakdown/go-ethereum
$ make geth
```

## Genesis Block 설정
다음 파일을 geth-breakdown/go-ethereum/mydata/genesis/default.json 경로에 작성

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

### Set Genesis Block

```bash
$ touch mydata/genesis/default.json
$ build/bin/geth --datadir ./mydata/ init ./mydata/genesis/default.json
```

## 새 account 생성
```bash
$ build/bin/geth --datadir ./mydata/ account new
```

### 생성 확인
다음 명령으로 account가 잘 생성됐는지 확인
```bash
$ build/bin/geth --datadir ./mydata/ account list
```

## 수정된 geth 실행
```bash
$ build/bin/geth console --datadir ./mydata/ --networkid 950327 --port 32222 --rpc --rpcport "8002" --rpcaddr "0.0.0.0" --rpccorsdomain "*" --rpcapi db,eth,net,web3,personal --nodiscover
```

`http://localhost:8002` or `http://"IP주소":8002`로 수정된 geth에 접근할 수 있다.

# Geth Console

## 채굴 시작하기

```bash
> miner.start()
```

## 채굴 중단하기

```bash
> miner.stop()
```

# 수정된 Solidity Compiler

* 추가된 OPCODE인 `ethash`를 포함한 스마트 컨트랙트를 compile할 수 있는 Solidity Compiler.

## 수정된 Solidity Compiler 설치

### Download
git을 이용해 다음 경로에서 다운로드 `https://github.com/twodude/solc-breakdown.git`

### Install Dependencies
```bash
$ ./scripts/install_deps.sh
```

### Command-Line Build
```bash
$ ./scripts/build.sh
```
or
```bash
$ mkdir build
$ cd build
$ cmake .. && make
```

## 수정된 Solidity Compiler 실행

`solc` 키워드로 실행 가능.

* For example:
```bash
$ solc --combined-json abi,bin RelayNaivePoW.sol > RelayNaivePoW.json
```
컴파일 후 abi(인터페이스)와 bin(바이너리)를 json 파일에 저장함.


# Relayer

* 서로 다른 체인의 정보를 교환할 수 있도록 하는 주체.

* A 체인의 블록 정보를 B 체인에 등록, 검증함.

* A 체인의 트랜잭션을 B 체인 상에서 유효성을 검증할 수 있다.

## Relay 다운로드

git을 이용해 다음 경로에서 다운로드 `https://github.com/twodude/eth-proof-sol.git`
