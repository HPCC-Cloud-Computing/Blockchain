# Mastering Bitcoin

#### 1. Introduction
* BTC không tồn tại dưới dạng tiền vật lý cũng như tiền điện tử. Nó được ngầm định qua các transaction.
* Người nào có key sẽ có thể unlock được những **transaction output**, và tiêu số BTC trong các transaction đó.
* Hoàn toàn phân tán, mạng lưới P2P, không có trung gian.
* BTC được tạo ra bằng việc **mining**, cứ 10 phút 1 lần.
* Đào thành công được thưởng một lượng BTC, phần thưởng này giảm 1 nửa sau mỗi 4 năm. Tính đến 2140 sẽ bằng 0, số lượng BTC tối đa là 21 triệu.
* Bitcoin cũng là tên một giao thức. Đồng BTC chỉ là một ứng dụng của nó.
* Bitcoin gồm:
	* Mạng lưới P2P phi tập trung.
	* Sổ cái công khai (Blockchain).
	* Tính toán phi tập trung.
	* Xác thực transaction phi tập trung.

#### 2. How bitcoin Works
###### VD mua cà phê: Alice gặp Joe đổi tiền lấy bitcoin. Alice mua cafe ở chỗ Bob bằng BTC

* Joe tạo tx để chuyển cho Alice BTC.

* Alice quẹt mã QR. Ứng dụng tạo transaction, giao dịch thành công.

* 1/100_000_000 BTC = 1 Satoshi

* Transaction:
  * Cho biết số lượng bitcoin được chuyển từ địa chỉ này tới địa chỉ khác.
  * Người nhận dùng transaction này để tạo transaction mới, chuyển BTC tới địa chỉ khác, tạo thành 1 chuỗi
  * Bao gồm các **input** - debits, và các **output** - credits chuyển tới cho địa chỉ khác. **input** thực chất chính là các **output** từ những tx trước đó.

* Yêu cầu ứng dụng của Alice phải lựa chọn những input phù hợp để xây dựng tx mới.

* Các ứng dụng ví thường có 1 database **unspent tx** của key phù hợp (lightweight clients).

* Output sẽ chứa 1 script để thể hiện số BTC này được chuyển cho Bob. Chỉ Bob có key phù hợp mới có thể dùng số lượng BTC này.

* Alice trả 0.1BTC, nhưng cốc cafe chỉ trị giá 0.015BTC. 2 output sẽ được tạo ra, 1 cho Bob, 1 trả lại tiền thừa cho Alice

* Phí giao dịch: tiền thừa là 0.085, nhưng ví sẽ chỉ tạo output với 0.0845 BTC, số còn lại - 0.005 BTC (chênh lệch input và output 0.1 - 0.015 - 0.0845) sẽ trở thành phí giao dịch.

* Mạng lưới Bitcoin sẽ tiến hành xử lý giao dịch này. Giao dịch được thêm vào blockchain.

* Ví của Alice không cần kết nối trực tiếp tới ví của Bob, cô sẽ dùng mạng Internet để gửi tx này cho các node khác, chúng sẽ xác thực và tiến hành xử lí. Bob sẽ nhận được thông báo vì tx output trên chứa key của Bob.

###### Mining

* Một tx sẽ không được thêm vào sổ cái nếu như nó chưa được xác thực và đào (mining)
  * `Mining` - đồng thuận: minh chứng cho việc block đã được tính toán trong một khoảng thời gian, càng nhiều block thì độ tin tưởng càng cao.
  * `Proof Of Work` - tính mã băm của header với một số ngẫu nhiên để ra được giá trị cần tìm. Hoàn thành nó đồng nghĩa với việc `Mining` hoàn tất, block mới được thêm vào. Đây là một cuộc đua, người tìm ra sớm nhất kết quả này sẽ là người chiến thắng và sẽ đưa chúng lên mạng lưới để xác thực.
  * Càng nhiều người đào, độ khó của việc `Mining` càng cao. Sẽ cần những thiết bị chuyên dụng như GPU. Hay các **mining pool** để chia sẻ công việc.
* Những tx mới sẽ được lưu trong một **pool** tại nội bộ mỗi node. Các pool này có thể khác nhau tùy node. Khi tiến hành `Mining`, node sẽ lấy tx từ pool này. Khi đào, miner sẽ tự động thêm 1 tx đặc biệt để gửi cho chính mình 25 BTC (phần thưởng từ việc đào).
* Đối với 1 tx trong block, cứ một block được thêm vào sau block ấy sẽ gọi là một **confirmation**. Vì vậy, càng nhiều **confirmation** thì sẽ càng đáng tin cậy.
* Hiện tại với BTC, nếu có 6 **confirmation** thì tx sẽ được xem như là giao dịch thành công.

