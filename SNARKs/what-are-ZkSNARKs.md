# What are zk-SNARKs

https://z.cash/technology/zksnarks/

zk-SNARks : zero-knowledge Succinct Non-interactive Argument of Knowledge'

- Zero-knowledge : Proofs allow one party (the prover) to prove to another (the verifier) that a statement is true, without revealing any information beyond the validity of the statement itself.
- Proof of Knowledge : The prover can convince the verifier not only that the number exists, but that they in fact know such a number – again, without revealing any information about the number.
- Succinct : Zero-knowledge proofs can be verified within a few milliseconds, with a proof length of only a few hundred bytes even for statements about programs that are very large.
  - In the first zero-knowledge protocols, the prover and verifier had to communicate back and forth for multiple rounds, but in ***“non-interactive” constructions, the proof consists of a single message sent from prover to verifier.***

## How zk-SNARKs are constructed in Zcash

**Computation → Arithmetic Circuit → R1CS → QAP → zk-SNARK**

In this `R1CS`(Rank 1 Constraint System) representation, the verifier has to check many constraints — one for almost every wire of the circuit. (For technical reasons, it turns out we only have a constraint for wires coming out of multiplication gates.)

- In a 2012 paper on the topic, Gennaro, Gentry, Parno and Raykova presented a nice way to “bundle all these constraints into one”.
  - This method uses a representation of the circuit called a Quadratic Arithmetic Program (QAP).
  - The single constraint that needs to be checked is now between polynomials rather than between numbers.
- The polynomials can be quite large, but this is alright because when an identity does not hold between polynomials, it will fail to hold at most points.
  - Therefore, you only have to check that the two polynomials match at one randomly chosen point in order to correctly verify the proof with high probability.