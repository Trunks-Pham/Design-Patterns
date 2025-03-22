const UserFactory = require('../factories/userFactory');
const User = require('../models/user');
const { HTTP_STATUS_CODES } = require('../httpStatus/statusCodes');

// Tạo người dùng mới
const createUser = async (req, res) => {
  try {
    const { name, email, role } = req.body;

    // Factory quyết định tạo loại người dùng nào
    const user = UserFactory.createUser(role, { name, email });

    await user.save();
    res.status(HTTP_STATUS_CODES.CREATED).json({ message: 'User created successfully', user });
  } catch (err) {
    res.status(HTTP_STATUS_CODES.BAD_REQUEST).json({ error: err.message });
  }
};
 
// Lấy danh sách người dùng
const getUsers = async (req, res) => {
  try {
    const users = await User.find();
    if (users.length === 0) {
      res.status(HTTP_STATUS_CODES.OK).json({ message: 'No users found' });
    } else {
      res.status(HTTP_STATUS_CODES.OK).json(users);
    }
  } catch (err) {
    res.status(HTTP_STATUS_CODES.INTERNAL_SERVER_ERROR).json({ error: err.message });
  }
};

// Tìm người dùng theo role
const getUsersByRole = async (req, res) => {
  try {
    const { role } = req.params;
    const users = await User.find({ role });

    if (users.length === 0) {
      res.status(HTTP_STATUS_CODES.OK).json({ message: `No users found with role ${role}` });
    } else {
      res.status(HTTP_STATUS_CODES.OK).json(users);
    }
  } catch (err) {
    res.status(HTTP_STATUS_CODES.INTERNAL_SERVER_ERROR).json({ error: err.message });
  }
};

// Cập nhật người dùng
const updateUser = async (req, res) => {
  try {
    const { id } = req.params;
    const { role } = req.body;
    const updatedUser = await User.findByIdAndUpdate(id, req.body, { new: true });

    if (!updatedUser) return res.status(HTTP_STATUS_CODES.NOT_FOUND).json({ message: 'User not found' });

    const updatedPermissions = UserFactory.getPermissions(role);
    updatedUser.permissions = updatedPermissions;

    // Lưu lại thay đổi vào database
    await updatedUser.save();

    res.status(HTTP_STATUS_CODES.OK).json({ message: 'User updated', user: updatedUser });
  } catch (err) {
    res.status(HTTP_STATUS_CODES.BAD_REQUEST).json({ error: err.message });
  }
};

// Xóa người dùng
const deleteUser = async (req, res) => {
  try {
    const { id } = req.params;
    const deletedUser = await User.findByIdAndDelete(id);

    if (!deletedUser) return res.status(HTTP_STATUS_CODES.NOT_FOUND).json({ message: 'User not found' });

    res.status(HTTP_STATUS_CODES.OK).json({ message: 'User deleted' });
  } catch (err) {
    res.status(HTTP_STATUS_CODES.BAD_REQUEST).json({ error: err.message });
  }
};

module.exports = { createUser, getUsers, getUsersByRole, updateUser, deleteUser };