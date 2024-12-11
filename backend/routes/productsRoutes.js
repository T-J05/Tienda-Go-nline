import Crud from "../controllers/productsController.js";
import express from "express";
import JWT from "../middlewares/jwt.js";
const router = express.Router();
const crud = new Crud();
const middlewares = new JWT();



router.get("/", crud.verProductos);

router.post("/",middlewares.verificarToken,middlewares.verificarRol,crud.crearProducto);

router.patch("/:id",middlewares.verificarToken,middlewares.verificarRol,crud.editarProducto);

router.delete("/:id",middlewares.verificarToken,middlewares.verificarRol,crud.eliminarProducto);



router.get('/editarProducto/:id',crud.mostrarFormulario)



export default router