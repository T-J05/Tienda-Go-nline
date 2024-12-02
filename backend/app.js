import express, { json } from "express";
import jwt from "jsonwebtoken";
import dotenv from "dotenv";
import productsRoutes from "./routes/productsRoutes.js";
import authRoutes from "./routes/authRoutes.js";
import connectDB from "./config/databaseJS.js";

dotenv.config();

connectDB();

const app = express();
const port = process.env.PORT;

app.set('view engine','pug');
app.set('views', './views');
app.use(express.json());
app.use(express.urlencoded({ extended: true }));


app.use("/producto",productsRoutes);
app.use("/admin",authRoutes);
app.listen(port,() => {
    console.log(`Servidor corriendo en el puerto ${port}`)
});
