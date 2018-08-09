# 1. INTRODUCATION

- 1 blockchain có thể định nghĩa là 1  sổ cái bất biến, để ghi lại các giao dịch đc duy trì trong 1 mạng phân tán của các peer ko tin cậy lẫn nhau.
- Các peer thực hiện 1 giao thức đồng thuận để xác thực giao dịch.
-  1 blockchain được phép có thể sử dụng đồng thuận chịu lỗi BFT
  - Nhiều ứng dụng phân tán chạy đồng thời
  - Các ứng dụng có thể được triển khai động với bất kì ai
  - Mã ứng dụng không đáng tin cậy, thậm chí có thể độc hại
- Kiến trúc thực thi lệnh
- Tất cả các peer thực hiện mọi giao dịch và giao dịch phải xác định
- Hạn chế của các kiến trúc blockchain trước:
  - Sự đồng thuận bị mã hóa cứng trong nền tảng
  - Mô hình tin cậy của xác thực giao dịch được xác định bởi giao thức đồng thuận và không thể thích ứng với các yêu cầu của hợp đồng thông minh
  - Hợp đồng thông minh phải được viết bằng ngôn ngữ cố định, cản trở việc áp dụng rộng rãi và có thể dẫn đến lỗi lập trình
  - Việc thực thi tuần tự của tất cả các giao dịch đều hạn chế hiệu suất và các biện pháp phức tạp là cần thiết để ngăn chặn các cuộc tấn công từ chối dịch vụ đối với nền tảng có nguồn gốc từ các hợp đồng không đáng tin cậy
  - Giao dịch phải xác định, có thể khó đảm bảo bằng chương trình
  - 
- Fabric kết hợp 2 pp tiếp cận nổi tiếng:
  - Sử dụng sao chép thu động hoặc chính như trong csdl phân tán, nhưng với xử lý cập nhật ko đối xứng dựa trên middleware và được chuyển đến các mt ko tin cật với lỗi BFT. Trong fabric mọi giao dịch được xác thực bởi tập hợp các peer, cho phép thực hiện song song và giải quyết các tiềm năng ko xác định. Thực hiện chính xác 1 hợp đồng thông minh
  - Fabric kết hợp hoạt động mở rộng theo các tác động của giao dịch trên trạng thái sổ cái chỉ đươc đưa vào khi đạt được sự đồng thuận về tổng số thứ tự , trong bước xác thưcj được thực hiện bởi các peer
- Fabric chứa khối xây dựng mô-đun  có các thành phần:
  - 1 service order phát trạng thái cập nhật peer và thiết lập sự đồng thuận về thứ tự giao  dịch
  - 1 msp chịu trách nhiệm liên kết các peer với các mã hóa. Duy trì bản chất của fabric
  - 1 service message ngang hàng tùy chọn các khối đầu ra cho tất cả các peer.
  - Các smart-contract chạy trong mt container để cách ly, được viết bằng ngôn ngữ lập trình chuẩn nhưng ko có quyền truy cập trực tiếp trạng thái sổ cái.
  - Mỗi peer duy trì 1 sổ cái dưới dạng blockchian chỉ cập nhât thêm (key-value)

# 2. BACKGROUND

## 2.1 Order-Execute Architecture for Blockchains

- Hệ thống blockchain trước đây tuân theo kiến trúc thực hiện lệnh.

## 2.2 **Limitations of Order-Execute** 

- **Thực hiện tuần tự:** 
  - Thực hiện giao dịch tuần tự trên các peer giới hạn thông lượng có thể đạt được trong blockchain.
  - Có khả năng bị tấn công DoS.
  - Giải pháp: Tạo chi phí giao dịch
  - 
- **Mã  không xác định:**
  - Các giao dịch không xác định, các hoạt động được thực hiện sau khi đồng thuận trong SMR phải được xác định (các peer lưu trữ trạng thái giống nhau)
  - Chỉ 1 smart contract ko xác định được tạo ra với mục đích độc hại là đủ để đưa toàn bộ hệ thống blockchain dừng lại
