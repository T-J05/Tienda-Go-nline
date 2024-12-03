import Users from "../models/userModel.js";
import JWT from "../middlewares/jwt.js";
import bcrypt from "bcrypt"
import dotenv from "dotenv";
import time from "timers"

dotenv.config();
const jwt = new JWT();
const secret_key = process.env.SECRET_KEY;
const salt_rouds = parseInt(process.env.SALT_ROUNDS, 10);


export default  class UsersClass{
    constructor(){

    }

    async register(req,res){
        const { username, contraseña, role } = req.body;
        try{
            const query = { username: username}
            const existe = await Users.findOne(query)
            if(!existe)
            {
            const hashcontraseña = await bcrypt.hash(contraseña,salt_rouds)
            const userNew = await Users.create({
                username: username,
                contraseña: hashcontraseña,
                role: role || "user"
            })
        }
            else{
                res.status(404).json({ExisteYa: username})
            }
            res.status(200).json({RegisterExitoso: username})
        }catch(error){
            res.json({ErrorRegister: error.message})
        }
    }

    async userfuction(req,res){
        const { username, contraseña } = req.body;
        try{
            const query = { username: username}
            const existsUser = await Users.findOne(query);
            const verificado = await bcrypt.compare(contraseña ,existsUser.contraseña)
            if (existsUser && verificado){
                if (existsUser.role == "admin"){
                    const payload = {
                        role: existsUser.role,
                        username: username
                    }
                    const token = jwt.generarJwt(payload,secret_key);
                    res.cookie("token",token).redirect("/producto")
                }
           
            }
            else{
                res.status(404).json("Contraseñas no coinciden")
                setTimeout(() => {
                    res.render("login");
                  }, 3000);
            }
           
        }catch(error){
            res.status(404).json({ErrorUser:error})
        }
    }

    login_render(req,res,next){
        res.render("login")
        next();
    }

    redirectt(req,res){
        res.redirect("/producto")
    }
}