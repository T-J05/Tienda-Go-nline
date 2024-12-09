import jwt from "jsonwebtoken"


export default class JWT{
    constructor(){

    }

    generarJwt(payload,secretKey){
        return jwt.sign(payload,secretKey,{algorithm:"HS256",expiresIn: "1h"})
    }
                                                
    
    verificarToken(req, res, next) {
        const token = req.body.token || req.query.token || req.headers['authorization'];
        console.log( token )
        if (!token) {
            return res.status(401).json({ error: 'Token no proporcionado' });
        }
    
        try {
            const decoded = jwt.verify(token, process.env.SECRET_KEY);
            req.user = decoded; 
            req.token = token
            next();
        } catch (error) {
            res.status(401).json({ error: 'Token inválido', message: error.message });
        }
    };
    

    verificarRol(req,res,next){
        const roleUser = req.user.role;
        try{
            if(roleUser && roleUser === "admin"){
                next();
            }
        }catch(error){
            return res.json({ErrorVerificarRol: error.message})
        }
    }
}