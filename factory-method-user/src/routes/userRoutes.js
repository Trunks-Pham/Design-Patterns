const express = require('express');
const router = express.Router();
const userController = require('../controllers/userController');

// Định tuyến CRUD
router.post('/users', userController.createUser);
router.get('/users', userController.getUsers);
router.get('/users/:role', userController.getUsersByRole);
router.put('/users/:id', userController.updateUser);
router.delete('/users/:id', userController.deleteUser); 

module.exports = router;
