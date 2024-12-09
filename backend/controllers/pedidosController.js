import mongoose from "mongoose"


export default class Pedidos{
    constructor(){

    }

    async obtenerPedidos (req, res){
        try {
            const db = mongoose.connection.db;
            const productsCollection = db.collection("Pedido");
            const pedidos = await productsCollection.find().toArray();
    
            if (pedidos.length === 0) {
               
                return res.status(404).json({ message: "No hay pedidos" });
            }
    
            return res.render('pedidos',{pedidos: pedidos})
    
        } catch (error) {
            res.status(400).json({ error: error.message, "message": "Error al obtener Pedido" });
        }
    };
}

