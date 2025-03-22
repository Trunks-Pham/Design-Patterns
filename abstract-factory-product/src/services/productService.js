const ProductFactory = require('../factories/productFactory');

class ProductService {
  async createProduct(type, data) {
    const product = ProductFactory.createProduct(type, data);
    return await product.save();
  }

  async getAllProducts() {
    return await ProductFactory.createProduct({}).constructor.find();
  }

  async getProductById(id) {
    return await ProductFactory.createProduct({}).constructor.findById(id);
  }

  async updateProduct(id, data) {
    return await ProductFactory.createProduct({}).constructor.findByIdAndUpdate(id, data, { new: true });
  }

  async deleteProduct(id) {
    return await ProductFactory.createProduct({}).constructor.findByIdAndDelete(id);
  }
}

module.exports = new ProductService();