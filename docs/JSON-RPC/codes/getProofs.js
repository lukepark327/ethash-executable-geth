const rlp = require('rlp');
const Web3 = require("web3");

const provider = "https://ropsten.infura.io";
const web3 = new Web3(provider);

// values
const cpInterval = 10;

// functions
function getProofs(address, storageKey, startBlockNumber, endBlockNumber) {
    var proofs = [];
    for (var i = startBlockNumber; i <= endBlockNumber; i += cpInterval) {
        var proof = web3.eth.getProof(address, storageKey, i).catch((err) => console.log(err));
        proofs.push(proof);
    }
    return proofs;
}

// main
var address = process.argv[2];          // "0x7224769b9eE714dAA816053732D6Ed0AA35714CB";
var storageKey = [];                    // Empty for EOA
var from = Number(process.argv[3]);     // 6011146;
var to = Number(process.argv[4]);       // 6011172;

// Get Proofs
var proofs = getProofs(address, storageKey, from, to);
Promise.all(proofs).then((res) => {
    var accountProofs = [];    
    res.forEach((proof) => {
        accountProofs.push(proof.accountProof);
    });

    /*
    var preRlp = {
        "address" : address,
        "startBlockNumber" : from,
        "accountProofs" : accountProofs
    };
    */
    var preRlp = [
        address,
        from,
        accountProofs
    ];

    // print
    // console.log(preRlp);
    console.log(rlp.encode(preRlp));
});
