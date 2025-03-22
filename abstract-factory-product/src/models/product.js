const mongoose = require('mongoose');

const ProductSchema = new mongoose.Schema({
  name: { type: String, required: true },         
  category: { type: String, required: true },     
  description: { type: String },                  
  price: { type: Number, required: true },        
  stock: { type: Number, default: 0 },        
  brand: { type: String },                        
  sku: { type: String, unique: true },          
  images: [{ type: String }],                 
  weight: { type: Number },                   
  dimensions: {                                   
    length: { type: Number },
    width: { type: Number },
    height: { type: Number }
  },                                             
  isAvailable: { type: Boolean, default: true },  
  createdAt: { type: Date, default: Date.now },  
  updatedAt: { type: Date, default: Date.now }   
});

module.exports = mongoose.model('Product', ProductSchema);
