import Product from "../models/productModels";


export default class Crud{
    constructor()
    {}

    async verProductos(req,res){
        try{
        const Products = await Product.find();
        res.status(200).json({Productos:Products})
    }catch(error){
        throw new Error({ErrorVerProductos:error})
    }
    }

    
    async crearProducto(req,res){
            const {name,price,image} = req.body;
            try{
            const newProduct = await Product.create(name,price,image)
            res.status(200).json({Producto_Nuevo:newProduct})
            }catch(error){
            throw Error({ErrorCrearProducto:error})
            }
    }


    editarProducto(req,res){

    }


    eliminarProducto(req,res){

    }


}