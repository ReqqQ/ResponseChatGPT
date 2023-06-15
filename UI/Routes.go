package UI

import (
	"github.com/gofiber/fiber/v2"

	"ResponseChatGPT/UI/API"
)

type Controllers struct {
	ResponseController API.ResponseController
}

func (co Controllers) InitGetRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1", func(c *fiber.Ctx) error {
		c.Set("Version", "v1")
		return c.Next()
	})
	v1.Post("/response/message/generateTooltip", func(c *fiber.Ctx) error {
		return c.JSON(co.ResponseController.GetMessageTooltip(c))
	})
}
