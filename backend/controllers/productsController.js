import Product from "../models/productModels.js";
import { ObjectId } from 'mongodb';


export default class Crud{
    constructor()
    {}

    async verProductos(req,res){
        try{
        const Products = await Product.find();
        res.render("productos",{ productos: Products });
    }catch(error){
        res.send({ErrorVerProductos:error.message})
    };
    };

    
    async crearProducto(req,res){
            const { name ,price ,image } = req.body;
            try{
            const newProduct = await Product.create({
                name:name,
                price:price,
                image:image
            })
            res.redirect("/productos");
            }catch(error){
            throw Error({ErrorCrearProducto:error})
            };
    };


    async editarProducto(req,res){
            const { id }  = req.params;
            const { name, price, imagen } = req.body;
            try{
               if (id){
                const updateProduct = await Product.findByIdAndUpdate(
                    id,
                    {$set: {name, price,imagen}},
                    { new: true, runValidators: true }
                    
                );
                if(updateProduct.matchedCount === 0){
                    return res.status(404).json(ErrorEditarProducto)
                }
                res.redirect("/productos");
               }
               
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
            res.redirect("/productos");
        }catch(error){
            res.status(404).json({ErrorEliminarProducto: error.message})
        }
    };


};