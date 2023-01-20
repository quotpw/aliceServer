package helpers

import (
	"aliceServer/alice"
	"aliceServer/device"
	"aliceServer/kuzya"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"gorm.io/gorm"
)

func InitWebApp(database *gorm.DB) *fiber.App {
	webApp := fiber.New(fiber.Config{
		// Pass view engine
		//Views: html.New("./views", ".html"),
		// Pass global error handler
		//ErrorHandler: handlers.Errors("./public/500.html"),
	})

	// Serve alice post requests
	webApp.Post("/whalice", alice.Webhook(database))

	// Serve kuzya get requests
	webApp.Get("/kuzya/:device/functions/:function", kuzya.Webhook(database))
	webApp.Post("/kuzya/:device/functions/:function", kuzya.Webhook(database))

	// Serve device websockets requests
	webApp.Get("/ws", websocket.New(device.Websocket(database), websocket.Config{}))

	//// Serve static assets
	//webApp.Static("/", "./public", fiber.Static{
	//	Compress: true,
	//})
	//
	//// Handle 404 errors
	//webApp.Use(handlers.NotFound("./public/404.html"))
	return webApp
}
