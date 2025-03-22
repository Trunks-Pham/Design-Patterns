const Product = require("../models/Product");

class ProductService {
  async getAllProducts() {
    return await Product.find();
  }

  async getProductById(productId) {
    return await Product.findById(productId);
  }
}

module.exports = new ProductService();
