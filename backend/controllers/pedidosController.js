import mongoose from "mongoose"
import connectDB from "../config/databaseJS.js"


const obtenerPedidos = async (req, res)=> {
    try{
    const db = mongoose.connection.db;
    const productsCollection = db.collection("Pedido");
    const products = await productsCollection.find().toArray();

    }
   
    catch(error){
        res.status(400).json({error: error.message, "message": "Error al obtener Pedido"})

    }
   
}