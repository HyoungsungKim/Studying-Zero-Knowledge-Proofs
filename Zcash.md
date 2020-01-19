# Zcash

주소 소유자는 주소를 다른 사람에게 노출 할지를 결정 할 수 있음

- z-주소 : 보호 주소(Shilded address)
- t-주소 : 공개(투명) 주소 (Transparent address)
  - t-주소 상의 트랜잭션은 비트코인과 유사하게 동작
- 두 주소는 서로 상호 운용이 가능 함

## 작동방식

Sender - Receiver

- 비공개 - 비공개
- 비공개 - 공개
- 공개 - 비공개
- 공개 - 공개

## Zcash 트랜젝션 상세 구조

https://electriccoin.co/ko_KR/blog/anatomy-of-zcash/

두개의 pool(Shilded pool, Transparent value pool)이 존재하고, pool들은 서로 상호 운용이 가능함. 즉 Shilded pool에서 TV pool로, TV pool에서 Shilded pool로 전송, 수신이 가능함.

- 비공개 거래 전송 -> Shilded pool -> 비공개 거래 수신
  - 보호 주소를 사용하기 위해서는 영지식 증명을 사용해야 함
    - Zero-knowledge proof
    - Shilded input ZEC
    - Shilded output ZEC
  - 보호 주소끼리 거래가 발생하면 Tx에는 입력 금액은 0 ZEC로 표시되고, 출력금액은 수수료만 표시 됨
- 공개 거래 전송 ->  Transparent value pool -> 공개 거래 수신

## Zcash의 주소변경

매 거래때마다 새로운 주소를 생성해서 거래를 하더라도, UTXO 같은 방식에서는 잔액을 돌려 받을 때 똑같은 주소를 사용하기 때문에 의미 있는 수준의 보안 레벨을 제공하지 않음

- Zcash에서는 Shilded address에서 ZEC를 보낼 때, 데이터의 은밀함(Private)이 유지 됨. 즉, Shilded address는 모두 동일하게 보이기 때문에 같은 주소를 사용하더라도 보안상으로 취약한점이 없음
  - 어떤 방식으로 동일하게 보이도록 만드는지 문헌조사 필요

## Shilded address와, Transparent address 사이의 거래

Zcash에서의 화폐인 ZEC는 다른 암호화폐와 다른 특징을 가지고 있음, 특징은 주소에 의해 결정되는데

1. ZEC 잔액을 가지고 있는 주소의 타입에 의해서
2. ZEC의 전 소유자의 주소의 타입에 따라서

- 만약 ZEC가 TP address에 있다면 잔앤을 누구나 확인 할 수 있음
  - Shilded address에서 ZEC를 보내더라도 받는 사람이 TP address면 보낸 금액을 확인 할 수 있음
- TP address에서 Shilded address로 금액을 보낸다면, 미래에 Shilded address로부터 이 금액이 다른 주소로 전달 되더라도 최초 Tx를 발생 시킨 TP address가 추적이 안된다는 장점이 있음
- Shilded address에서 TP address로 보내면 TP address가 얼마를 받았는지 확인 되지만, 보낸 사람의 주소와 얼마를 보냈는지는 확인이 안됨.

## 보호 주소간의 트랜젝션 동작 방식

https://electriccoin.co/blog/zcash-private-transactions/

비트코인 기반의 UTXO에 Zcash 특징을 붙여서 설명 함. 설명을 위해 비트코인의 UTXO를 Unspent note라고 정의하고, 이 note에는 사용자의 address/public key, 금액의 총 량이 적혀 있음.

가정

1. 모든 UTXO는 1개의 BTC만 가지고 있음
2. 1개의 노트에는 1개의 주소만 존재 함. 즉 각각의 노트는 소유자를 의미 함

$$
Note_4 = (PK_4), Note_2 = (PK_2), Note_3 = (PK_3)
$$

- PK : Public key

여기에 Privacy를 위해 임의의 serial number `r`을 추가 함
$$
Note_4 = (PK_4,r_1), Note_2 = (PK_2, r_2), Note_3 = (PK_3, r_3)
$$
Privacy를 위한 첫번째 단계는 오직 암호화 된 정보만 저장하는 것임. 따라서 각 Note를 Hash함수에 넣은 뒤 저장 함
$$
H_1=HASH(Note_1), H_2 = HASH(Note_2), H_3= HASH(Note_3)
$$
Privacy를 위한 두번째 단계는 UTXO가 사용되더라도 계속 데이터베이스에 저장 함. 즉, 사용되지 않는 UTXO와 사용 된 UTXO 모두 저장 됨

- 데이터 베이스 = hashed note + nullifier set
- 사용된 것과 사용되지 않은것을 Privacy문제 없이 구분 할 방법이 필요 함
- 이를 위해 nullifier가 사용 됨
  - Nullifier란 노트의 serial number의 해시 값임
- 예를 들어 Note2가 사용되면 데이터 베이스의 hashed note에 Note2의 hash값이 계속 남아있고,
- 데이터 베이스의 Nullifier set에 r2(Note2의 serial number)의 해시 값이 저장 됨

### 트랜젝션 생성 및 검증 방법

Note1을 소유한 A가 PK4를 가진 B에게 1BTC를 전송하려 함

1. A가 무작위로 새로운 serial number를 생성하고 새로운 노트를 정의 함(Note4 = (PK4, r4))
2. A가 B에게 Note4를 private하게 보냄
3. A가 Note1의 nullifier(nf1 = HASH(r1))를 모든 노드에게 전송함
   - r4가 아니라 Note1의 r1을 저장 함
4. A가 H4(H4 = HASH(note4))를 모든 노드에게 전송 함

- 모든 노드가 H4와 nf1을 받았을 때 nf1이 nullifier set에 존재함을 확인해서 이전에 사용된 적이 없는지 확인 함.
  - 만약 존재하지 않는다면 추가 함
  - H4 역시 hash note에 추가
- 트랜젝션이 검증 됨
- 하지만 이 과정에서 Note1이 사용되었는지 확인하는게 없음
  - 만약 Note1이 이미 사용됐다면 Note4를 만들 수 없음.
  - Note1이 사용되었는지를 확인하는 방법은 A가 Note1을 공개하는 것이지만, 이것은 pricacy하지 않음
    - 데이터 베이스에는 HASH(Note1)이 저장 되어 있음
  - 이 과정에서 영지식 증명이 필요 함

#### 영지식 증명을 이용한 검증

A는 Note1의 소유자이기 때문에 Pk1, Sk1(private key), r1을 모두 알고있음. A는 자신이 Note1의 소유자임을 다른 노드들에게 납득시키기 위해 proof-string($$\pi$$)를 생성 함. Proof-string으로는 어떠한 정보도 노출되지 않음.

