import UsersClass from "../controllers/authController.js";
import express from "express";
const router = express.Router();
const userclass = new UsersClass();




router.post("/register",userclass.register)

router.post("/login",userclass.userfuction)

router.get("/login", (req,res) =>{
    try {
        console.log("tratando de visualizar")
        res.render('login');  // Renderiza la vista
      } catch (error) {
        console.error('Error al renderizar la vista:', error.message);
        res.status(500).send('Hubo un error al intentar renderizar la vista.');
      }
})


export default router;