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

# Private -> Testnet
Relayer를 사용해 Private 체인의 트랜잭션을 이더리움 테스트넷에서 검증하는 example:

## Compile Smart Contract

미리 작성한 script로 쉽게 compile할 수 있음.

```bash
$ sh naivePoW/deps.sh
```

## Relayer 기본 정보 작성

* Relay를 위한 기본 정보를 `config.json` 파일에 작성

* Relayer의 address
* 트랜잭션 발생을 위한 Relayer's address의 비밀번호(geth상 비밀번호)
* 처음 Relay로 등록한 제네시스 블록의 정보
    * 블록 해시
    * 엉클 블록 해시
    * rlp 인코딩된 bytes 값
* 스마트 컨트랙트의 정보
    * 컨트랙트가 배포될 때 포함된 트랜잭션 ID
    * 컨트랙트의 address

다음 파일을 naivePoW/config.json 경로에 작성
    
```json
{
  "owner": {
    "address": "사용할 ethereum 계좌 주소",
    "password": "geth에서 설정한 비밀번호"
  }
}
```

### `genesis.js` 실행
유효한 blockNumber를 같이 넘겨줘야 함.

```bash
$ node naivePoW/genesis.js [blockNumber]
```

정상 구동 시 `config.json` 파일에 genesisBlock 필드가 추가됨.

* `*.js``const provider =` 부분에 해당하는 Private chain과 Testnet chain의 주소를 올바르게 작성해야 함.
    * 현재는 연구실 컴퓨터에서 구동 중인 geth 클라이언트에 연결됨.

### `contractDeploy.js` 실행
```bash
$ node naivePoW/contractDeploy.js
```

정상 구동 시 `config.json` 파일에 contract 필드가 추가됨.

## Relay 구동

```bash
$ node naivePoW/relay.js 
```

1. 스마트 컨트랙트를 통해 등록된 블록 헤더 중 가장 높은 blockNumber를 가지는 블록을 찾음.
2. 그 이후 블록에서부터 블록 정보를 받아와 등록함.
3. 가장 최신 블록까지 등록을 완료하면, 추가로 채굴이 될 때 까지 대기했다가 다시 등록을 시작함.

## 트랜잭션 검증하기

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
