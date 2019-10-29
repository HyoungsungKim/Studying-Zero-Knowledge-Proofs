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

## R1CS to QAP

The next step is taking this R1CS and converting it into QAP form, which implements the exact same logic except using polynomials instead of dot products.

We go

- From four groups of three vectors of length six
- To six groups of three degree-3 polynomials, where evaluating the polynomials at each *x coordinate* represents one of the constraints.

That is, if we evaluate the polynomials at x=1, then we get our first set of vectors If we evaluate the polynomials at x=2, then we get our second set of vectors, and so on.

If you have a set of points (ie. (x, y) coordinate pairs), then doing a Lagrange interpolation on those points gives you a polynomial that passes through all of those points.

We do this by decomposing the problem:

- For each x coordinate, we create a polynomial that has the desired y coordinate at that x coordinate.
- A y coordinate of 0 at all the other x coordinates we are interested in, and then to get the final result we add all of the polynomials together.

## Checking the QAP

we can now check *all of the constraints at the same time* by doing the dot product check *on the polynomials*.

Because in this case the dot product check is a series of additions and multiplications of polynomials, the result is itself going to be a polynomial.

- If the resulting polynomial, evaluated at every x coordinate that we used above to represent a logic gate, is equal to zero then that means that all of the checks pass;
- If the resulting polynomial evaluated at at least one of the x coordinate representing a logic gate gives a nonzero value, then that means that the values going into and out of that logic gate are inconsistent (ie. the gate is `y = x * sym_1` but the provided values might be `x = 2`,`sym_1 = 2` and `y = 5`).