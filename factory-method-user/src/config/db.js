const mongoose = require('mongoose');
require('dotenv').config();

class Database {
  constructor() {
    if (Database.instance) {
      return Database.instance; // Trả về instance đã tồn tại
    }

    Database.instance = this;
  }

  async connect() {
    if (this.connection) {
      console.log('Reusing existing MongoDB connection');
      return this.connection; // Trả về kết nối nếu đã có
    }

    try {
      this.connection = await mongoose.connect(process.env.MONGO_URI);
      console.log('MongoDB connected successfully');
      return this.connection;
    } catch (err) {
      console.error('MongoDB connection failed:', err.message);
      process.exit(1);
    }
  }
}

module.exports = new Database();
