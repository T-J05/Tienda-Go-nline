import UsersClass from "../controllers/authController.js";
import express from "express";
const router = express.Router();
const userclass = new UsersClass();




router.post("/register",userclass.register)

router.post("/login",userclass.userfuction)

router.get("/login")

export default router;