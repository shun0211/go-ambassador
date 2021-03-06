package routes

import (
	"ambassador/src/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// MEMO: app.Post("/api/admin/register", controllers.Register)は下のように書ける
	api := app.Group("api")
	admin := api.Group("admin")
	admin.Post("/register", controllers.Register)
	admin.Post("/login", controllers.Login)
	admin.Get("/user", controllers.User)
	admin.Post("/logout", controllers.Logout)
}
