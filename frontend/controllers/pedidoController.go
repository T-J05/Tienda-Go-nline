package controllers

import (
	"context"
	"frontend/config"
	"frontend/models"
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/gin-contrib/sessions"
	"strconv"
	"strings"
	"github.com/pkg/errors"
)

func ConfirmarPedido(c *gin.Context) {
	// Obtener la sesión
	session := sessions.Default(c)

	// Obtener el carrito de la sesión
	carrito := session.Get("carrito")
	if carrito == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "El carrito está vacío"})
		return
	}

	// Convertir el carrito en un slice de ProductoPedido
	productosCarrito, ok := carrito.([]models.ProductoPedido)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al procesar el carrito"})
		return
	}

	// Crear el pedido
	nuevoPedido := models.Pedido{
		Productos: productosCarrito,
	}

	// Llamar a CreatePedido para insertar el pedido en la base de datos
	CreatePedido(c, nuevoPedido)

	// Limpiar el carrito
	session.Delete("carrito")
	session.Save()
	// Verificar el estado de la sesión después de limpiarlo
	fmt.Println("Estado del carrito después de eliminar:", session.Get("carrito"))

	session.Set("alerta", "Gracias por su compra!")
    session.Save()

	c.Redirect(http.StatusFound, "/tiendaOnline")
}


func CreatePedido(c *gin.Context, pedido models.Pedido) {
	
	collection, err := config.GetPedidosCollection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al conectar a la base de datos"})
		return
	}

	if pedido.ID.IsZero() {
		pedido.ID = primitive.NewObjectID()
	}

	_, err = collection.InsertOne(context.TODO(), pedido)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al guardar el pedido en la base de datos"})
		return
	}


}

func AgregarAlCarrito(c *gin.Context) {
	// Obtener la sesión
	session := sessions.Default(c)

	// Obtener los datos del formulario
	productoID := c.PostForm("id")
	productoID = strings.TrimPrefix(productoID, "ObjectID(\"")
	productoID = strings.TrimSuffix(productoID, "\")")   
	productoNombre := c.PostForm("nombre")
	productoPrecio := c.PostForm("precio")


	objectID, err := primitive.ObjectIDFromHex(productoID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID de producto inválido"})
		return
	}

	// Crear un objeto de producto
	precio, err := strconv.ParseFloat(productoPrecio, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Precio inválido"})
		return
	}

	producto := models.ProductoPedido{
		ID:     objectID,
		Nombre: productoNombre,
		Precio: precio,
	}

	// Obtener el carrito de la sesión (si existe)
	carrito := session.Get("carrito")
	
	if carrito == nil {
		carrito = []models.ProductoPedido{}
	}
	fmt.Println("carritoSession", carrito)
	carrito = append(carrito.([]models.ProductoPedido), producto)
	fmt.Println("carritoProdut", carrito)
	session.Set("carrito", carrito)
	fmt.Println("Sesión antes de guardar carrito:", session.Get("carrito"))

	if err := session.Save(); err != nil {
        wrappedErr := errors.Wrap(err, "Error al intentar guardar la sesión")
        c.JSON(http.StatusInternalServerError, gin.H{
            "message": "Error al añadir el producto al carrito. Inténtalo nuevamente.",
            "error":   wrappedErr.Error(), 
        })
        fmt.Printf("Error detallado: %+v\n", wrappedErr)
        return
    }

	session.Set("alerta", "Producto agregado correctamente al carrito!")
    session.Save()
	c.Redirect(http.StatusFound, "/tiendaOnline")
}

func VerCarrito(c *gin.Context) {
	// Obtener la sesión
	session := sessions.Default(c)

	// Obtener el carrito de la sesión
	carrito := session.Get("carrito")
	fmt.Println("carrito en verCarrito luego de obtener el carrito de la sesion",carrito)
	if carrito == nil {
		carrito = []models.ProductoPedido{}
	}
	
	// Calcular el total del carrito
	var total float64
	for _, producto := range carrito.([]models.ProductoPedido) {
		total += producto.Precio
	}

	// Renderizar la página con el carrito y el total
	c.HTML(http.StatusOK, "carrito.html", gin.H{
		"carrito": carrito,
		"total":   total,
	})
}

func CreatePedidoHandler(c *gin.Context) {
    // Extraer el cuerpo de la solicitud y llamar a la función CreatePedido
    var pedido models.Pedido
    if err := c.ShouldBindJSON(&pedido); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Datos inválidos"})
        return
    }
    CreatePedido(c, pedido)
}
