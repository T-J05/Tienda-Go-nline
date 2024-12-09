package routes



import( 
	"github.com/gin-gonic/gin"
	"frontend/controllers"
)


func SetupRoutes(r *gin.Engine) {

	r.GET("/tiendaOnline", controllers.GetProducts)

	r.POST("/pedido", controllers.CreatePedidoHandler)

	r.POST("/agregar-al-carrito",controllers.AgregarAlCarrito)

	r.GET("/ver-carrito", controllers.VerCarrito)  

	r.POST("/confirmar-pedido", controllers.ConfirmarPedido)
	
}