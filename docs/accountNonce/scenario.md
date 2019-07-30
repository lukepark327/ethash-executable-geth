# Geth

## Pre-mining

1. `miner.start()`

2. `miner.stop()`

3. `eth.blockNumber`

  * 12
  
## Send Tx

1. `personal.unlockAccount(eth.coinbase);`

2. `eth.sendTransaction({from:eth.coinbase, to:'0x1111111111111111111111111111111111111111', value: web3.toWei(0.05, "ether"), gas:21000});`

  * INFO [07-25|11:20:15.580] RESET NONCE
  * base=832 nonce=0
    * 832 is 64 * 13 (MaxTransactionLimit * NextBlockNumber)
  * Now next block's nonce is 832
  * Current nonce is 0
    
3. `eth.sendTransaction({from:eth.coinbase, to:'0x1111111111111111111111111111111111111112', value: web3.toWei(0.05, "ether"), gas:21000});`

  * INFO [07-25|11:28:42.098] NONCE++
  * base=832 nonce=832
  * Now next block's nonce is 833
  * Current nonce is 1
  
4. `eth.sendTransaction({from:eth.coinbase, to:'0x1111111111111111111111111111111111111113', value: web3.toWei(0.05, "ether"), gas:21000});`

  * INFO [07-25|11:28:42.098] NONCE++
  * base=832 nonce=833
  * Now next block's nonce is 834
  * Current nonce is 2
  
## Mining

1. `miner.start()`

2. `miner.stop()`

3. `eth.blockNumber`

  * 14
  
4. `eth.getBlock(13)`

  * transactions: ["0x485329e3c114c5cf9a10acfb6c80b02f42b61b647241235875a7ab78d98009d6", "0xbf215ed47372801998d75e3318f8f14d970e3edee1e603427ee88dd8cbea64fd", "0xd5c68561df7291592ee32e145091d01e30324cdfd1c85fcb6095ead6cc1c36fc"]
  
## Transactions

1. `eth.getTransaction("0x485329e3c114c5cf9a10acfb6c80b02f42b61b647241235875a7ab78d98009d6")`

  * nonce: 0
 
2. `eth.getTransaction("0xbf215ed47372801998d75e3318f8f14d970e3edee1e603427ee88dd8cbea64fd")`

  * nonce: 1
  
3. `eth.getTransaction("0xd5c68561df7291592ee32e145091d01e30324cdfd1c85fcb6095ead6cc1c36fc")`

  * nonce: 2
  
## Send Tx

1. `eth.sendTransaction({from:eth.coinbase, to:'0x1111111111111111111111111111111111111111', value: web3.toWei(0.05, "ether"), gas:21000});`

  * INFO [07-25|12:56:21.539] RESET NONCE
  * base=960 nonce=834
    * 960 is 64 * 15 (MaxTransactionLimit * NextBlockNumber)
  * Now next block's nonce is 960
  * Current nonce is 834
    
2. `eth.sendTransaction({from:eth.coinbase, to:'0x1111111111111111111111111111111111111112', value: web3.toWei(0.05, "ether"), gas:21000});`

  * INFO [07-25|12:57:45.881] NONCE++
  * base=960 nonce=960
  * Now next block's nonce is 961
  * Current nonce is 835
  
