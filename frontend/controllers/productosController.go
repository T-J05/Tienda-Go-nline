package controllers

import (
	"context"
	"frontend/config"
	"net/http"
	"time"
	"log"
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

	c.JSON(http.StatusOK, products)
}