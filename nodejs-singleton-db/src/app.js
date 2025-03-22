const express = require("express");
const app = express();
const userRoutes = require("./routes/userRoutes");
require("./config/database");

// Middleware xử lý JSON trước khi log request
app.use(express.json());

// Middleware log thông tin request
app.use((req, res, next) => {
    console.log(`[${new Date().toISOString()}]  ${req.method} ${req.url}`);
    
    if (Object.keys(req.body).length) {
        console.log(" Body:", req.body);
    }
    next();
});

// Routes
app.use("/api", userRoutes);

// Middleware xử lý lỗi chung
app.use((err, req, res, next) => {
    console.error(`[${new Date().toISOString()}] Lỗi:`, err.message);
    res.status(500).json({ error: "Lỗi server nội bộ" });
});

module.exports = app;
