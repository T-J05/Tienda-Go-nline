package routes



import( 
	"github.com/gin-gonic/gin"
	"frontend/controllers"
)


func SetupRoutes(r *gin.Engine) {

	r.GET("/tiendaOnline", controllers.GetProducts)


	r.POST("/pedido", controllers.CreatePedido)

	r.GET("pedido", controllers.VerPedidos)
	
}