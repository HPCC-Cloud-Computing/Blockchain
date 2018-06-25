# Hướng dẫn xây dựng blockchain đơn giản (Golang)

(xin phép thầy đặt phần này ở đầu để các bạn dễ thấy)
--> Bích tạo repo mới cho nhóm và quản lý bên đó nhé. Repo này chỉ để meeting log và tài liệu chia sẻ.

* [Repo](https://github.com/Jeiwan/blockchain_go)

* [Documentation](https://jeiwan.cc/posts/building-blockchain-in-go-part-1/)

# Nội dung các buổi họp nhóm
  
## Meeting 22/06/2018

### Nội dung:
    1. Huy (nhóm 4) trình bày chương 5
    2. Hùng (nhóm 5) trình bày chương 6
    3. Huyền (nhóm 2) trình bày chương 7

### Câu hỏi:
        1. Giao dịch offline được hiểu như thế nào?
        2. Khi lan truyền các giao dịch, các nút có xác thực nội dung giao dịch hay không? Việc xác thực thông tin giao dịch được thực hiện ở đâu?
        3. Làm rõ transaction input và transaction output
        4. Cách tính phí giao dịch
        5. Thảo luận về các loại nút trong mạng bitcoin
  6. Thảo luận về Bloom Filter, độ phức tạp của thuật toán
  7. Một SPV node xác minh sự tồn tại của 1 giao dịch thuộc block bằng Merkle path như thế nào?

### Cuộc thi Hackathon
  1. Các mảng đề tài tham gia cuộc thi
  2. Thầy Chung đưa ra 1 số đề tài tham khảo: xây dựng hệ thống quản lí dân cư, xây dựng blockchain riêng (dùng Go)
  3. Các nhóm thảo luận đề tài
      - Nhóm chẵn: Code blockchain
      - Nhóm lẻ: Tạo ứng dụng trên HyberLegde

### Buổi họp tiếp theo: 9h thứ 6, ngày 29/6/2018

## Meeting 15/06/2018

### Nội dung
1. Luật trình bày chương 2 &3 cuốn mastering Bitcoin
2. Trung Anh trình bày chương 4
3. Chia nhóm tham gia Hackathon https://www.facebook.com/events/472684833186093/

### Câu hỏi
1. Quy định 10 phút như thế nào? Nhật sẽ gửi bài báo
2. Phần code nào của bitcoin quy định 10 phút? 
3. Chọn giao dịch dựa trên cơ chế nào? nhiều tiền phí hay đến trước sau?
4. Tại sao dùng ví thay cho địa chỉ public key
5. Có cần ẩn public key đi không?
6. Dùng phép XOR hay AND khi kết hợp parent public key và child key?

### Buổi tiếp theo
+ 9h thứ sáu tuần sau (22/6)
+ 3 chương tiếp theo của quyển mastering Bitcoin