- **Bảo mật thực thi:** 
  - public blockchain cho phép chạy tất cả smart contract trên tất cả các peer.
  - Các biện pháp mã hóa, từ mã hóa có chi phí cao hoặc ko thực thi được.

## 2.3 **Further Limitations of Existing Architectures**

# 3 **ARCHITECTURE**

## **3.1	Fabric Overview** 

- Thực hiện các ứng dụng phân tán.
- Theo dõi lịch sử thực thi 1 cách an toan trong cấu trúc dữ liệu ledger.
- Gồm 2 phần:
  - Chaincode: Phát triển bởi dev ko tin cậy
  - Một chính sách chứng thực được đánh giá trong giai đoạn xác nhận: Ko thể chọn hoặc bị sửa đổi bởi dev ko tin cậy, như 1 thư viện tĩnh để xác thực các giao dịch.
- _Sau khi thực hiện, các giao dịch được đưa vào order, sử dụng giao thức đồng thuận **plug-gable** để tạo ra một chuỗi các giao dịch đã được xác nhận và được nhóm theo khối.  Mỗi peer sau đó xác nhận các thay đổi trạng thái từ các giao dịch được xác nhận  với chính sách chứng thực và tính nhất quán của việc thực thi_
- Fabric giới thiệu mô hình lai mới trong mô hình Byzantine: Kết hợp sự sao chép thụ động (tính toán đồng thuận trước các cập nhật trạng thái) và nhân rộng hoạt động (xác nhận sau khi kết quả đồng thuận thực hiện và thay đổi trạng thái).
- **Order hoàn toàn không biết về trạng thái đơn đăng ký và không tham gia thực hiện cũng như trong quá trình xác thực giao dịch.**

## **3.2	Execution Phase** 

- Client ký và gửi giao dịch cho các client khác xác nhận
- Yêu cầu -> Xác nhận -> Order -> Xác thực
- Trạng thái của blockchain được tạo ra bởi chaincode có phạm vi độc quyền cho chaincode đó và không thể truy cập trực tiếp từ 1 chaincode khác. (**GetState, PutState và DelState**)
- 1 bản ghi giá trị là các cập nhật trạng thái
- Thực hiện 1 transaction trước giai đoạn gửi đến order là rất quan trọng
- Chấp nhận các thực thi không xác định sẽ giải quyết đc cuộc tấn công DoS từ chuổi mã ko tin cậy 

## **3.3	Ordering Phase** 

- 1 client khi đã tập hợp đủ xác nhận, nó sẽ lắp ráp transaction và gửi đến cho order
- Transaction: Tham số, metadata giao dịch và 1 tập hợp xác nhận.
- Order thiết lập số transaction đóng vào block và thiết lập sự đồng thuận về giao dịch
- Order đảm bảo các khối được phân phối trên kênh, đảm bảo thứ tự các thuộc tính.
- Fabric cấu hình để sử dụng dịch vụ gossip để đưa các khối đến tất cả các peer

 **Fabric trở thành hệ thống blockchain đầu tiên hoàn toàn tách biệt khỏi việc thực hiện và xác nhận hợp lệ**

##  **3.4	Validation Phase** 

- Các block sẽ được gửi đến các peer thông qua order hoặc gossip.
- Quá trình xác thực:
  -  Một thư viện tĩnh là một phần của cấu hình blockchain và chịu trách nhiệm xác nhận chứng thực đối với chính sách xác nhận mà được cấu hình cho chaincode và quá trình chứng thực này xảy ra song song cho tất cả giao dịch trong khối.
  - Kiểm tra xung đột đọc-ghi được thực hiện cho tất cả các giao dịch trong khối tuần tự.
  - Giai đoạn cập nhật sổ cái.

## **3.5	Trust and Fault Model** 

