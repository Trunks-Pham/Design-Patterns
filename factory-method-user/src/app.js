const express = require('express');
const bodyParser = require('body-parser');
const dotenv = require('dotenv');
const Database = require('./config/db');
const userRoutes = require('./routes/userRoutes');

dotenv.config();
const app = express();

// Kết nối MongoDB thông qua Singleton
Database.connect();

// Middleware
app.use(bodyParser.json());

// Routes
app.use('/api', userRoutes);

// Route kiểm tra server
app.get('/', (req, res) => {
  res.send('User Factory API with Singleton is running');
});

// Khởi động server
const PORT = process.env.PORT;
app.listen(PORT, () => console.log(`Server running on port ${PORT}`));
