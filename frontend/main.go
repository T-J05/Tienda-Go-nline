package main

import (
	"log" // Importa log para manejar errores
	"context" // Importa context para el uso de contexto en MongoDB
	"encoding/gob"
	"github.com/gin-gonic/gin"
	"frontend/routes"
	"frontend/config"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"frontend/models"
)


func init() {
    // Registrar el tipo ProductoPedido para su uso con gob
	gob.Register(models.ProductoPedido{})  // Registro del tipo ProductoPedido
    gob.Register([]models.ProductoPedido{})  // Registro de un slice de ProductoPedido
}


func main() {
	// Inicializa el router de Gin
	r := gin.Default()

	// Configura el almacenamiento de sesiones en cookies
	store := cookie.NewStore([]byte("clave_muyyyy_secreta"))
	r.Use(sessions.Sessions("josexd", store))

	// Configura las rutas
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
	
	r.Static("/static", "./static")
	
	// Carga las plantillas HTML
	r.LoadHTMLGlob("templates/*")

	// Inicia el servidor Gin
	r.Run()
}
