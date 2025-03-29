
Ý tưởng là làm sao để thêm hình mới (như tam giác) mà không phải đụng chạm gì đến hàm `printAreas` – nghe hay ho quá đúng không =)))

## Câu chuyện bắt đầu từ đâu? 

Ban đầu, thầy Trọng có một đoạn code tính diện tích hình tròn và hình vuông đưa cho Minh Thảo. Nhưng mà, nó dùng `if-else` trong hàm `printAreas` để tính diện tích dựa trên `shape.Type`. Vấn đề là nếu mình muốn thêm tam giác, mình phải sửa hàm `printAreas` – mà như thế thì vi phạm yêu cầuuu mất rồi! Thế là mình phải nghĩ cách "lách luật" một chút! 

## Minh Thảo đã làm gì?

### Ăn Yang Nè: Không sửa `printAreas` (Lách luật một chút!)

Mình đã nghĩ ra một cách để thêm tam giác mà không sửa hàm `printAreas`. Bí kíp là:
- Thêm các trường `Base` và `Height` vào struct `Shape` để hỗ trợ tam giác.
- Thêm một trường `Area` để lưu diện tích đã tính.
- Tạo một phương thức `CalculateArea()` để tính diện tích trước cho tất cả các hình (bao gồm tam giác) và lưu vào `Area`.
- Vì `printAreas` không xử lý tam giác, mình in diện tích của tam giác riêng trước khi gọi `printAreas`.

Code đây nè trong Before á =)))

**Tự bình phẩm một chútttt**: Ăn yang quá mom

---

### Cách 2: Làm đúng chuẩn OCP (Sửa `printAreas` cho đẹp luôn!)

Thực ra, để áp dụng OCP đúng cách, ta nên sửa lại hàm `printAreas` để nó không phụ thuộc vào `if-else`. Mình dùng một **interface** để mỗi hình tự tính diện tích của nó. Như thế, khi thêm hình mới, `printAreas` không cần sửa gì cả – đúng chuẩn OCP luôn!

Code chuẩn trong After á =)))

**Tự bình phẩm một chútttt**: Cách này đúng chuẩn OCP luôn! Thêm hình mới chỉ cần tạo struct mới và triển khai interface `Shape`, không cần đụng đến `printAreas`. Đỉnh của chóp mười điểm cho Minh Thảo!