###### Tiêu transaction

* Full node lưu đầy đủ blockchain nên có thể truy xuất được nguồn gốc của đồng BTC đó. Tuy nhiên với lightweight node, nó sẽ thực hiện `Simplified Payment Verification` để xác nhận tx này có trong blockchain và đã có một vài block sau nó (hợp lệ).

#### 3. The Bitcoin Client
###### Bitcoin Core

* Để thao tác với BTC, ta cần Bitcoin Core
  * Với người dùng bình thường có thể download bộ cài từ bitcoin.org.
  *  Với lập trình viên, có thể build từ source code trên github, sau đó thao tác qua CLI. Dữ liệu trả về dưới dạng JSON.

```shell
$ bitcoin-cli <args>
```
#### 4. Keys, Addresses, Wallets

###### Introduction

* Để chứng minh quyền sở hữu đối với BTC cần: chìa khóa - **key**, địa chỉ - **address**, chữ kí - **signature**. Chúng được lưu giữ lại máy của từng node (gọi là ví - **wallet**).
* Tx nào cũng cần có một signature hợp lệ. Signature này được tạo bởi key của người dùng.
* Key gồm: **public key** và **private key** (cần lưu trữ cẩn thận).
* Khi giao dịch, thường sẽ dùng **bitcoin address** (được mã hóa từ public key) thay cho public key.

###### Public key cryptography and crypto-currency

* Bitcoin sử dụng **elliptic curve** trong việc tạo key.
  * Public key: dùng để nhận bitcoin (nằm trong tx output).
  * Private key: dùng để tạo signature (nằm trong tx input) - tiêu bitcoin.
  * Public key có thể tính toán từ private key.
* Signature sinh bởi private key có thể được xác thực bằng public key mà không cần dùng đến private key.

###### Private and Public Keys

Các bước sinh key:

* Sinh ngẫu nhiên private key (1 số).
* Từ private key, dùng elliptic curve (1 chiều) sinh ra public key.
* Từ public key, sử dụng hàm băm một chiều sinh ra bitcoin address.
* Không có chiều ngược lại.

###### Private Keys

* Range: 1 - 2<sup>256</sup> (xấp xỉ 10<sup>77</sup>, trong khi số nguyên tử có trong phần vũ trụ nhìn thấy được là khoảng 10<sup>80</sup>)
* Không nên sử dụng hàm random của các NNLT, thuật toán này liên quan đến bảo mật của key.

###### Bitcoin Addresses

* Địa chỉ bitcoin được công khai, người khác sẽ giao dịch dựa trên địa chỉ này (địa chỉ bitcoin đóng vai trò là địa chỉ nhận).
* Được mã hóa từ public key thông qua các thuật toán SHA256 và RIPEMD160.
  * `public_key_hash = RIPEMD160(SHA256(public_key))`
  * `checksum = SHA256(SHA256(prefix + public_key_hash)).to_hex.slice(0, 4) # Lấy 4 byte đầu`
  * `address = Base58(prefix + public_key_hash + checksum)`
* Bitcoin Address được bắt đầu bằng "1", do sử dụng prefix 0x00.

###### Wallets

* Random Wallets:
  * Private key được sinh random
  * Phải lưu lại bản copy của nó, phải backup thường xuyên
  * Đây là một lựa chọn tồi.
  * Bitcoin Core Client sử dụng loại này, nhưng các Bitcoin developer không khuyến khích dùng.
* Deterministic Wallets:
  * Private key được sinh từ một seed phổ biến thông qua hàm băm một chiều.
  * Seed là một số ngẫu nhiên cùng với một data (`index number` hoặc `chaincode`).
  * Có thể sinh ra nhiều key khác nhau từ một seed.
* Mnemonic Code Words:
  * Các từ tiếng anh mô tả một con số ngẫu nhiên dùng trong deterministic wallet.
  * Một chuỗi các từ tiếng Anh là đủ để tái tạo lại tất cả các key.

#### Advanced Keys and Addresses

###### Encrypted Private Keys

* Private key phải được lưu trữ kĩ càng, riêng tư.
* BIP0038 đưa ra một phương án mã hóa private key với một passphrase (giống Linux ssh-agent ?), sau đó mã hóa Base58, kết quả này có thể được công khai.

