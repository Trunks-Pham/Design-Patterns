const User = require('../models/user');

class UserFactory {
  static createUser(role, { name, email }) {
    let userData = { name, email, role };

    switch (role) {
      case 'admin':
        userData.permissions = ['read', 'write', 'delete'];
        break;

      case 'member':
        userData.permissions = ['read', 'write'];
        break;

      case 'guest':
      default:
        userData.permissions = ['read'];
        break;
    }

    return new User(userData);
  }

  static getPermissions(role) {
    switch (role) {
      case 'admin':
        return ['read', 'write', 'delete'];
      case 'member':
        return ['read', 'write'];
      case 'guest':
      default:
        return ['read'];
    }
  }
}

module.exports = UserFactory;
