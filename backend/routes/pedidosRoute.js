import express from "express";
import Pedido from "../controllers/pedidosController.js";

const pedido = new Pedido();
const router = express.Router()




router.get("/get",pedido.obtenerPedidos)



export default router