###### Pay to Script hash (P2SH) and Multi-Sig Addresses.
* Bitcoin Address bắt đầu bằng "3" được gọi là P2SH (hoặc Multi-Signature Address).
* Bằng cách mã hóa với prefix 5, ta sẽ được address bắt đầu bằng "3".
* Không như tx gửi Bitcoin đến Bitcoin address "1...", nó sẽ gửi đến địa chỉ "3...", mà sẽ cần nhiều thông tin hơn ngoài public key hash và signature.
* Như tên gọi, nó cần nhiều hơn 1 signature để xác nhận, cụ thể là M (gọi là `threshold`) nhỏ hơn hoặc bằng N (N key). Từ đó N key kia đều có thể dùng được các transaction output này (tương đương tài khoản chung).

  VD: ông A chuyển 5 BTC cho ông B. Ông B sử dụng 1-of-2 multi-sig, dùng 2 key của ông và của vợ ông. Sau đó ông và vợ mình đều có thể sử dụng key của mình để tiêu transaction này.

###### Vanity Addresses

* Đây là địa chỉ bao gồm những từ có nghĩa. Nó yêu cầu phải thử hàng tỉ địa chỉ khác nhau.
* Tính bảo mật không khác gì các địa chỉ khác.

#### 5. Transactions

###### Transaction Lifecycle

* Tạo transaction
* Transaction được kí để tiêu số BTC trong tx
* Transaction được broadcast lên mạng lưới btc, mỗi node xác thực và gossip cho các node khác.
* Transaction được Miner xác thực và thêm vào trong block.
* Sau khi được chấp nhận (sau một vài block) thì transaction sẽ được coi là hoàn toàn hợp lệ, và bắt đầu một vòng tuần hoàn mới.

###### Tạo transaction

* Có thể tạo transaction ở bất cứ đâu, bất cứ node nào

* Nó chứa đầy đủ thông tin để có thể thực hiện việc giao dịch.

###### Broadcasting

* Tx không chứa các thông tin nhạy cảm, có thể public cho toàn bộ, nên việc tin tưởng người broadcast các tx là không cần thiết, các tx có thể gửi bằng bất cứ kết nối mạng nào.

###### Propagating

* Node sẽ kết nối tới 1 vài node xung quanh để bắt đầu việc gửi tx.
* Nếu tx hợp lệ, các node xung quanh đó lại gửi đến 1 vài node xung quanh chúng tiếp, dần dần sẽ hết mạng lưới.

###### Transaction Structure

* Transaction dùng input, ouput để xác định người gửi, và người nhận cũng như số tiền gửi, và nhận.

* Cấu trúc:

| Kích thước | Trường    |
| ---------- | --------- |
| 4 bytes    | Version   |
| 1-9 bytes  | Số input  |
| Variable   | Các input |
| 1-9 bytes  | Số output |
| Variable   | Output    |
| 4 bytes    | Timestamp |

###### Transaction Outputs and Inputs

* Output:
  * Các output có thể tiêu được gọi là unspent transaction output (UTXO).
  * Các UTXO này được suy ra từ blockchain.
  * Các ứng dụng ví btc tính toán số dư của một tài khoản bằng cách tính tổng số BTC trong các UTXO nằm rải rác trong blockchain.
  * UTXO chứa thông tin về số lượng tiền giao dịch, mỗi đơn vị gọi là 1 Satoshi.
  * Các UTXO không thể chia nhỏ. Nếu lớn hơn số tiền cần giao dịch, nó sẽ tự động gửi lại số dư cho ta.
* Input:
  * Các UTXO được sử dụng trong mỗi tx được gọi là tx input.
  * Như vậy, transaction sinh ra UTXO. Ta dùng UTXO này để tiêu, nó sẽ trở thành các input, và tạo ra UTXO mới.
* Có một ngoại lệ là **coinbase tx**, tx này dành cho miner đã bỏ công sức ra để đào block, nó không có input mà chỉ có output là phần thưởng (hiện tại là 12.5 BTC) gửi đến địa chỉ của miner.

###### Transaction Output

* UTXO được track bởi các full-node, được gọi là UTXO set hoặc UTXO pool.
* Output gồm:
  * Lượng Satoshi cần giao dịch
  * Locking script: người giữ public key phù hợp sẽ có thể mở khóa các output này.

###### Transaction Input

