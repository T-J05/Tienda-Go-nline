import express from "express";
import Pedido from "../controllers/pedidosController.js";
import JWT from "../middlewares/jwt.js";

const auth = new JWT();
const pedido = new Pedido();
const router = express.Router()




router.get("/get",auth.verificarToken,pedido.obtenerPedidos)

router.delete("/eliminarPedido/:id",auth.verificarToken,pedido.eliminarPedido)


export default router