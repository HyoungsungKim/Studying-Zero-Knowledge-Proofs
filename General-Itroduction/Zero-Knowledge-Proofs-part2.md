# Zero Knowledge Proofs: An illustrated primer, Part 2

Based on : https://blog.cryptographyengineering.com/2017/01/21/zero-knowledge-proofs-an-illustrated-primer-part-2/

Three critical properties

- **Completeness**: If the Prover is honest, then she will eventually convince the Verifier.
- **Soundness:** The Prover can only convince the Verifier if the statement is true.
- **Zero-knowledge(ness):** *The Verifier learns no information beyond the fact that the statement is true.*

The real challenge turns out to be finding a way to formally define the last property. 

