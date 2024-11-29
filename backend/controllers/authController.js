import Users from "../models/userModel.js";


export default  class UsersClass{
    constructor(){

    }
    async userfuction(req,res){
        const { username, contraseña } = req.body;
        try{
            const query = { username: username}
            const existsUser = await Users.findOne(query);
            if (existsUser){
                
                return res.status(200).json()
            }
        }catch(error){
            res.status(404).json({ErrorUser:error})
        }
    }
}