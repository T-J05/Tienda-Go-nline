import Crud from "../controllers/productsController.js";
import express from "express";
const router = express.Router();
const crud = new Crud


export default
router.get("/producto", crud.verProductos);

router.post("/producto",crud.crearProducto);

router.patch("/producto",crud.editarProducto);

router.delete("/producto",crud.eliminarProducto);