// main.go

package main

import (

	"log"
	"net/http"


	"frontend/config" // Importar el paquete donde está la conexión
)

func main() {
	// Definir la URI de la base de datos MongoDB
	uri := "mongodb://localhost:27017/tienda_online"

	// Establecer la conexión a MongoDB
	client, cancel, err := db.ConnectMongoDB(uri)
	if err != nil {
		log.Fatal(err)
	}
	defer cancel() // Asegurarse de que el contexto se cancele al final
	defer db.DisconnectMongoDB(client, client.Context()) // Asegurarse de desconectar MongoDB al final

	// Inicializar el servidor Gin
	r := gin.Default()

	// Ruta simple para verificar que el servidor está funcionando
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Pong"})
	})

	// Rutas adicionales de la API

	// Iniciar el servidor en el puerto 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error al ejecutar el servidor:", err)
	}
}
