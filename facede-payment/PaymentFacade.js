const ProductService = require("./services/ProductService");
const OrderService = require("./services/OrderService");
const PaymentService = require("./services/PaymentService");

class PaymentFacade {
  async placeOrder(productId, quantity) {
    try {
      const product = await ProductService.getProductById(productId);
      const order = await OrderService.createOrder(product, quantity);
      const payment = PaymentService.processPayment(order.totalPrice);
      return { order, payment };
    } catch (error) {
      return { error: error.message };
    }
  }
}

module.exports = new PaymentFacade();
