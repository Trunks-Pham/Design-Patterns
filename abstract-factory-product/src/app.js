require('dotenv').config(); 
const express = require('express');
const bodyParser = require('body-parser');
require('./config/db'); 
const productRoutes = require('./routes/productRoutes');

const app = express();
const PORT = process.env.PORT;

// Middleware
app.use(bodyParser.json());

// Routes
app.use('/api', productRoutes);

app.get('/', (req, res) => {
  res.send('Product Management API is running');
});

app.listen(PORT, () => {
  console.log(`Server running on http://localhost:${PORT}`);
});