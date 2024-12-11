package controllers

import (
	"context"
	"fmt"
	"frontend/config"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetProducts(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() 
	db, err := config.GetDatabase("tienda_online")
	if err != nil {
		log.Printf("Error al obtener la base de datos: %v", err) 
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "No se pudo conectar a MongoDB", "Detalle": err.Error()})
		return
	}


	collection := db.Collection("products")

	cursor, err := collection.Find(ctx, bson.D{{}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	var products []bson.M
	if err = cursor.All(context.TODO(), &products); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	session := sessions.Default(c)

    // Obtener y eliminar el mensaje de alerta de la sesión
    alerta := session.Get("alerta")
    if alerta != nil {
        session.Delete("alerta") // Borrar el mensaje para que no se muestre de nuevo
        session.Save()
    }

	fmt.Println("producto en go",products)

	c.HTML(http.StatusOK, "index.html", gin.H{"productos": products, "alerta": alerta})
	
}