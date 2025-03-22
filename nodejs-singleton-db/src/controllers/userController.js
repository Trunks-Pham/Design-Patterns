const User = require("../models/User");

// Tạo user
exports.createUser = async (req, res) => {
    try {
      const { email } = req.body;
      const existingUser = await User.findOne({ email });
      if (existingUser) {
        return res.status(400).json({ error: "Email đã tồn tại" });
      }
  
      console.log(`[${new Date().toISOString()}] Đang tạo user với dữ liệu:`, req.body);
      const user = await User.create(req.body);
      console.log(`[${new Date().toISOString()}] User đã tạo:`, user);
      res.status(201).json(user);
    } catch (error) {
      console.error(`[${new Date().toISOString()}] Lỗi khi tạo user:`, error.message);
      
      if (error.name === "ValidationError") {
        return res.status(400).json({ error: "Dữ liệu không hợp lệ" });
      }
  
      res.status(500).json({ error: "Lỗi server" });
    }
  };

// Lấy danh sách user
exports.getUsers = async (req, res) => {
  try {
    console.log(`[${new Date().toISOString()}] Đang lấy danh sách user...`);
    const users = await User.find();
    console.log(`[${new Date().toISOString()}] Tổng số user: ${users.length}`);

    if (users.length === 0) {
      res.status(200).json({ message: "Không có user" });
    } else {
      res.status(200).json(users);
    }
  } catch (error) {
    console.error(`[${new Date().toISOString()}] Lỗi khi lấy user:`, error.message);
    res.status(500).json({ error: "Lỗi server" });
  }
};

// Xóa user
exports.deleteUser = async (req, res) => {
  try {
    const { id } = req.params;
    const user = await User.findByIdAndDelete(id);
    if (!user) {
      return res.status(404).json({ error: "User không tồn tại" });
    }
    res.status(200).json({ message: "Đã xóa thành công" });
  } catch (error) {
    console.error(`[${new Date().toISOString()}] Lỗi khi xóa user:`, error.message);
    res.status(500).json({ error: "Lỗi server" });
  }
};

// Cập nhật user
exports.updateUser = async (req, res) => {
  try {
    const { id } = req.params;
    const user = await User.findByIdAndUpdate(id, req.body, { new: true });
    if (!user) {
      return res.status(404).json({ error: "User không tồn tại" });
    }
    res.status(200).json({ message: "Đã cập nhật thành công" });
  } catch (error) {
    console.error(`[${new Date().toISOString()}] Lỗi khi update user:`, error.message);
    res.status(500).json({ error: "Lỗi server" });
  }
};