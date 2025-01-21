package routes

import (
	"go-server/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func WebSocketRoutes(app *fiber.App, wsController *controllers.WebSocketController, audioController *controllers.AudioSocketController) {
	app.Get("/note", websocket.New(wsController.HandleYWebRTC))
	app.Get("/webrtc/audio", websocket.New(audioController.HandleWebRTC))
}
