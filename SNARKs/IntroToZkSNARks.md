# Introduction to zk-SNARKs with Examples

https://media.consensys.net/introduction-to-zksnarks-with-examples-3283b554fc3b

Zk-SNARKs : Non-interactive zero knowledge proofs

## Definition of a zk-SNARK

A zk-SNARK consists of three algorithms `G, P, V` defined as follows:

- The *key generator* `G` takes a secret parameter `lambda` and a program `C`, and generates two publicly available keys, a *proving key* `pk`, and a *verification key* `vk`.
  - G(lambda, C, pk, vk)
  - These keys are public parameters that only need to be generated once for a given program `C`.
- The *prover* `P` takes as input the proving key `pk`, a public input `x` and a private witness `w`.
  - The algorithm generates a *proof* `prf = P(pk, x, w)` that the prover knows a witness `w` and that the witness satisfies the program.
- The *verifier* `V` computes `V(vk, x, prf)` which returns `true` if the proof is correct, and `false` otherwise.
  - Thus this function returns true if the prover knows a witness `w` satisfying `C(x,w) == true`.