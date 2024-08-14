package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thitiphongD/go-backend/controllers"
)

func SetUpRoutes(app *fiber.App) {
	// Test route to verify application setup
	app.Get("/", controllers.Hello)
	app.Post("/api/signUp", controllers.SignUp)
	app.Post("/api/signIn", controllers.SignIn)

	app.Get("/api/manga", controllers.GetMangas)
	app.Get("/api/manga/:id", controllers.GetMangas)
	app.Post("/api/manga", controllers.AddManga)
	app.Put("/api/manga/:id", controllers.UpdateManga)
	app.Delete("/api/manga/:id", controllers.RemoveManga)
}
