const mongoose = require('mongoose');
require('dotenv').config();

class Database {
  constructor() {
    if (!Database.instance) {
      console.log('MONGO_URI:', process.env.MONGO_URI); // Kiểm tra giá trị của MONGO_URI
      mongoose.connect(process.env.MONGO_URI)
        .then(() => console.log("MongoDB Connected"))
        .catch(err => console.error("MongoDB connection error:", err));

      Database.instance = this;
    }
    return Database.instance;
  }
}

module.exports = new Database();