- Các peer được nhóm lại thành các tổ chức và tạo thành 1 miền tin cậy
- Tính toàn vẹn của mạng Fabric phụ thuộc vào sự nhất quán của dịch vụ order

# 4 **FABRIC COMPONENTS** 

Fabric viết bằng Go và sử dụng giao thức gRPC để định hướng các chức năng cần gọi

## **4.1	Membership Service**

- MSP duy trì danh tính của tất cả các node trong hệ thống và cấp chứng chỉ node để xác thực và ủy quyền.(chữ ký số).
- Nhiệm vu: Xác thực giao dịch, xác minh tính toàn vẹn của các giao dịch, ký và xác thực tất cả các xác nhận và các hoạt động blockchain khác. Bao gồm cả quản lý khóa và đăng ký các node.
- MSP cho phép liên kết danh tính (khi nhiều tổ chức hđ trong 1 mạng blockchain)

## **4.2	Ordering Service** 

- Được khởi động với 1 khối genesis trên kênh hệ thống, khối này mang giao dịch cấu hình.
-  Việc đóng khối khi: 
  - Khối chứa số lượng giao dịch tối đa
  - Khối đã đạt kích thước tối đa
  - Qua 1 khoảng time chỉ định

## **4.3	Peer Gossip** 

- Peer gossip sử dụng multicast cho mục đích chuyển giao trạng thái cho các peer mới và peer đã bị ngắt kết nối trong time dài. Khi nhận được tất cả các khối, peer độc lập cập nhật blockchain của nó và xác minh tính toàn vẹn dữ liệu.
- Giao tiếp trong gossip dựa trên gRPC và sử dụng TLS.
- Sử dụng 2 giai đoạn để phổ biến thông tin:
  - Quá trình push: Mỗi peer chọn ngẫu nhiên 1 tập hợp các hàng xóm đang hoạt động và chuyển tiếp thông báo.
  - Qua trình pull: Mỗi peer ngang hàng sẽ thăm dò 1 nhóm các peer được chọn ngẫu nhiên và yêu cầu gửi các message bị thiếu

## 4.4 Ledger

- Các block là bất biến và theo 1 trật tựu nhất định.

## **4.5	Chaincode Execution** 

- Chaincode được thực thi trong môi trường ghép nối ko ổn định với phần còn lại của peer
- Peer là không thể hiểu đươc ngôn ngữ cái mà chaincode đang thực hiện
- Chaincode chạy trong môi trường container Docker.

## **4.6	Configuration and System Chaincodes** 

- fabric được tùy chỉnh thông qua cấu hình kênh và thông qua các chaincode đặc biệt gọi là **system chaincode**
- Cấu hình kênh:
  - Định nghĩa MSP cho các nút tham gia.
  - Đại chỉ mạng của OSN.
  - Cấu hình chia sẻ cho việc thực hiện đồng thuận và dịch vụ
  - Quy tắc điều chỉnh quyền truy cập vào các dịch vụ order
  - Quy tắc điều chỉnh cách cấu hình kênh có thể được sửa đổi
- Cấu hình của kênh được cập nhật bằng cách sử dụng giao dịch cập nhật cấu hình kênh.(chứa những thay đổi và tập hợp các chứ ký)
- 

# **5	EVALUATION** 

- **fabric** chưa được điều chỉnh và tối ưu hóa hiệu suất
- **fabric** là 1 hệ thống phân tán phức tạp
- Hiệu suất **fabric** phụ thuộc vào:
  - Lựa chọn ứng dụng phân tán
  - Kich thước giao dịch
  - Dich vụ order , sự đồng thuận và tham số
  - Thông số mạng, cấu trúc liên kết của các node trong mạng
  - Phần cứng trên node nào chạy
  - Số node và kênh
  - Thông số cấu hình và mạng động lực học (?)

## **5.1	Fabric Coin (Fabcoin)** 

## **5.2	Experiments** 

# **6	APPLICATIONS AND USE CASES** 

# 7  **RELATED WORK**