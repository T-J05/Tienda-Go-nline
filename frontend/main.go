package main

import (

	"log" // Importa log para manejar errores
	"context" // Importa context para el uso de contexto en MongoDB
	"github.com/gin-gonic/gin"
	"frontend/routes"
	"frontend/config"
)

func main() {
	// Inicializa el router de Gin
	r := gin.Default()
	routes.SetupRoutes(r)

	// Inicializa la conexión a MongoDB
	err := config.InitMongoClient()
	if err != nil {
		log.Fatalf("Error al inicializar MongoDB: %v", err)
	}

	// Asegúrate de cerrar la conexión cuando termine
	defer func() {
		// Captura el cliente y el error de desconexión
		client, err := config.GetMongoClient()
		if err != nil {
			log.Fatalf("Error al obtener el cliente MongoDB: %v", err)
		}
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatalf("Error al desconectar MongoDB: %v", err)
		}
	}()

	// Ruta de prueba
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hola mundo",
		})
	})

	// Inicia el servidor Gin
	r.Run()
}
