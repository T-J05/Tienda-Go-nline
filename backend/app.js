import express, { json } from "express"
import jwt from "jsonwebtoken"
import { configDotenv } from "dotenv"
import productsRoutes from "./routes/productsRoutes.js" 
import connectDB from "./config/databaseJS.js"

connectDB();

const app = express();
const port = process.env.PORT;

app.set('view engine','ejs');
app.use(express.json());
app.use(express.urlencoded({ extended: true }));


app.use("/producto",productsRoutes);

app.listen(port,() => {
    console.log("Servidor corriendo...")
});
