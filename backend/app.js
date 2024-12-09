import express, { json } from "express";
import dotenv from "dotenv";
import productsRoutes from "./routes/productsRoutes.js";
import authRoutes from "./routes/authRoutes.js";
import pedidosRoutes from "./routes/pedidosRoute.js"
import connectDB from "./config/databaseJS.js";
import methodOverride from "method-override"
dotenv.config();



connectDB();

const app = express();
const port = process.env.PORT;

app.set('view engine', 'pug');
app.set('views', 'views');

app.use(methodOverride('_method'));
app.use(express.static('public'));
app.use(express.json());
app.use(express.urlencoded({ extended: true }));


app.use("/producto",productsRoutes);
app.use("/admin",authRoutes);
app.use("/pedidos", pedidosRoutes)

app.listen(port,() => {
    console.log(`Servidor corriendo en el puerto ${port}`)
});
