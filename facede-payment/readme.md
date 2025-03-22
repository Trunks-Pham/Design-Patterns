### **Ý tưởng về `PaymentFacade`**  

`PaymentFacade` đóng vai trò là một **lớp trung gian** giúp đơn giản hóa quy trình đặt hàng và thanh toán. Thay vì để client (người dùng hoặc API) trực tiếp gọi từng service riêng lẻ (`ProductService`, `OrderService`, `PaymentService`), `PaymentFacade` gom tất cả các bước này lại thành một phương thức duy nhất:  

### **Mục tiêu chính của `PaymentFacade`**:
1. **Đơn giản hóa quy trình thanh toán**  
   - Không cần gọi từng service một cách thủ công.  
   - Chỉ cần gọi `PaymentFacade.placeOrder()`, nó sẽ tự động xử lý tất cả các bước.  

2. **Giảm sự phụ thuộc giữa các thành phần**  
   - Nếu sau này muốn thay đổi logic xử lý đơn hàng, chỉ cần sửa trong `PaymentFacade`, không cần sửa nhiều nơi trong hệ thống.  

3. **Dễ bảo trì & mở rộng**  
   - Nếu cần thêm chức năng (ví dụ: kiểm tra tồn kho, gửi email xác nhận), chỉ cần cập nhật trong `PaymentFacade` mà không ảnh hưởng đến các service khác.  

---

### **Cách hoạt động của `PaymentFacade`**
Khi một người dùng muốn **mua một sản phẩm**, thay vì gọi từng bước như sau:  
1. Lấy thông tin sản phẩm từ **ProductService**  
2. Tạo đơn hàng bằng **OrderService**  
3. Xử lý thanh toán bằng **PaymentService**  

Thì `PaymentFacade` giúp gom tất cả lại thành **một bước duy nhất**! 🎯  

**Mã nguồn `PaymentFacade.js`**:
```javascript
const ProductService = require("./services/ProductService");
const OrderService = require("./services/OrderService");
const PaymentService = require("./services/PaymentService");

class PaymentFacade {
  async placeOrder(productId, quantity) {
    try {
      // 1 Lấy thông tin sản phẩm
      const product = await ProductService.getProductById(productId);
      if (!product) throw new Error("Product not found!");

      // 2 Tạo đơn hàng
      const order = await OrderService.createOrder(product, quantity);

      // 3 Thanh toán
      const payment = PaymentService.processPayment(order.totalPrice);

      // Trả về kết quả cuối cùng
      return { order, payment };
    } catch (error) {
      return { error: error.message };
    }
  }
}

module.exports = new PaymentFacade();
```

---

###  **Ví dụ sử dụng `PaymentFacade`**
Trong API `POST /api/order`, thay vì gọi từng service, chỉ cần gọi `PaymentFacade`:
```javascript
router.post("/order", async (req, res) => {
  const { productId, quantity } = req.body;
  const result = await PaymentFacade.placeOrder(productId, quantity);
  res.json(result);
});
```
 **Chỉ với một dòng `PaymentFacade.placeOrder()`, API đã thực hiện tất cả các bước!** 

---

## **Tóm tắt lại lợi ích của `PaymentFacade`**
 **Giúp code ngắn gọn hơn** (gom nhiều thao tác vào một hàm duy nhất).  
 **Giảm sự phụ thuộc giữa các service** (các thành phần không cần gọi trực tiếp nhau).  
 **Dễ dàng mở rộng & bảo trì** (nếu muốn thêm chức năng, chỉ cần sửa trong `PaymentFacade`).  
 