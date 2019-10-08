# On Interactive Proofs and Zero-Knowledge: A Primer

https://medium.com/magicofc/interactive-proofs-and-zero-knowledge-b32f6c8d66c3

- NIZKs : Non-interactive zero-knowledge proofs 
- SNARKs : succinct non-interactive arguments of knowledge
- STARKs : succinct transparent arguments of knowledge

## 1. Introduction

Zero-knowledge is a wonderful property of proof systems to guarantee privacy of the computation. It enables to verify the computation over the dataset without revealing any partial information about the data.

### 1.1 What is a proof system

In computational complexity theory, an interactive proof system is an abstract machine that models computation as the exchange of messages between two parties.

- The parties, known as prover **P**
  - The prover is all-powerful and possesses unlimited computational resources, but cannot be trusted
- The Verifier **V**, interact by exchanging messages in order to ascertain(알아내다) whether a statement is correct.
  - The verifier has bounded computational power.

Messages are sent between the verifier and prover ***until the verifier can validate or falsify a statement.***

All interactive proof systems have two key properties:

- **Completeness:** if the statement is true, the honest verifier (that is, one following the protocol properly) can be convinced of this fact by a honest prover.
- **Soundness:** if the statement is false, no cheating/malicious prover can convince the honest verifier that it is true, except with some negligible probability.

> Completeness : 만약 prover가 사실을 말한다면 verifier도 prover를 인정함
>
> Soundness : 만약 prover가 거짓을 말한다면 verifier은 사실로 인정하지 않음(매우 적은 확률을 제외하고는)

### 1.2 Privacy through Zero-Knowledge

If the prover has some secret information — in zk-jargon referred to as **witness**(Zero-knowledge용어로 witness라고 불리는) — and wants the verifier to learn no partial information about it.

Hence, the third magical property of proof systems can be summarized as

- **Zero-knowledge:** If the statement is true, no cheating verifier learns anything other than this fact.
  - verifier(검증자)는 학습하지 않음

> 동굴 예제
>
> 1. Alice가 먼저 동굴에 들어감
> 2. Bob이 나와! 라고 외침
> 3. Alice는 사전에 합의 된 출구로 나옴
> 4. 만약 Alice가 두개의 입구 중 옳바른 입구로 들어갔으면 그대로 나오고 반대편으로 갔을 경우 동굴 안에 있는 문을 열고 옳바른 입구로 나감
> 5. 만약 Alice가 합의 되지 않은 출구로 나오면 Alice는 거짓을 말하는 중임
>
> 군대 암구호도 zero-knowledge proof? - 내 생각... 여기서는 witness를 정의하기 좀 애매한듯
>
> 1. 저녁에 초소를 지키고 있는 군인이 있음
> 2. 누군가 초소에 접근
> 3. 초소를 지키고 있는 군인은 질문을 날림
>    - 암구호는?!?!
> 4. 만약 정상적으로 대답하면 정상 통과
>    - 초소에 접근한 사람은 자신의 신원을 들어내지 않으면서 자신의 신원을 증명함
> 5. 초소를 지키고 있던 군인은 어떠한 정보도 학습하지 않음
>    - 예를 들어 목소리 톤, 발음, 억양 등등을 학습하지 않고 참, 거짓 여부만 판단
>    - prover가 증명에 사용하는 witness는 verifier는 알 수 없음(알면 안됨)
>      - 만약 verifier가 알 경우 도용 가능
>      - zero-knowledge가 아니게 됨

## 2. Formal Definitions: The Devil is in the Detail

In order to formalize the properties we will need a model of computation to express what an interaction between two computing devices is.

### 2.1 Interactive Turing Machine

An interactive Turing machine (ITM) is a (deterministic) multi-tape Turing machine.

- A tapes are a read-only input tape
- A read-only random tape
- A read-and-write work tape
- A write-only output tape
- A pair of communication tapes
  - One communication tape is read-only, and the other is write-only.
- A read-and-write switch tape consisting of a single cell.

You may now want to call the interactive turing machine **M₁** the prover **P** and **M₂** the verifier **V**, respectively*.* Like humans, machines need to understand a language **L** to properly interact. 

All what we need to know is that the statement the prover **P** wants to prove must — very loosely speaking — be encoded in a particular language. With **(x,w) ∈ L** and **(x,w) ∉ L** we express a correct/incorrect proof statement or in other words the pair **(x,w)** is a member of the language **L**.

Typically the value **x** is public and known to both the prover and verifier and the parameter **w** (called the witness) is private and known to the prover only.