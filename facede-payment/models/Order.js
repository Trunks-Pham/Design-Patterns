const mongoose = require("mongoose");

const OrderSchema = new mongoose.Schema({
  product: { type: mongoose.Schema.Types.ObjectId, ref: "Product" },
  quantity: Number,
  totalPrice: Number,
});

module.exports = mongoose.model("Order", OrderSchema);