* Mỗi transaction input sẽ trỏ đến một transaction output nào đó.
* Input cũng có 1 script để xác minh việc tiêu output nó trỏ đến là hợp lệ. Script liên quan đến signature (sinh bởi private key).

###### Transaction Fees

* Phí giao dịch được trả cho miner.
* Phí giao dịch được tính dựa trên dung lượng của transaction (tính theo KB).
* Miner có thể sắp xếp các transaction trong transaction pool dựa trên fee, cũng có thể làm free.

###### Adding fees to Transactions

* **Fees = Sum(Input) - Sum(Output)**
* VD transaction càng nhiều input thì kích thước càng lớn.

###### Transaction Chaining and Orphan Transactions

* Transaction tạo output, output lại trở thành input của transaction mới => Chain
* Coi transaction tạo output = parent, transaction dùng output đó = child. Đôi lúc, child xuất hiện trước parent (do độ trễ?), thay vì loại bỏ chúng thì các node sẽ cho vào một temp pool. Để tránh DDoS, số orphan transaction này bị giới hạn.

###### Transaction Scripts and Script Language

* Transaction được xác thực bằng các script (viết bằng Script language). 
* Bitcoin dùng Pay-to-Public-Key-Hash script.

###### Script Construction (Lock + Unlock)

* Locking script: được đặt trong các output. Gọi là `scriptPubKey`. Nó chứa public key hoặc btc address.
* Unlocking script: được đặt trong các input. Gọi là `scriptSig`.  Nó chứa signature được tạo bởi private key.
* Validate: nối 2 script này lại. `script = scriptSig + scriptPubKey`.

###### Scripting Language

* Được gọi là **Script**.
* Sử dụng cơ chế Stack.

###### Stateless Verification

* Trước và sau khi chạy script đều không lưu lại trạng thái, mọi thứ cần thiết cho script này đều nằm trong script.
* Đều cho một kết quả ở mọi hệ thống.

###### Standard Transactions

* Có 5 kiểu tiêu chuẩn cho script: P2PKH, Public-Key, Multi-Sig, P2SH, và OP_RETURN.

#### 6. The Bitcoin Network

###### Peer-to-Peer Network Architecture

* Tất cả các node đều ngang bằng nhau, không có node nào đặc biệt hơn node nào.
* Không có server chung, không trung gian, không phân cấp bậc.

###### Nodes Types and Roles

* Các node trong mạng lưới btc có thể có vai trò (role) khác nhau.

* Btc node: 

  | Role                | Full node | Lightweight node |
  | :------------------ | :-------: | :--------------: |
  | Routing             |    Yes    |       Yes        |
  | Blockchain database |    Yes    |        No        |
  | Mining              | Optional  |     Optional     |
  | Wallet services     | Optional  |     Optional     |

###### The Extended Bitcoin Network

* Các node:
  * Chạy các phiên bản Client khác nhau (Bitcoin Core).
  * Dùng các giao thức P2P khác nhau (BitcoinJ, Libbitcoin, btcd).
  * Nhiều role khác nhau.
* Dựa trên nền Bitcoin Core, có thể tạo ra rất nhiều ứng dụng khác: exchanges, wallets, block explorers, ...

###### Network Discovery

Khi một node mới xuất hiện:

* Nó phải kết nối tới các node khác xung quanh (ít nhất là 1).

* Thường sử dụng cổng 8333 TCP.
* Không có node đặc biệt, nhưng một số node ổn định được list lại: **seed nodes**.
* Khi kết nối thành công, nó gửi IP đến các node xung quanh, các node xung quanh lại gửi ra các node xung quanh khác.
* Nó cũng có thể gửi thông điệp để lấy danh sách các peer của neighbor nó.

###### Full Nodes

* Chứa toàn bộ blockchain (Full blockchain nodes).
* Có thể build, verify một cách độc lập, nhưng vẫn phải kết nối tới network để update block mới.

###### Exchanging "Inventory"

* Xây dựng blockchain là điều đầu tiên mà full node cần làm.
* Sync dựa trên `BestHeight`.
  * Node có blockchain có thể nhận biết được đối phương cần phải thêm block nào để có thể đuổi kịp, sau đó nó sẽ gửi 500 block để share.
  * Node nhận có thể track số block mỗi peer gửi (đã request mà chưa nhận được, hay vượt quá số block cho phép - định nghĩa trong `MAX_BLOCKS_IN_TRANSIT_PER_PEER`).

###### Simplified Payment Verification (SPV) Nodes

