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
	(*router).Patch("/:id", controller.ChangePrice)
	(*router).Get("/:id/price", controller.GetPriceHistory)
}

func ConfigReviewRouter(router *fiber.Router){
	(*router).Post("/", controller.CreateReview)
	(*router).Get("/:id", controller.GetReview)
}

func ConfigUserRouter(router *fiber.Router){
	(*router).Post("/", controller.CreateUser)
	(*router).Get("/:id", controller.GetUserByID)
}



