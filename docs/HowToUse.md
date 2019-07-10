* Relayer를 사용해 Private 체인의 트랜잭션을 이더리움 테스트넷에서 검증하는 example:

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
