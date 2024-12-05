package controllers

import (
	"context"
	"frontend/config"
	"frontend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)



func CreatePedido(c *gin.Context){

	collection, err := config.GetPedidosCollection()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error al conectar a la base de datos",
		})
		return
	}
	
	var nuevoPedido models.Pedido
	if err := c.ShouldBindJSON(&nuevoPedido); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
		})
		return
	}

	// Paso 4: Asignar un ID único al pedido si no tiene uno (Mongo lo generaría automáticamente)
	if nuevoPedido.ID.IsZero() {
		nuevoPedido.ID = primitive.NewObjectID()
	}

	// Paso 5: Insertar el pedido en la colección 'Pedidos'
	_, err = collection.InsertOne(context.TODO(), nuevoPedido)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error al guardar el pedido en la base de datos",
		})
		return
	}

	// Paso 6: Devolver una respuesta indicando que el pedido fue creado con éxito
	c.JSON(http.StatusOK, gin.H{
		"message": "Pedido creado con éxito",
		"id":      nuevoPedido.ID.Hex(), // Devuelves el ID generado para el nuevo pedido
	})

}

func VerPedidos(c *gin.Context){
	

}