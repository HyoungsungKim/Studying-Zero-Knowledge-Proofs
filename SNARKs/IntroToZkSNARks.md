# Introduction to zk-SNARKs with Examples

https://media.consensys.net/introduction-to-zksnarks-with-examples-3283b554fc3b

Zk-SNARKs : Non-interactive zero knowledge proofs

## Zero-knowledge proofs

The goal of zero-knowledge proofs is for a *verifier* to be able to convince herself that a *prover* possesses knowledge of a secret parameter, called a *witness*, satisfying some relation, without revealing the witness to the verifier or anyone else.

## Definition of a zk-SNARK

Verifier generates pk, vk -> Prover makes prf using pk and x -> Verifier validates prf using vk and x

- Verifier and prover know x but don't share it

A zk-SNARK consists of three algorithms `G, P, V` defined as follows:

- The *key generator* `G` takes a secret parameter `lambda` and a program `C`, and generates two publicly available keys, a *proving key* `pk`, and a *verification key* `vk`.
  -  pk, vk := G(lambda, C)
  - These keys are public parameters that only need to be generated once for a given program `C`.
- The *prover* `P` takes as input the proving key `pk`, a public input `x` and a private witness `w`.
  - The algorithm generates a *proof* `prf = P(pk, x, w)` that the prover knows a witness `w` and that the witness satisfies the program.
- The *verifier* `V` computes `V(vk, x, prf)` which returns `true` if the proof is correct, and `false` otherwise.
  - Thus this function returns true if the prover knows a witness `w` satisfying `C(x,w) == true`.
- This parameter(`lambda`) sometimes makes it tricky to use zk-SNARKs in real-world applications.
  - The reason for this is that ***anyone who knows this parameter can generate fake proofs*** such that `V(vk, x, fake_prf)` evaluates to `true` without knowledge of the secret `w`.

## zk-SNARKs in Ethereum

Developers have already started integrating zk-SNARKs into Ethereum.

- Concretely, the building blocks of the verification algorithm is added to Ethereum in the form of precompiled contracts. The usage is the following:
  - The generator is run `off-chain` to produce the proving key and verification key.
  - Any prover can then use the proving key to create a proof, also `off-chain`.
  - The general verification algorithm can then be run inside a smart contract, using the proof, the verification key and the public input as input parameters.
  - ***The outcome of the verification algorithm can then be used to trigger other on-chain activity***.

## Comparison of the most popular zkp systems

|                                       |                     SNARKs |                             STARKs |         Bulletproofs |
| ------------------------------------: | -------------------------: | ---------------------------------: | -------------------: |
|        Algorithmic complexity: prover |              O(N * log(N)) |                 O(N * poly-log(N)) |        O(N * log(N)) |
|      Algorithmic complexity: verifier |                      ~O(1) |                     O(poly-log(N)) |                 O(N) |
| Communication complexity (proof size) |                      ~O(1) |                     O(poly-log(N)) |            O(log(N)) |
|              - size estimate for 1 TX |  Tx: 200 bytes, Key: 50 MB |                              45 kB |               1.5 kb |
|         - size estimate for 10.000 TX | Tx: 200 bytes, Key: 500 GB |                             135 kb |               2.5 kb |
|    Ethereum/EVM verification gas cost |            ~600k (Groth16) |         ~2.5M (estimate, no impl.) |                  N/A |
|               Trusted setup required? |             YES :unamused: |                         NO :smile: |           NO :smile: |
|                   Post-quantum secure |              NO :unamused: |                        YES :smile: |        NO :unamused: |
|                    Crypto assumptions |          Strong :unamused: | Collision resistant hashes :smile: | Discrete log :smirk: |