# Explaining SNARKs

## Part 1 : Homomorphic Hidings

- Similar with discrete log

An HH E(x) of a number x is a function satisfying the following properties:

- For most x’s, given E(x) it is hard to find x
- Different inputs lead to different outputs – so if x≠y, then E(x)≠E(y)
- If someone knows E(x) and E(y), they can generate the HH of ***arithmetic expressions in*** x *and* y. For example, they can compute E(x+y) from E(x) and E(y)

Here’s a toy example of why HH is useful for Zero-Knowledge proofs:

1. Alice sends E(x) and E(y) to Bob

2. Bob computes E(x+y) from these values (which he is able to do since E an HH

3. Bob also computes E(7), and now checks whether E(x+y)=E(7). 
   - He accepts Alice’s proof only if equality holds.

## Part 2 : Blind Evaluation of Polynomials

$$
E(ax+by)=g^{ax+by}=g^{ax}⋅g^{by}=(g^x)^a⋅(g^y)^b=E(x)^a⋅E(y)^b
$$

- discrete logarithm is possible to do linear combination

## Part 3 : The Knowledge of Coefficients Test and Assumption

There was something missing in that protocol – the fact that Alice is *able* to compute `E(P(s))` ***does not guarantee she will indeed send `E(P(s))` to Bob, rather than some completely unrelated value.***

- Thus, we need a way to “force” Alice to follow the protocol correctly.
  - *Knowledge of Coefficient (KC) Test*

## Part 4 : How to make Blind Evaluation of Polynomials Verifiable

### An Extended KCA(Knowledge of Coefficient Assumption)

## Part 5 : From Computations to polynomials

In 2013, Gennaro, Gentry, Parno and Raykova defined an extremely useful translation of computations into polynomials called a *Quadratic Arithmetic Program* (QAP).

- QAPs have become the basis for modern zk-SNARK constructions, in particular those used by Zcash.