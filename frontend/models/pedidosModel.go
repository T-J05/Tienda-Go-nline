package models

import "go.mongodb.org/mongo-driver/bson/primitive"


type Pedido struct {
    ID        primitive.ObjectID   `bson:"_id,omitempty"`
    Productos []ProductoPedido     `bson:"productos"`
    Direccion string               `bson:"direccion"`
}


type ProductoPedido struct {
    ID     primitive.ObjectID `bson:"id"`     
    Nombre string             `bson:"nombre"` 
    Precio float64            `bson:"precio"`
}