# Design Patterns - Mẫu Thiết Kế

## Giới thiệu
Design Patterns (Mẫu thiết kế) là các giải pháp đã được chuẩn hóa để giải quyết các vấn đề phổ biến trong thiết kế phần mềm. Chúng cung cấp cách tiếp cận tái sử dụng, giúp mã nguồn dễ bảo trì, mở rộng và hiểu hơn. Tài liệu này giới thiệu một số mẫu thiết kế phổ biến và cách áp dụng chúng.

## Mục tiêu
- Cung cấp tổng quan về các mẫu thiết kế cơ bản.
- Hỗ trợ lập trình viên áp dụng các mẫu này vào dự án thực tế.
- Tăng tính linh hoạt và hiệu quả trong quá trình phát triển phần mềm.

## Các loại Design Patterns
Design Patterns thường được chia thành 3 nhóm chính theo sách *"Design Patterns: Elements of Reusable Object-Oriented Software"* của Gang of Four (GoF):

### 1. Creational Patterns (Mẫu tạo lập)
- Tập trung vào việc tạo đối tượng một cách linh hoạt.
- Ví dụ:
  - **Singleton**: Đảm bảo chỉ có một instance duy nhất của một lớp.
  - **Factory Method**: Tạo đối tượng mà không cần chỉ định lớp cụ thể.
  - **Abstract Factory**: Tạo tập hợp các đối tượng liên quan mà không phụ thuộc vào lớp cụ thể.

### 2. Structural Patterns (Mẫu cấu trúc)
- Xác định cách tổ chức các lớp và đối tượng để hình thành cấu trúc lớn hơn.
- Ví dụ:
  - **Adapter**: Chuyển đổi giao diện của một lớp để tương thích với lớp khác.
  - **Decorator**: Thêm tính năng cho đối tượng mà không cần thay đổi mã nguồn.
  - **Facade**: Cung cấp giao diện đơn giản cho một hệ thống phức tạp.

### 3. Behavioral Patterns (Mẫu hành vi)
- Quản lý cách các đối tượng tương tác và phân chia trách nhiệm.
- Ví dụ:
  - **Observer**: Định nghĩa cơ chế để thông báo nhiều đối tượng khi trạng thái thay đổi.
  - **Strategy**: Cho phép thay đổi thuật toán linh hoạt trong thời gian chạy.
  - **Command**: Đóng gói yêu cầu dưới dạng đối tượng để thực hiện hoặc hoàn tác.

## Cách sử dụng
1. **Xác định vấn đề**: Hiểu rõ vấn đề thiết kế mà bạn đang gặp phải.
2. **Chọn mẫu phù hợp**: Dựa trên ngữ cảnh, chọn mẫu thiết kế thích hợp từ danh sách trên.
3. **Triển khai**: Áp dụng mẫu vào mã nguồn, tuân theo nguyên tắc SOLID và các phương pháp lập trình tốt.
4. **Kiểm tra**: Đảm bảo giải pháp hoạt động đúng và không làm phức tạp hệ thống.

## Tài liệu tham khảo
- *"Design Patterns: Elements of Reusable Object-Oriented Software"* - Erich Gamma et al.
- Các bài viết và ví dụ trên mạng về Design Patterns (GoF).
