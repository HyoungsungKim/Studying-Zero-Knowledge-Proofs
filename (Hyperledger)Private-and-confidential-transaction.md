# (Hyperledger)Private and confidential transaction

https://developer.ibm.com/tutorials/cl-blockchain-private-confidential-transactions-hyperledger-fabric-zero-knowledge-proof/

How this open source blockchain framework ensures custom levels of privacy and confidentiality

>“*In transaction systems, such as blockchain networks, privacy is  frequently achieved by mechanisms that enforce the confidentiality of  data and the anonymity of transaction participants.*“

- *Data confidentiality* mechanisms ensure that individuals or organizations are prevented from accessing data that they are not authorized to access, such as classified information of other organizations’ transactions.
  - Data confidentiality는 허가된 사람이나 조직만이 접근을 허락 함
- *Anonymity* requires that participants of transactions are concealed.
  - Anonymity는 거래의 참가자를 감춤(숨김)

## Privacy in public vs. private blockchain systems

The public availability of the content of transactions in such systems can be problematic for numerous business use cases. Why?

- Imagine a  business obtaining computer parts from a vendor. Given the large volume of computer parts purchased, the supplier provides a discount to the business when trading the asset for currency. For the supplier, the  actual discount is sensitive business information, as the supplier may not want to provide the same discount to businesses who purchase lower volumes.
- 공급자가 특정 기업에게 할인을 제공 할 경우 공급자는 이 정보는 민감한 정보이기 때문에 공개 하고 싶지 않을수가 있음
- 이럴 경우 public한 blockchain을 사용하는건 적절하지 않음

Permissioned blockchains have emerged as an alternative to public ones to address enterprise needs for having known and identifiable participants. They’ve achieved the scale, confidentiality, and privacy necessary to enable enterprise applications.

## Privacy mechanisms in Hyperledger Fabric

Starting from its permissioned nature, Hyperledger Fabric offers mechanisms to accommodate multiple flavors of privacy, depending on the use case.

### Channels in Hyperledger Fabric

Hyperledger Fabric’s channel architecture can offer privacy in certain cases.

> “*A channel is like a virtual blockchain network that sits on top of a physical blockchain network with its own access rules.* ***Channels employ their own transaction ordering mechanism and thus provide scalability,***  *ultimately allowing for effective ordering and partition of huge amounts of data.*“

***Channels in Hyperledger Fabric are configured with access policies*** that govern access to the channel’s resources (chaincodes, transactions, and  ledger state), thus preserving the privacy and confidentiality of information exclusively within the nodes that are in the channel.

- Channel은 접근 정책이 설정 되어 있음
  - 따라서 privacy와 confidentiality는 채널 안에 있는 노드들만에 의해 보존 됨
  - 허락된 노드 만이 접근 가능
- Channels achieve better quality of robustness when a node is down given alternate paths to get to the destination, while also providing scalability allowing for effective sharing of huge amounts of data.
  - Channel은 대체 경로를 제공하기 때문에 노드가 다운 되어도 좋은 퀄리티의 강인함을 보여줌
- From a privacy perspective, channels are useful in cases where a subgroup of the blockchain network’s participants have a lot of transactions in common (enough to justify the creation of a new broadcast order channel), and these transactions can be processed with no dependency on state controlled by entities outside this group.
  - Privacy관점에서도, channel은 블록체인 네트워크의 함여자의 하위 그룹이 매우 많은 양의 트랜젝션을 공통으로 가지고 있을때 매우 유용 함.
    - 이 거래들은 이 그룹 외부에서 통제되는 state에 의존하지 않고 처리 될 수 있음

***Channels can be further used in combination with private transactions and zero-knowledge proof technologies***

## Private transactions in Hyperledger Fabric

Private transactions offer transaction privacy at a more fine-grained level than channels.

- Private transactions은 channel보다 더 향상된 privacy를 제공 함

> “Verified. Me by Secure-Key Technologies relies on strong privacy and confidentiality requirements. ***The solution depends on minimizing data disclosure and retention to appropriate parties, while still retaining the evidence of actions taken in the network.*** This balance is fundamental, and the ***foundation is provided in Hyperledger Fabric with the combination of Channels and Private Transactions.*** Watch the video to understand the real business value at work utilizing Hyperledger Fabric technology. Greg Wolfond, CEO, Secure-Key Technologies“

- Solution은 data가 들어나는 정도를 최소화하고, 적절한 파티를 유지하면서, 여전히 네트워크에 거래에 대한 증거를 남겨야 함
- Hyperledger fabric에서 Channel과 Private transaction을 통해 제공 함.
- ***By video, ID is worthier than money***

The database storing the private data is updated alongside the public ledger as transactions containing references to private data are committed.

- In fact, the hashes on the public ledger serve as verifiable proof of the data.
- ***Private transactions can be combined with anonymous client authentication to avoid leaking the connection between the identity of the transaction’s creator and the ledger stored (hashed) data.***
  - 거래 생성자의 identity와 ledger에 저장된 데이터 사이의 연관성이 유출 되는걸 막기 위해 private transactions은 익명의 사용자 인증과 결합 될 수 있음. 
- This feature is especially useful in cases where, for regulatory or legal reasons, private data is not allowed to reside off the premise(전제) of the parties involved in the transaction.
- 이 특성은은 규제 또는 법적 이유로 개인 데이터가 거래 당사자의 전제에서 벗어날 수 없는 경우에 특히 유용함.
  - Consider an example in the healthcare sector where a patient’s health information should be released only for a specified amount of time.
    - Healthcare 같은 특정 시간 동안만 환자의 정보가 공개되어야 하는 예를 고려해보자.
  - For example, a patient’s prescription history may be made available to a specialist for a period of time before a specific surgery occurs. 
    - 예를 들어 의사는 환자의 수술 전 환자의 처방기록에 대하여 조회가 가능해야 한다.
  - Private transactions would ensure data confidentiality in only allowing the patient and the specialist to see the information for a specified amount of time while also recording the hash of the data as evidence that the transaction occurred.
    - Private transaction은 오직 환자와 의사가 특정한 양의 시간 만큼만 정보를 보는것을 허락 할 하고, transaction이 발생 했을때 데이터의 해시를 증거로 저장하여 데이터의 기밀성을 보장 할 것이다. 
- Privacy is achieved in that there is control over who can access the actual sensitive data.
  - If anonymous client authentication is used in addition to this, stronger privacy would be offered as the identity of the entity that introduced or updated the (hashed) record will also be concealed.
    - Privacy는 실제 민감한 데이터 접근에 통제가 되어야 이루어 질 수 있다.
    - 만약 익명의 사용자 인증이 추가로 사용된다면, 감춰진 해시된 기록을 도입하거나, 업그레이드 한 독립된 인증으로 제공 될 것이기 때문에 더 강력한 프라이버시가 제공 될 것이다. 

