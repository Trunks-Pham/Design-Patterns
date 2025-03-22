const mongoose = require("mongoose");
require("dotenv").config();

class Database {
  constructor() {
    if (!Database.instance) {
      this._connect();
      Database.instance = this; // Lưu instance duy nhất
    }
    return Database.instance;
  }

  _connect() {
    const MONGO_URI = process.env.MONGO_URI;

    mongoose
      .connect(MONGO_URI)
      .then(() => console.log("Kết nối MongoDB thành công!"))
      .catch((err) => console.error("Lỗi kết nối MongoDB:", err));

    this.connection = mongoose.connection;

    // Xử lý khi kết nối bị mất
    this.connection.on("error", (err) => {
      console.error("Mất kết nối MongoDB:", err);
    });

    // Khi server tắt, đóng kết nối MongoDB
    process.on("SIGINT", async () => {
      await mongoose.connection.close();
      console.log("Đã ngắt kết nối MongoDB do ứng dụng Node.js đã tắt");
      process.exit(0);
    });
  }

  getConnection() {
    return this.connection;
  }
}

module.exports = new Database();