* Các SPV Node có thể thao tác với blockchain mà không cần lưu trữ toàn bộ blockchain.
* Nó chỉ lưuduy nhất block header mà không lưu các transaction (nhỏ hơn khoảng 1000 lần full blockchain).
* SPV không có cái nhìn toàn cảnh về UTXO.
* Nó dựa vào các peer khác để thực hiện verify các transaction.
* SPV verify transaction dựa trên **depth** của blockchain thay vì **height**.
  * Full node verify chuỗi các block và transaction.
  * SPV node chỉ verify chuỗi các block, liên kết chúng tới các giao dịch liên quan.
  * VD để kiểm nghiệm 1 transaction ở block thứ 300_000 
    * Full node sẽ duyệt toàn bộ blockchain, build UTXO db sau đó kiểm tra các input là hợp lệ.
    * SPV node sẽ tạo một liên kết giữa transaction và block chứa nó, sử dụng Merkle Path. Sau đó đợi thêm 6 block nữa (300_001 - 300_006), rồi verify. Nếu đã có 6 block sau đó, chứng tỏ mạng lưới đã chấp nhận rằng transaction đó là hợp lệ (không double-spend).
* SPV node có thể nhận biết transaction có tồn tại trong block hay không, nhờ (request) Merkle Path, và xác thực bằng chuỗi PoW (6 block).
  * Tuy nhiên SPV không thể kiểm tra được việc transaction đó sử dụng chung UTXO.
  * Để tránh loại tấn công này, SPV node cần kết nối tới một vài node ngẫu nhiên.

###### Bloom Filters

* Là một **probabilistic search filter** (Search xác xuất?), giúp mô tả 1 pattern mà không cần nêu chính xác nó.
* Dùng cho SPV khi yêu cầu các peer khác để lấy transaction ứng với một pattern.

###### Bloom Filters and Inventory Updates

* Bloom filter được dùng để lọc các transaction (và block chứa chúng), được SPV sử dụng khi gửi request tới các peer.
  * Khi phản hồi, peer sẽ gửi các block header ứng với các block tìm được, cùng với Merkle Path cho mỗi transaction tìm được.

###### Transaction Pool

* Lưu trữ các transaction chưa được xác nhận.
* Ngoài ra, một số node còn lưu trữ các orphaned transaction (input tham chiếu tới output không tồn tại).
* Chỉ lưu trên local storage, không lưu trên persistent storage (?) - page 160

###### Alert Messages

* Dùng để gửi tín hiệu cho tất cả các node khi gặp sự cố nghiêm trọng.

#### 7. The Blockchain

###### Introduction

* Cấu trúc dạng link list ngược (dùng con trỏ prev).

* Sử dụng LevelDB.

* **Height**: khoảng cách so với block 1. **Top** hoặc **tip**: block mới nhất.

* Mỗi block được xác định bởi mã hash (SHA256). Mỗi block cũng tham chiếu tới block trước nó. Block đầu tiên được gọi là **genesis block**

* Do mỗi block chứa hash của block trước nó, điều này làm ảnh hưởng tới mã hash của chính nó. Mã hash của nó sẽ bị thay đổi nếu block trước nó thay đổi.

  => Thay đổi 1 block sẽ làm thay đổi cả blockchain 

  => Tính toán từ đầu

###### Structure of a Blockchain

* Block gồm:
  * Header: metadata
  * Transaction counter
  * Transactions

###### Block Header:

* Version
* Previous Block Hash
* Merkle Root
* Timestamp
* Difficult Target
* Nonce

###### Block Identifier - Block Header Hash and Block Height

* ID = block hash = băm header 2 lần bằng SHA256. ID là duy nhất.
* Có thể sử dụng **height** thay cho block hash để xác định block. VD: genesis block có height 0.
* Không như block hash. **Height** không phải duy nhất (Khi các block tranh nhau để được network chấp nhận có height bằng nhau).

###### The Genesis Block

* Là block đầu tiên, được tạo tự động bằng bitcoin client.

###### Linking Blocks in the Blockchain

* Khi nhận block mới, node sẽ kiểm tra prev block hash của nó.
* Nếu trùng với hash block cuối cùng, nó sẽ thêm vào blockchain local trong máy của nó.

###### Merkle Trees

* Merkle Tree được dùng để tổng hợp các transaction trong 1 block.
* Merkle Tree trong btc dùng double-SHA256.
* Số lá phải là số chẵn. Nếu lẻ, nó sẽ tự động duplicate lá cuối cùng.
* Để kiểm tra 1 transaction có trong 1 block hay không, ta cần Merkle Path, thời gian 2O(logn).
* Merkle Tree không được lưu trong block.

