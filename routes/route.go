package routes
import(
	"github.com/gofiber/fiber/v2"
	"github.com/NamDuongkiwi/shopping-backend/controller"
)

func ConfigProductRouter(router *fiber.Router) {
	(*router).Get("/", controller.GetAllProducts)
	(*router).Post("/", controller.CreateProduct)
	(*router).Get("/:id", controller.GetProductbyID)
	(*router).Delete("/:id", controller.DeleteProduct)
}