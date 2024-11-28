import Product from "../models/productModels.js";


export default class Crud{
    constructor()
    {}

    async verProductos(req,res){
        try{
        const Products = await Product.find();
        res.status(200).json({Productos:Products})
    }catch(error){
        throw new Error({ErrorVerProductos:error})
    };
    };

    
    async crearProducto(req,res){
            const { name ,price ,image } = req.body;
            try{
            const newProduct = await Product.create(name,price,image)
            res.status(200).json({Producto_Nuevo: newProduct})
            }catch(error){
            throw Error({ErrorCrearProducto:error})
            };
    };


    async editarProducto(req,res){
            const { id }  = req.params;
            const { name, price, imagen } = req.body;
            try{
               if (id){
                const updateProduct = await Product.updateOne(
                    { _id: new ObjectId(id) },
                    {$set: {name, price,imagen}}
                );
                if(updateProduct.matchedCount === 0){
                    return res.status(404).json(ErrorEditarProducto)
                }
               }
               res.status(200).json({ProductoActualizado: updateProduct})
            }catch(error){
                res.status(400).json({ErrorActualizarProducto: error.message})
            };
    };


    async eliminarProducto(req,res){
        const { id }  = req.params;
        try{
            const deletedProduct = await Product.deleteOne(
                {_id: new ObjectId(id)}
            )
            if(deletedProduct.deletedCount === 0 ){
                return res.status(404).json({ProductoEliminado: deletedProduct,msj:("No encontrado")})
            }
            res.status(200).json({EliminadoCorrecto: deletedProduct})
        }catch(error){
            res.status(404).json({ErrorEliminarProducto: error.message})
        }
    };


};