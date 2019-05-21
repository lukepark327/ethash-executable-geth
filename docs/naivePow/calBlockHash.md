# Research

## MixDigest

Mixhash is actually calculated from nonce as intermediate value when validating PoW with Hashimoto algorithm. But this calculation is still pretty heavy and a node might be DDoSed by blocks with incorrect nonces. mixhash is included into block to perform lightweight PoW 'pre-validation' to avoid such attack, as generating a correct mixhash still requires at least some work for attacker.



# References

[1] https://ethereum.stackexchange.com/questions/5833/why-do-we-need-both-nonce-and-mixhash-values-in-a-block   