###### Merkle Trees and SPV

* Khi cần kiểm tra một transaction
  * SPV sẽ gửi bloom filter cho các peer khác.
  * Peer gửi lại block header và merkle path tương ứng với transaction tìm được.
  * SPV sử dụng Merkle Path để kiểm tra transaction.
  * SPV sử dụng Block Header kiểm tra block.

#### 8. Mining and Consensus

###### Introduction

* Mining là quá trình thêm block vào blockchain, nó cũng ngăn chăn sự gian lận, double spend.
* Người đào thành công được trả công bằng btc:
  * Lượng btc thưởng khi có block mới.
  * Phí giao dịch.
* Quá trình đào ~ 10 phút.
* **Mining** phục vụ cho việc phi tập trung, mục đích chính không phải để sinh ra btc mới.

###### Bitcoin Economics and Currency Creation

* Thời kì đầu, phần thưởng cho block mới là 50 BTC, cứ 4 năm giảm 1 nửa.

###### De-centralized Concensus

* Câu hỏi: làm thế nào để các peer có thể đồng thuận với 1 kết quả  mà không cần trung gian?
* Gồm 4 quá trình chính:
  * Verify transaction độc lập.
  * Tập hợp các transaction để tạo block mới độc lập.
  * Verify block độc lập.
  * Lựa chọn độc lập (sức mạnh tính toán lớn nhất - PoW).

###### Independent Verification of Transactions

* Các node sẽ truyền tay nhau transaction, để lan rộng ra toàn bộ network.
* Trước khi nhận, sẽ cần verify transaction.
  * Transaction syntax, data structure
  * Input/Output có rỗng hay không.
  * Transaction size
  * Output value
  * Transaction không phải coinbase (input có hash = 0, N = -1)
  * scriptSig, scriptPubKey
  * ...

###### Mining Nodes

* Việc 1 block mới đc lan truyền đồng nghĩa với việc kết thúc cuộc thi "đào".

###### Aggregating Transactions into Blocks

* Sau khi validate transaction, nó sẽ được đưa vào tx pool
* Trong khi đào block n, nó tạo trước **candidate block** để chuẩn bị cho block thứ n+1.

###### Transaction Age, Fees, and Priority

* Transaction được ưu tiên theo **age** của UTXO trong input của nó.
* Các transaction được ưu tiên thêm vào block có thể được gửi mà không tốn phí.
* Độ ưu tiên = Sum (Input value * Input Age) / Transaction size
* Độ ưu tiên cao: lớn hơn **57_600_000**
  * 1 btc (100_000_000)
  * 1 ngày (144 block).
  * transaction 250 byte.
* 50KB trong danh sách transaction được dành cho transaction ưu tiên cao.
* Thứ tự:
  * Coinbase.
  * 50KB: Transaction ưu tiên cao
  * Các transaction khác (phí cao - thấp)
  * Nếu thừa dung lượng, có thể thêm 1 vài transaction không có phí.

=> Transaction còn lại ngày càng có độ ưu tiên cao hơn, các transaction không có phí cũng vậy.

* Không có timeout, tuy nhiên có thể mất nếu node đó restart máy (vì lưu trong bộ nhớ trong).
* Các ứng dụng ví thường thiết lập có thể tạo lại transaction với phí giao dịch cao hơn nếu quá 1 khoảng thời gian nhất định.

###### The Generation Transaction

* Là transaction đầu tiên trong block
* Còn gọi là **coinbase transaction**
* Chỉ có một input, gọi là coinbase.
* Hash: 0, Output index: 0xFFFFFFFF

###### Coinbase Data

* Ở genesis block: "The Times 03/Jan/ 2009 Chancellor on brink of second bailout for banks", được dùng như 1 thông điệp.
* Giờ thường dùng để chứa extra nonce.

###### Constructing the Block Header

| Size     | Field               |
| -------- | ------------------- |
| 4 bytes  | Version             |
| 32 bytes | Previous Block Hash |
| 32 bytes | Merkle Root         |
| 4 bytes  | Timestamp           |
| 4 bytes  | Difficult Target    |
| 4 bytes  | Nonce               |

* Các bước thực hiện:
  * Version: 2
  * Thêm prev hash
  * Tạo Merkle tree
  * Thêm timestamp
  * Tính target
  * Nonce: 0
* Sau đó có thể tiến hành đào.

###### Mining the Block