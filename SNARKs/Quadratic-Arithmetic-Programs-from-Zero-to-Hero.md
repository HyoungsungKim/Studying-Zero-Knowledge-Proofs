# Quadratic Arithmetic Programs: from Zero to Hero

https://medium.com/@VitalikButerin/quadratic-arithmetic-programs-from-zero-to-hero-f6d558cea649

Computation -> Algebraic Circuit -> R1Cs -> QAP -> Linear PCP -> Linear Interactive Proof -> zkSNARKS

The steps here can be broken up into two halves. 

- First, zk-SNARKs cannot be applied to any computational problem directly;
  - Rather, ***you have to convert the problem into the right “form” for the problem to operate on.***
  - The form is called a ***“quadratic arithmetic program” (QAP)***, and transforming the code of a function into one of these is itself highly nontrivial.
- Along with the process for converting the code of a function into a QAP is another process that can be run alongside
  - So that if you have an input to the code you can create a corresponding solution (sometimes called “witness” to the QAP).
- After this, there is another fairly intricate process for creating the actual “zero knowledge proof” for this witness, and a separate process for verifying a proof that someone else passes along to you.

## Flattening

The first step is a “flattening” procedure, where we convert the original code, which may contain arbitrarily complex statements and expressions, into a sequence of statements

```
// out = x**3 + x + 5
// Flattening
sym_1 = x * x
y = sym_1 * x
sym_2 = y + x
~out = sym_2 + 5
```

## Gates to R1CS

R1CS : Rank-1 constraint system

We convert flattening into something called a rank-1 constraint system (R1CS).

- An R1CS is a sequence of groups of three vectors `(a, b, c)`, ***and the solution to an R1CS is a vector `s`***, where `s` must satisfy the equation `s . a * s . b - s . c = 0`
  - `.` represents the dot product - in simpler terms, if we "zip together" `a` and `s`, multiplying the two values in the same positions, and then take the sum of these products, then do the same to `b` and `s` and then `c` and `s`, ***then the third result equals the product of the first two results.***
- 