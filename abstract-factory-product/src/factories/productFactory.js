const Product = require('../models/product');

class ProductFactory {
  createProduct(type, data) {
    switch (type) {
      case 'phone':
        return new Product({ ...data, category: 'Phone', warranty: '12 months' });
      case 'laptop':
        return new Product({ ...data, category: 'Laptop', warranty: '24 months' });
      case 'accessory':
        return new Product({ ...data, category: 'Accessory', warranty: '6 months' });
      case 'tablet':
        return new Product({ ...data, category: 'Tablet', warranty: '18 months' });
      case 'smartwatch':
        return new Product({ ...data, category: 'Smartwatch', warranty: '12 months' });
      case 'desktop':
        return new Product({ ...data, category: 'Desktop', warranty: '36 months' });
      default:
        return new Product(data); // Loại sản phẩm chung
    }
  }
}

module.exports = new ProductFactory();