# Introducing Sonic: A Practical zk-SNARK with a Nearly Trustless Setup

https://www.benthamsgaze.org/2019/02/07/introducing-sonic-a-practical-zk-snark-with-a-nearly-trustless-setup/

Unlike other SNARKs, Sonic does not require a trusted setup for each circuit, but only a single setup for all circuits

- Further, the setup for Sonic never has to end, so it can be continuously secured by accumulating more contributions.
- This property makes it ideal for any system where there is not a trusted party, and there is a  need to validate data without leaking confidential information.

## More about zk-SNARKs

Zcash uses zk-SNARKs to send private cryptocurrency transactions, and if their setup was compromised then a small number of users could generate an unlimited supply of currency without detection.

### Characteristics of zk-SNARKs

Good

- Can be used to build many cryptographic protocols
- Very small proof sizes
- Very fast verifier time

Not bad

- Average prover time

Bad

- Requires a ***trusted setup***
- Security assumes non-standard cryptographic assumptions

## Updatability

Updatability means that any user, at any time, can update the parameters, including after the system goes live. This property means that a distrustful user could update the parameters themselves and have personal confidence in the parameters from that  point forward. The update proofs are short and quick to verify.

## Universality

Universality means that the same parameters can be used for any application using this zk-SNARK.

## Why Use Sonic?

Sonic is universal, updatable, and has a small set of global parameters  (in the order of megabytes). Proof sizes are small (256 bytes) and  verifier time is competitive with the fastest zk-SNARKs in the  literature. 

The advantage of Sonic over Groth et al.’s design is that ***Sonic uses a smaller set of global parameters.***

- As a result, storing, updating, and verifying Sonic’s parameters could be achieved on a personal laptop.
- On the other hand, Groth et al.’s public parameters are expensive to store, update, and verify, to the extent that it would require a distributed set of computers to carry out these tasks.
- Also, to use Groth et al.’s zk-SNARK with a specific application, some party has to run an expensive (but untrusted) derivation process on the global parameters to obtain application-specific parameters.
  - ***Sonic does not require any derivation process.***
- Sonic would be used in different applications to zk-STARKs.
  - STARKs are feasibly postquantum secure and have a fully trustless setup, but  have high concrete overhead.
  - Even for modestly-sized applications, current STARK proof sizes are at least 40kB.
    - This property makes them more ***suitable for applications where the STARK is run as a one-off cost*** and is less suitable for applications where multiple proofs for the same program run, stored and verified.

## Techniques

Sonic describes its problem statements using the same techniques as Bulletproofs.

- It is built over pairing based groups and makes heavy use of a constant-sized polynomial commitment scheme by [Kate et al.](https://www.iacr.org/archive/asiacrypt2010/6477178/6477178.pdf).
- It assumes two rounds of interaction between the prover and the verifier and is then made non-interactive in the random oracle model.

***Sonic has two modes of operation, the “helper” mode, and the “unhelped” mode.***

- The helper mode is simpler and more efficient but assumes the presence of untrusted third parties that aggregate a number of proofs together.
- The unhelped mode assumes neither batching or aggregation, however, is more expensive (estimated proof sizes over 256-bit groups is 1kb).
  - For each proof, the Sonic verifier requires knowledge of the evaluation of a (sparse) polynomial *s(z,y)* for known polynomial *s(X,Y)*.
  - Computing this polynomial for themselves requires an undesirable level of work for the verifier, but a reasonable level of work for the prover or helper.

## Helper Mode

The helper commits to a batch of proofs. They then calculate the verifiers’ polynomials and provide a short, easy to verify proof of the correct calculation. 

- In the blockchain setting, they could be run by the miners.
- Given *m* proofs, the helper runs in time proportional to *m* provers.
- The helpers’ proof sizes are small.
  - Verifying the help requires a one-off calculation of the polynomial *s(X,Y)* at known points *z,y* and then a small computation per proof (this computation is independent of the size of the application).

## Unhelped Mode

In the unhelped mode, ***the prover*** calculates the polynomial for the verifier and provides a short, easy to verify proof of the correct calculation. 

- This mode is ideal for settings either where proofs cannot be batched or where there are no helpers available.
- While the proof size and verifier computation remain independent of the size of the computation, they are concretely several times larger than the proof size and the verifier computation of the helped proofs.