import mongoose from "mongoose"
import { ObjectId } from 'mongodb';


export default class Pedidos{
    constructor(){

    }

    async obtenerPedidos (req, res){
        try {
            const { token } = req.query || req.body;
            const db = mongoose.connection.db;
            const productsCollection = db.collection("Pedido");
            const pedidos = await productsCollection.find().toArray();
    
            if (pedidos.length === 0) {
               
                return res.status(404).json({ message: "No hay pedidos" });
            }
    
            return res.render('pedidos',{pedidos: pedidos,token:token})
    
        } catch (error) {
            res.status(400).json({ error: error.message, "message": "Error al obtener Pedido" });
        }
    };
    

    async eliminarPedido(req,res){
        try{
            const { id } = req.params;
            const objectid = new mongoose.Types.ObjectId(id)
            const token  = req.token;
           
            const db = mongoose.connection.db;
            const productsCollection = db.collection("Pedido");
            const pedidos = await productsCollection.find().toArray();
            const pedidoDelete = await productsCollection.findOneAndDelete(
                {_id:objectid}
            );

           res.redirect(`/pedidos/get?token=${token}`)
         
           
        }catch(error){
            res.json({"Error al eliminar el pedido":error.message})
        }
    }
}

