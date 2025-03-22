const express = require("express");
const PaymentFacade = require("./PaymentFacade");
const ProductService = require("./services/ProductService");
const OrderService = require("./services/OrderService");

const router = express.Router();
 
router.get("/products", async (req, res) => {
  const products = await ProductService.getAllProducts();
  res.json(products);
});
  
router.post("/order", async (req, res) => {
  const { productId, quantity } = req.body;
  const result = await PaymentFacade.placeOrder(productId, quantity);
  res.json(result);
});
 
router.get("/orders", async (req, res) => {
  const orders = await OrderService.getAllOrders();
  res.json(orders);
});

module.exports = router;
