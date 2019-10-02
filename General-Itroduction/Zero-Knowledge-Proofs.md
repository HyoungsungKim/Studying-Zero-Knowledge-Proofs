# Zero Knowledge Proofs: An illustrated primer

https://blog.cryptographyengineering.com/2014/11/27/zero-knowledge-proofs-illustrated-primer/

In this series of posts try to give a (mostly) *non*–*mathematical* description of what ZK proofs are, and what makes them so special.

## Origin of Zero Knowledge

Researchers were working on problems related to [interactive proof systems](http://en.wikipedia.org/wiki/Interactive_proof_system), theoretical systems where a first party (called a ‘Prover’) exchanges messages with a second party (‘Verifier’) to convince the Verifier that some mathematical statement is true.

> Interactive proof systems (wikipedia)
> In computational complexity theory, an **interactive proof system** is an abstract machine that models computation as the ***exchange of messages between two parties.***
>
> The parties, the verifier and the prover, interact by exchanging messages in order to ascertain(알아내다) whether a given **string** belongs to a **language** or not.
>
> - The prover is all-powerful and possesses unlimited computational resources, but cannot be trusted,
> - While the verifier has bounded computation power. Messages are sent between the verifier and prover until the verifier has an answer to the problem and has "convinced" itself that it is correct.
>
> 검증자와 증명자 사이에서 메세지를 교환하면서 메세지가 정해진 정답에 포함되어 있는지 여부를 어떻게 증명 할 것인가
>
> - 증명자는 무한한 컴퓨팅 파워를 가지고 있지만 신뢰되지 않음
> - 검증자는 제한된 컴퓨팅 파워를 가지고 있음
>   - 제한된 컴퓨팅 파워를 가지고 있기 때문에 최소한의 노력으로 증명자의 신뢰 여부를 결정해야 함
>   - 증명자는 무한한 컴퓨팅 파워를 가지고 있기 때문에 악의적인 공격이 가능 함
>   - 예를 들어 사용자가 로그인 하기 위해 비밀 번호 입력
>     - 검증자 : 사용자
>     - 증명자 : 서버

What Goldwasser, Micali and Rackoff did was to turn this problem on its head. Instead of worrying only about the Prover, they asked: ***what happens if you don’t trust the Verifier? ***

> 검증자는 어떻게 믿을 것인가?
>
> 증명한 사람이 사실만을 말해도 검증자는 신뢰 할 수 있는가?

The specific concern they raised was *information leakage.* Concretely, they asked, how much extra information is the Verifier going to learn during the course of this proof, beyond the mere fact that the statement is true?

> 검증자는 얼마나 많은 정보를 믿고 신뢰 할 것인가?

What Goldwasser, Micali and Rackoff proposed was a new hope for conducting such proofs. If fully realized, zero knowledge proofs would allow us to prove statements like the one above, while provably revealing *no information* beyond the single bit of information corresponding to ‘this statement is true’.

> 현실 세계에서 사용하는 비밀번호는 서버로 평문(cleartext)을 전송하면 평문을 사용하여 해쉬 연산 후 결과 값을 비교 함
>
> - 서버가 사용자의 비밀번호 평문 학습이 가능 함
> - 사용자는 서버를 일방적으로 믿을 수 밖에 없음
>
> 최소한의 정보만 제공하여 검증-증명 할 필요가 있음

## A crazy technical solution (with hats!)

1. 내가 방에 들어가서 바닥에 종이를 깔아 놓음
2. 다른 사람이 들어와서 종이를 색칠하고 무작위로 섞어 놓은 뒤 모자로 종이들을 가려 놓음
3. 내가 방에 들어와서 몇몇 모자를 들고 종이의 색을 확인 함
   1. 만약 종이가 색칠 되어 있지 않다면 거짓 말을 치고있는 것임
   2. 만약 종이가 색칠 되어 있다면 거짓말을 치지 않을 ***확률***이 존재
      - 우연히 내가 확인한 부분만 색칠 되어 있을 수도 있기 때문에

- 테스트를 계속 해서 반복한다면 거짓말을 치고 있는지 아닌지 확인 할 수 있음
- 우연히 맞을수는 있지만 반복 횟수가 증가 할 수록 거짓말이 걸릴 확률이 증가하기 때문에

## What makes it ‘zero knowledge’?

The first rule of modern cryptography is *never to trust people* who claim such things without proof.

Goldwasser, Micali and Rackoff proposed ***three following properties  that every zero-knowledge protocol must satisfy. Stated informally, they  are:***

1. ***Completeness.*** If Google is telling the truth, then they will eventually convince me (at least with high probability).
2. ***Soundness.*** Google can *only* convince me *if* they’re actually telling the truth.
3. ***Zero-knowledgeness.*** (Yes it’s really called this.) I don’t learn anything else about Google’s solution.
   - 나는 참, 거짓 만 알 수 있지 학습 할 수는 없음.

> 1. 만약 계속 해서 사실을 말 한다면 나는 결국 믿을 것임
> 2. 나는 사실만을 말했을 때 믿음
> 3. 나는 솔루션을 학습하지 않음

The hard part here is the ***‘zero knowledgeness’*** property. To do this, we need to conduct a very strange thought experiment.

## A thought experiment(with time machines)

Their idea is to sneak into the GoogleX workshop and borrow Google’s prototype time machine.

- Initially the plan is to travel backwards a few years and use the extra working time to take another crack at solving the problem.
- Unfortunately it turns out that, like most Google prototypes, the time machine has some limitations. Most critically: it’s only capable of going backwards in time *four and a half minutes.*

***It turns out that even this very limited technology can still be used to trick me.***

Inevitably, though, I’m going to pull off a pair of hats and discover two vertices of the *same* color. In the normal protocol, Google would now be totally busted. And this is where the time machine comes in. 

In effect, the time machine allows Google to ‘repair’ any accidents that happen during their bogus protocol execution, which makes the experience look totally legitimate to me. 

***In fact, from my perspective, being unaware that the time machine is in the picture, the resulting interaction is exactly the same as the real thing.***

> 실제로 나에게는 정상적으로 동작하는 것 처럼 보임

## What the hell is the point of this?

What we’ve just shown is that if time doesn’t run only forward — specifically, if Google can ‘rewind’ my view of time — then they can fake a valid protocol run ***even if they have no information at all about the actual graph coloring.***

Believe it or not, this proves something very important.

Specifically, assume that I (the Verifier) have some strategy that ‘extracts’ useful information about Google’s coloring after observing an execution of the honest protocol. Then my strategy should work equally well in the case where I’m being fooled with a time machine. The protocol runs are, from my perspective, statistically identical. I physically cannot tell the difference.

> 만약 내가 구글의 정직한 coloring 정보를 추출(학습) 하고 전략을 세웠다면, 이 전략은 Time machine을 사용 했을 때에도 동일하게 동작해야 함. 물리적으로 차이가 없기때문에 나에게는 완전히 동일하게 보임

Thus if the amount of information I can extract is identical in the ‘real experiment’ and the ‘time machine experiment’, yet the amount of information Google puts into the ‘time machine’ experiment is exactly  zero — then this implies that even in the real world the protocol must not leak any useful information.

> 따라서 만약 실제 실험이나 타임 머신 실험에서 얻는 정보의 양이 같다면, 구글이 타임 머신 실험에서 제공하는 정보가 0이면 실제 실험에서 내가 얻는 정보의 양도 0임을 의미 함.(어떠한 정보도 유출 하지 않음)

Thus it remains only to show that computer scientists have time machines. We do! (It’s a well-kept secret.)

## Getting rid of the hats (and time machines)

To tie things together, we first need to bring our protocol into the digital world.

-  This requires that we construct the digital equivalent of a ‘hat’: something that both hides a digital value
- While simultaneously ‘binding’ (or ‘committing’) the maker to it, so she can’t change her mind after the fact.

Fortunately we have a perfect tool for this application. It’s called a digital ***commitment scheme***. A commitment scheme allows one party to ‘commit’ to a given message while keeping it secret, and then later ‘open’ the resulting commitment to reveal what’s inside. They can be built out of various ingredients, including (strong) cryptographic ***hash functions.***

>A **commitment scheme** is a cryptographic primitive that allows one to commit to a chosen value (or chosen statement) while keeping it hidden to others, with the ability to reveal the committed value late
>
>cryptographic primitive : one-way hash 함수 같은 방법

What we can now prove is the following theorem:

- If you could ever come up with a computer program (for the Verifier) that extracts useful information after participating in a run of the protocol, 
- then it would be possible to use a ‘time machine’ on that program in order to make it extract the same amount of useful information from a ‘fake’ run of the protocol where the Prover doesn’t put in any information to begin with.

> 증명자가 실제 실행에서 계속해서 정보를 추출 할 수 있다면 그에 맞는 컴퓨터 프로그램을 제시 할 수 있음.
>
> 따라서 이 정보를 이용한다면 증명자가 어떠한 정보도 입력하지 않은 가짜 실행(fake run)으로부터 같은 양의 정보를 추출 해 내는것이 가능 함.

And since we’re now talking about *computer programs*, it should  be obvious that rewinding time isn’t such an extraordinary feat at all. In fact, we rewind computer programs all the time. For example, consider using virtual machine software with a snapshot capability.

Ultimately what we get is the following theorem.

- If there exists any Verifier computer program that successfully extracts information by interactively(쌍방향으로) running this protocol with some Prover, then we can simply use the rewinding trick on that program to commit to a random solution,
  - then ‘trick’ the Verifier by rewinding its execution whenever we can’t answer its challenge correctly.

> 검증 방법에 대한 정보를 추출하여 프로그램을 개발 했다면 rewind하여 랜덤 솔루션으로 속임
> (틀릴 때 까지 계속 실험 -> 1/2 라도 반복 횟수 증가 할 수록 우연히 맞출 확률 떨어짐)

- The same logic holds as we gave above: if such a Verifier succeeds in extracting information after running the real protocol,
  - then it should be able to extract the *same amount of information* from the simulated, rewinding-based protocol.
  - But since there’s no information going into the simulated protocol, there’s no information to extract. Thus the information the Verifier can extract must always be zero.

> 검증자도 rewinding-based protocol을 할 수있지만 가상 실험에 사용 할 수있는 정보가 없기 때문에 얻을 수 있는 정보도 0이다...? -> zero-knowledge

## Ok, so what does this all mean?

We showed that any Verifier program that succeeds in extracting information must also be able to extract information from a protocol run where rewinding is used and *no information is available in the first place.* Which leads to an obvious contradiction, and tells us that the protocol can’t leak information in either situation.

> 비밀번호로 로그인 하려 할 때 검증자(서버)는 얼마든지 테스트 해 볼 수 있지만 증명자(사용자)는 검증 과정을 모르기 때문에 테스트를 해도 정보를 얻을 수 없음.
>
> - 증명자는 참, 거짓 여부 만 알 수 있음
> - 검증자는 rewind를 통해 계속해서 실험 반복하여 정말로 증명자가 정답을 알고 있는지 우연히 맞춘건지 검증 함.

I didn’t simply *edit* in the same way Google might have done using the time machine. This means that protocol transcripts(기록, 증명) themselves contain no information. The protocol is only meaningful if I myself participated, and I can be sure that it happened in real time.