const Order = require("../models/Order");

class OrderService {
  async createOrder(product, quantity) {
    if (!product) throw new Error("Product not found!");
    const totalPrice = product.price * quantity;
    const order = new Order({ product: product._id, quantity, totalPrice });
    return await order.save();
  }

  async getAllOrders() {
    return await Order.find().populate("product");
  }
}

module.exports = new OrderService();
